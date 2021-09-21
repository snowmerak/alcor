package main

import (
	"alcor/db"
	"alcor/model/candidate"
	"alcor/model/info"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/snowmerak/wrapper/auth"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Print("input server uri: ")
	var uri string
	fmt.Scanln(&uri)
	res, err := http.Get(fmt.Sprintf("http://%s:9999/bundle", uri))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(res.Status)
	}
	currentBundle, err := io.ReadAll(res.Body)
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
		if b {
			break
		}

		bundleID := base64.RawURLEncoding.EncodeToString(currentBundle)
		log.Println("current bundle hash: " + bundleID)

		bundledb := new(db.Bundle)
		bundledb.Hash = currentBundle
		if err := db.SelectBundle(ctx, bundledb); err == nil {
			log.Println(bundleID + " already exists")
			break
		}

		res, err = http.Get(fmt.Sprintf("http://%s:9999/bundle/%s", uri, bundleID))
		if err != nil {
			panic("panic: " + bundleID + ": " + err.Error())
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			panic("panic: " + bundleID + ": " + res.Status)
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

		for _, subHash := range bundlepb.SubHashes {
			paperid := base64.RawURLEncoding.EncodeToString(subHash)
			log.Println("subhash: " + paperid)

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
			paperdb.RandomBytes = paperpb.Message
			paperdb.Signature = paperpb.RandomBytes
			paperdb.Timestamp = paperpb.Timestamp
			paperdb.VoterID = paperpb.VoterID
			if err := db.InsertPaper(ctx, paperdb); err != nil {
				panic("panic: " + paperid + ": " + err.Error())
			}

			voterpb := new(info.Voter)
			voterpb.HashID = paperpb.VoterID
			res, err := http.Get(fmt.Sprintf("http://%s:9999/voter/info/%s", uri, paperpb.VoterID))
			if err != nil {
				panic("panic: " + hex.EncodeToString(paperpb.VoterID) + ": " + err.Error())
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				panic("panic: " + hex.EncodeToString(paperpb.VoterID) + ": " + res.Status)
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

	copy(currentBundle, lastBundleHash)
	for {
		b := false
		for _, v := range currentBundle {
			if v == 0 {
				b = true
				break
			}
		}
		if b {
			break
		}

		bundledb.Hash = currentBundle
		if err := db.SelectBundle(ctx, bundledb); err != nil {
			panic("panic: " + hex.EncodeToString(currentBundle) + ": " + err.Error())
		}

		list := make([]string, 0, 8)
		sha := sha512.New()
		for _, subHash := range bundledb.SubHashes {
			paperdb.Hash = make([]byte, len(subHash))
			copy(paperdb.Hash, subHash)
			if err := db.SelectPaper(ctx, paperdb); err != nil {
				panic("panic: " + hex.EncodeToString(subHash) + ": " + err.Error())
			}

			sha := sha256.New()
			sha.Write(paperdb.RandomBytes)
			sha.Write(paperdb.VoterID)
			sha.Write(paperdb.Timestamp)
			sha.Write(paperdb.Message)
			hashed := sha.Sum(nil)
			if !bytes.Equal(hashed, paperdb.Hash) {
				panic("panic: " + hex.EncodeToString(subHash) + ": hash mismatch")
			}

			voterdb.HashID = paperdb.VoterID
			if err := db.SelectVoter(ctx, voterdb); err != nil {
				panic("panic: " + hex.EncodeToString(paperdb.VoterID) + ": " + err.Error())
			}

			pubkey, err := auth.DeserializePublicKey(voterdb.PublicKey)
			if err != nil {
				panic("panic: " + hex.EncodeToString(paperdb.VoterID) + ": " + err.Error())
			}

			r, s, err := auth.DeserializeSignature(paperdb.Signature)
			if err != nil {
				panic("panic: " + hex.EncodeToString(paperdb.VoterID) + ": " + err.Error())
			}

			if !ecdsa.Verify(pubkey, paperdb.Hash, r, s) {
				panic("panic: " + hex.EncodeToString(paperdb.VoterID) + ": signature mismatch")
			}

			sha.Write(paperdb.Hash)
			list = append(list, string(paperdb.Message))
		}

		hash := sha.Sum(nil)
		if !bytes.Equal(bundledb.Hash, hash) {
			panic("panic: " + hex.EncodeToString(bundledb.Hash) + ": hash mismatch")
		}

		for _, candidate := range list {
			getVotes[candidate]++
		}
	}

	log.Println("all voted")

	f, err := os.Create("result.txt")
	if err != nil {
		panic("panic: " + err.Error())
	}
	defer f.Close()

	w := csv.NewWriter(f)
	for k, v := range getVotes {
		if err := w.Write([]string{k, strconv.Itoa(v)}); err != nil {
			panic("panic: " + err.Error())
		}
	}
	w.Flush()
}
