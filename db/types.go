package db

import "github.com/edgedb/edgedb-go"

type ID struct {
	ID edgedb.UUID `edgedb:"id"`
}

type Candidate struct {
	ID       edgedb.UUID `edgedb:"id"`
	Name     string      `edgedb:"name"`
	Markdown string      `edgedb:"markdown"`
}

type Voter struct {
	ID            edgedb.UUID `edgedb:"id"`
	HashID        []byte      `edgedb:"hash_id"`
	PublicKey     []byte      `edgedb:"public_key"`
	RandomBytes   []byte      `edgedb:"random_bytes"`
	Timestamp     []byte      `edgedb:"timestamp"`
	Voted         bool        `edgedb:"voted"`
	IsCandidate   bool        `edgedb:"is_candidate"`
	RepliedSurvey bool        `edgedb:"replied_survey"`
}

type Enrolled struct {
	ID            edgedb.UUID `edgedb:"id"`
	Name          string      `edgedb:"name"`
	PrivateNumber string      `edgedb:"private_number"`
	PhoneNumber   string      `edgedb:"phone_number"`
}

type Paper struct {
	ID          edgedb.UUID `edgedb:"id"`
	VoterID     []byte      `edgedb:"voter_id"`
	Timestamp   []byte      `edgedb:"timestamp"`
	Message     []byte      `edgedb:"message"`
	RandomBytes []byte      `edgedb:"random_bytes"`
	Hash        []byte      `edgedb:"hash"`
	Signature   []byte      `edgedb:"signature"`
}

type Bundle struct {
	ID        edgedb.UUID `edgedb:"id"`
	Hash      []byte      `edgedb:"hash"`
	Prev      []byte      `edgedb:"prev"`
	SubHashes [][]byte    `edgedb:"sub_hashes"`
}

type Data struct {
	ID        edgedb.UUID `edgedb:"id"`
	Candidate string      `edgedb:"candidate"`
	Gender    bool        `edgedb:"gender"`
	Age       int16       `edgedb:"age"`
	Region    string      `edgedb:"region"`
	Job       string      `edgedb:"job"`
	Education string      `edgedb:"education"`
	Married   bool        `edgedb:"married"`
	Divorced  bool        `edgedb:"divorced"`
	HasCar    bool        `edgedb:"has_car"`
	HouseType string      `edgedb:"house_type"`
	Salary    int16       `edgedb:"salary"`
	HasDebt   bool        `edgedb:"has_debt"`
	Ideology  string      `edgedb:"ideology"`
	Number    int64       `edgedb:"number"`
}
