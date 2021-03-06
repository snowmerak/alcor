package db

import (
	"alcor/worker/rdb"
	"context"
	"sync"
)

func SelectCandidate(ctx context.Context, candidate *Candidate) error {
	cmd := `
	select Candidate {
		name,
		markdown,
	}
	filter Candidate.name = <str>$0
	limit 1
	`
	return Voters.QueryOne(ctx, cmd, candidate, candidate.Name)
}

func UpdateCandidate(ctx context.Context, candidate *Candidate) error {
	id := new([]ID)
	cmd := `
	update Candidate
	filter Candidate.name = <str>$0
	set {
		markdown := <str>$1
	}
	`
	return Voters.Query(ctx, cmd, id, candidate.Name, candidate.Markdown)
}

func DeleteCandidate(ctx context.Context, candidate *Candidate) error {
	cmd := `
	delete Candidate
	filter Candidate.name = <str>$0
	limit 1
	`
	return Voters.QueryOne(ctx, cmd, candidate, candidate.Name)
}

func SelectCandidates(ctx context.Context, candidates *[]Candidate) error {
	cmd := `
	select Candidate {
		name,
	}
	`
	return Voters.Query(ctx, cmd, candidates)
}

func InsertCandidate(ctx context.Context, candidate *Candidate) error {
	id := new(ID)
	cmd := `
	insert Candidate {
		name := <str>$0,
		markdown := <str>$1,
	}
	`
	return Voters.QueryOne(ctx, cmd, id, candidate.Name, candidate.Markdown)
}

func SelectVoter(ctx context.Context, voter *Voter) error {
	cmd := `
	select Voter {
		hash_id,
		public_key,
		random_bytes,
		timestamp,
		voted,
		is_candidate,
		replied_survey,
	}
	filter Voter.hash_id = <bytes>$0
	limit 1
	`
	return Voters.QueryOne(ctx, cmd, voter, voter.HashID)
}

func InsertVoter(ctx context.Context, voter *Voter) error {
	id := new(ID)
	cmd := `
	insert Voter {
		hash_id := <bytes>$0,
		public_key := <bytes>$1,
		voted := <bool>$2,
		is_candidate := false,
		replied_survey := false,
	}
	`
	return Voters.QueryOne(ctx, cmd, id, voter.HashID, voter.PublicKey, voter.Voted)
}

func UpdateVoter(ctx context.Context, voter *Voter) error {
	id := new([]ID)
	cmd := `
	update Voter
	filter .hash_id = <bytes>$0
	set {
		random_bytes := <bytes>$1,
		timestamp := <bytes>$2,
		voted := <bool>$3,
	}
	`
	return Voters.Query(ctx, cmd, id, voter.HashID, voter.RandomBytes, voter.Timestamp, voter.Voted)
}

func UpdateVoterRepliedSurvey(ctx context.Context, voter *Voter) error {
	id := new([]ID)
	cmd := `
	update Voter
	filter .hash_id = <bytes>$0
	set {
		replied_survey := true,
	}
	`
	return Voters.Query(ctx, cmd, id, voter.HashID)
}

func ChangeVoter2Candidate(ctx context.Context, voter *Voter) error {
	id := new([]ID)
	cmd := `
	update Voter
	filter .hash_id = <bytes>$0
	set {
		is_candidate := <bool>$1
	}
	`
	return Voters.Query(ctx, cmd, id, voter.HashID, true)
}

func NotExistEnrolled(ctx context.Context, enrolled *Enrolled) bool {
	cmd := `
	select Enrolled {
		name,
		private_number,
		phone_number,
	}
	filter Enrolled.name = <str>$0 and Enrolled.private_number = <str>$1 and Enrolled.phone_number = <str>$2
	limit 1
	`
	if err := Voters.QueryOne(ctx, cmd, enrolled, enrolled.Name, enrolled.PrivateNumber, enrolled.Name); err != nil {
		return true
	}
	return false
}

func InsertEnrolled(ctx context.Context, enrolled *Enrolled) error {
	id := new(ID)
	cmd := `
	insert Enrolled {
		name := <str>$0,
		private_number := <str>$1,
		phone_number := <str>$2,
	}
	`
	return Voters.QueryOne(ctx, cmd, id, enrolled.Name, enrolled.PrivateNumber, enrolled.PhoneNumber)
}

func SelectPaper(ctx context.Context, paper *Paper) error {
	cmd := `
	select Paper {
		voter_id,
		timestamp,
		message,
		random_bytes,
		signature,
		hash,
	}
	filter Paper.hash = <bytes>$0
	limit 1
	`
	return Papers.QueryOne(ctx, cmd, paper, paper.Hash)
}

func InsertPaper(ctx context.Context, paper *Paper) error {
	id := new(ID)
	cmd := `
	insert Paper {
		voter_id := <bytes>$0,
		timestamp := <bytes>$1,
		message := <bytes>$2,
		random_bytes := <bytes>$3,
		hash := <bytes>$4,
		signature := <bytes>$5,
	}
	`
	return Papers.QueryOne(ctx, cmd, id, paper.VoterID, paper.Timestamp, paper.Message, paper.RandomBytes, paper.Hash, paper.Signature)
}

func SelectBundle(ctx context.Context, bundle *Bundle) error {
	cmd := `
	select Bundle {
		hash,
		prev,
		sub_hashes,
	}
	filter Bundle.hash = <bytes>$0
	limit 1
	`
	return Papers.QueryOne(ctx, cmd, bundle, bundle.Hash)
}

func InsertBundle(ctx context.Context, bundle *Bundle) error {
	id := new(ID)
	cmd := `
	insert Bundle {
		hash := <bytes>$0,
		prev := <bytes>$1,
		sub_hashes := <array<bytes>>$2,
	}
	`
	return Papers.QueryOne(ctx, cmd, id, bundle.Hash, bundle.Prev, bundle.SubHashes)
}

func SelectData(ctx context.Context, data *Data) error {
	cmd := `
	select Data {
		candidate,
		gender,
		age,
		region,
		job,
		education,
		married,
		divorced,
		has_car,
		house_type,
		salary,
		has_debt,
		ideology,
	}
	filter Data.number = <int64>$0
	limit 1
	`
	return Voters.QueryOne(ctx, cmd, data, data.Number)
}

var insertDataLock = sync.Mutex{}

func InsertData(ctx context.Context, data *Data) error {
	insertDataLock.Lock()
	defer insertDataLock.Unlock()
	id := new(ID)
	number, err := rdb.GetSurveyIndex()
	if err != nil {
		number = 0
	}
	cmd := `
	insert Data {
		number := <int64>$13,
		candidate := <str>$0,
		gender := <bool>$1,
		age := <int32>$2,
		region := <str>$3,
		job := <str>$4,
		education := <str>$5,
		married := <bool>$6,
		divorced := <bool>$7,
		has_car := <bool>$8,
		house_type := <str>$9,
		salary := <int16>$10,
		has_debt := <bool>$11,
		ideology := <str>$12,
	}
	`
	if err := Voters.QueryOne(ctx, cmd, id, data.Candidate, data.Gender, data.Age, data.Region, data.Job, data.Education, data.Married, data.Divorced, data.HasCar, data.HouseType, data.Salary, data.HasDebt, data.Ideology, int64(number)); err != nil {
		return err
	}
	number++
	if err := rdb.SetSurveyIndex(number); err != nil {
		return err
	}
	return nil
}
