package db

import (
	"context"
	"log"
	"strings"

	"github.com/edgedb/edgedb-go"
)

var Voters *edgedb.Pool
var Papers *edgedb.Pool
var GlobalCtx context.Context = context.Background()

func Init() {
	opt := edgedb.Options{
		MinConns: 10,
		MaxConns: 20,
	}
	voters, err := edgedb.ConnectDSN(GlobalCtx, "edgedb://merak:test@voters:16419", opt)
	if err != nil {
		log.Fatal(err)
	}
	papers, err := edgedb.ConnectDSN(GlobalCtx, "edgedb://merak:test@papers:16420", opt)
	if err != nil {
		log.Fatal(err)
	}

	Voters = voters
	Papers = papers

	tempIDs := new([]ID)

	if err := Voters.Query(GlobalCtx, `select Candidate`, tempIDs); err != nil {
		cmd := `
		create type Candidate {
			create required property name -> str;
			create required property markdown -> str;

			create index on (.name);
		}
		`
		if err := Voters.Execute(GlobalCtx, cmd); err != nil {
			if !strings.Contains(err.Error(), "already") {
				log.Fatal(err)
			}
		}
	}

	if err := Voters.Query(GlobalCtx, `select Voter`, tempIDs); err != nil {
		cmd := `
		create type Voter {
			create required property hash_id -> bytes;
			create required property public_key -> bytes;
			create property random_bytes -> bytes;
			create property timestamp -> bytes;
			create required property voted -> bool;
			create required property is_candidate -> bool;
			create required property replied_survey -> bool;

			create index on (.hash_id);
		}
		`
		if err := Voters.Execute(GlobalCtx, cmd); err != nil {
			if !strings.Contains(err.Error(), "already") {
				log.Fatal(err)
			}
		}
	}

	if err := Voters.Query(GlobalCtx, `select Enrolled`, tempIDs); err != nil {
		cmd := `
		create type Enrolled {
			create required property name -> str;
			create required property private_number -> str;
			create required property phone_number -> str;

			create index on (.name);
		}
		`
		if err := Voters.Execute(GlobalCtx, cmd); err != nil {
			if !strings.Contains(err.Error(), "already") {
				log.Fatal(err)
			}
		}
	}

	if err := Papers.Query(GlobalCtx, `select Paper`, tempIDs); err != nil {
		cmd := `
		create type Paper {
			create required property voter_id -> bytes;
			create required property timestamp -> bytes;
			create required property message -> bytes;
			create required property random_bytes -> bytes;
			create required property hash -> bytes;
			create required property signature -> bytes;

			create index on (.hash);
		}
		`
		if err := Papers.Execute(GlobalCtx, cmd); err != nil {
			if !strings.Contains(err.Error(), "already") {
				log.Fatal(err)
			}
		}
	}

	if err := Papers.Query(GlobalCtx, `select Bundle`, tempIDs); err != nil {
		cmd := `
		create type Bundle {
			create required property hash -> bytes;
			create required property prev -> bytes;
			create required property sub_hashes -> array<bytes>;

			create index on (.hash)
		}
		`
		if err := Papers.Execute(GlobalCtx, cmd); err != nil {
			if !strings.Contains(err.Error(), "already") {
				log.Fatal(err)
			}
		}
	}

	if err := Voters.Query(GlobalCtx, `select Data`, tempIDs); err != nil {
		cmd := `
		create type Data {
			create required property candidate -> str;
			create property gender -> bool;
			create property age -> int16;
			create property region -> str;
			create property job -> str;
			create property education -> str;
			create property married -> bool;
			create property divorced -> bool;
			create property has_car -> bool;
			create property house_type -> str;
			create property salary -> int16;
			create property has_debt -> bool;
			create property ideology -> str;
		}
		`
		if err := Voters.Execute(GlobalCtx, cmd); err != nil {
			if !strings.Contains(err.Error(), "already") {
				log.Fatal(err)
			}
		}
	}
}
