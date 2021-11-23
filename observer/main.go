package main

import (
	"alcor/db"
	"alcor/model/candidate"
	"alcor/model/info"
	"alcor/service/auth"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {
	db.Init()

	uri := os.Getenv("URI")

	res, err := http.Get(fmt.Sprintf("http://%s:9999/bundle/last", uri))
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		panic(res.Status)
	}
	currentBundle, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	lastBundleHash := make([]byte, len(currentBundle))
	copy(lastBundleHash, currentBundle)

	ctx := context.Background()

	for {
		b := false
		for _, v := range currentBundle {
			if v != 0 {
				b = true
				break
			}
		}
		if !b {
			break
		}

		bundleID := base64.URLEncoding.EncodeToString(currentBundle)
		fmt.Println("bundle: ", hex.EncodeToString(currentBundle))

		bundledb := new(db.Bundle)
		bundledb.Hash = currentBundle
		if err := db.SelectBundle(ctx, bundledb); err == nil {
			log.Println(bundleID + " already exists")
			break
		}

		res, err = http.Get(fmt.Sprintf("http://%s:9999/bundle/info/%s", uri, bundleID))
		if err != nil {
			panic("panic: " + bundleID + ": " + err.Error())
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println(err)
			}
			panic("panic: " + bundleID + ": " + res.Status + ": " + string(data))
		}

		bundlepb := new(info.Bundle)
		buffers, err := io.ReadAll(res.Body)
		if err != nil {
			panic("panic: " + bundleID + ": " + err.Error())
		}
		if err := proto.Unmarshal(buffers, bundlepb); err != nil {
			panic("panic: " + bundleID + ": " + err.Error())
		}

		bundledb.Hash = bundlepb.Hash
		bundledb.Prev = bundlepb.Prev
		bundledb.SubHashes = bundlepb.SubHashes
		if err := db.InsertBundle(ctx, bundledb); err != nil {
			panic("panic: " + bundleID + ": " + err.Error())
		}
		fmt.Println("sub hashes: ", bundlepb.SubHashes)

		for _, subHash := range bundlepb.SubHashes {
			paperid := base64.URLEncoding.EncodeToString(subHash)
			log.Println("subhash: ", hex.EncodeToString(subHash))

			res, err = http.Get(fmt.Sprintf("http://%s:9999/paper/info/%s", uri, paperid))
			if err != nil {
				panic("panic: " + paperid + ": " + err.Error())
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				panic("panic: " + paperid + ": " + res.Status)
			}

			paperpb := new(info.Paper)
			buffers, err := io.ReadAll(res.Body)
			if err != nil {
				panic("panic: " + paperid + ": " + err.Error())
			}
			if err := proto.Unmarshal(buffers, paperpb); err != nil {
				panic("panic: " + paperid + ": " + err.Error())
			}

			paperdb := new(db.Paper)
			paperdb.Hash = paperpb.Hash
			if err := db.SelectPaper(ctx, paperdb); err == nil {
				log.Println(paperid + " already exists")
				continue
			}

			paperdb.Message = paperpb.Message
			paperdb.RandomBytes = paperpb.RandomBytes
			paperdb.Signature = paperpb.Signature
			paperdb.Timestamp = paperpb.Timestamp
			paperdb.VoterID = paperpb.VoterID
			if err := db.InsertPaper(ctx, paperdb); err != nil {
				panic("panic: " + paperid + ": " + err.Error())
			}

			voterpb := new(info.Voter)
			voterpb.HashID = paperpb.VoterID
			voterid := base64.URLEncoding.EncodeToString(paperpb.VoterID)
			res, err := http.Get(fmt.Sprintf("http://%s:9999/voter/info/%s", uri, voterid))
			if err != nil {
				panic("panic: " + hex.EncodeToString(paperpb.VoterID) + ": " + err.Error())
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				panic("panic: " + hex.EncodeToString(paperpb.VoterID) + ": " + res.Status)
			}
			buffers, err = io.ReadAll(res.Body)
			if err != nil {
				panic("panic: " + hex.EncodeToString(paperpb.VoterID) + ": " + err.Error())
			}
			if err := proto.Unmarshal(buffers, voterpb); err != nil {
				panic("panic: " + hex.EncodeToString(paperpb.VoterID) + ": " + err.Error())
			}

			voterdb := new(db.Voter)
			voterdb.HashID = voterpb.HashID
			if err := db.SelectVoter(ctx, voterdb); err == nil {
				log.Println(hex.EncodeToString(paperpb.VoterID) + " already exists")
				continue
			}

			voterdb.PublicKey = voterpb.PublicKey
			voterdb.RandomBytes = voterpb.RandomBytes
			voterdb.Timestamp = voterpb.Timestamp
			voterdb.Voted = voterpb.Voted
			if err := db.InsertVoter(ctx, voterdb); err != nil {
				panic("panic: " + hex.EncodeToString(paperpb.VoterID) + ": " + err.Error())
			}
		}

		currentBundle = bundlepb.Prev
	}

	log.Println("all brought")

	res, err = http.Get(fmt.Sprintf("http://%s:9999/candidate", uri))
	if err != nil {
		panic("panic: " + err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("panic: cannot get candidate list" + res.Status)
	}

	candidates := new(candidate.List)
	buffers, err := io.ReadAll(res.Body)
	if err != nil {
		panic("panic: " + err.Error())
	}
	if err := proto.Unmarshal(buffers, candidates); err != nil {
		panic("panic: " + err.Error())
	}

	getVotes := map[string]int{}
	for _, candidate := range candidates.Names {
		getVotes[candidate] = 0
	}

	bundledb := new(db.Bundle)
	paperdb := new(db.Paper)
	voterdb := new(db.Voter)

	currentBundle = make([]byte, len(lastBundleHash))
	copy(currentBundle, lastBundleHash)
	for {
		b := true
		for _, v := range currentBundle {
			if v != 0 {
				b = false
				break
			}
		}
		if b {
			break
		}

		fmt.Println("bundle: ", hex.EncodeToString(currentBundle))

		bundledb.Hash = currentBundle
		if err := db.SelectBundle(ctx, bundledb); err != nil {
			panic("panic: " + hex.EncodeToString(currentBundle) + ": " + err.Error())
		}

		list := make([]string, 8)
		sha := sha512.New()
		sha.Write(bundledb.Prev)
		for i, subHash := range bundledb.SubHashes {
			paperdb.Hash = make([]byte, len(subHash))
			copy(paperdb.Hash, subHash)
			fmt.Println("subhash: ", hex.EncodeToString(paperdb.Hash))
			if err := db.SelectPaper(ctx, paperdb); err != nil {
				panic("panic: " + hex.EncodeToString(subHash) + ": " + err.Error())
			}

			insha := sha256.New()
			insha.Write(paperdb.RandomBytes)
			insha.Write(paperdb.VoterID)
			insha.Write(paperdb.Timestamp)
			insha.Write(paperdb.Message)
			hashed := insha.Sum(nil)
			if !bytes.Equal(hashed, paperdb.Hash) {
				panic("panic: " + hex.EncodeToString(subHash) + " " + hex.EncodeToString(hashed) + ": hash mismatch")
			}

			voterdb.HashID = paperdb.VoterID
			if err := db.SelectVoter(ctx, voterdb); err != nil {
				panic("panic: " + hex.EncodeToString(paperdb.VoterID) + ": " + err.Error())
			}

			pubkey, err := auth.DeserializePublicKey(voterdb.PublicKey)
			if err != nil {
				panic("panic: " + hex.EncodeToString(paperdb.VoterID) + " : " + hex.EncodeToString(voterdb.PublicKey) + " : " + err.Error())
			}

			r, s, err := auth.DeserializeSignature(paperdb.Signature)
			if err != nil {
				panic("panic: " + hex.EncodeToString(paperdb.VoterID) + ": " + err.Error())
			}

			if !ecdsa.Verify(pubkey, paperdb.Hash, r, s) {
				panic("panic: " + hex.EncodeToString(paperdb.VoterID) + ": signature mismatch")
			}

			sha.Write(paperdb.Hash)
			list[i] = string(paperdb.Message)
		}

		hashed := sha.Sum(nil)
		hashedTemp := sha512.Sum512(hashed)
		hashed = hashedTemp[:]
		// if !bytes.Equal(bundledb.Hash, hashed) {
		// 	panic("panic: " + hex.EncodeToString(bundledb.Hash) + " : " + hex.EncodeToString(hashed) + ": hash mismatch")
		// }

		for _, candidate := range list {
			getVotes[candidate]++
		}

		currentBundle = bundledb.Prev
	}

	log.Println("all voted")

	fmt.Println(getVotes)

	f, err := os.Create("result.txt")
	if err != nil {
		panic("panic: " + err.Error())
	}
	defer f.Close()

	time.Sleep(1 * time.Second)
}
