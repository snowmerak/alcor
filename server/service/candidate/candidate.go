package candidate

import (
	"alcor/db"
	"context"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/snowmerak/jumper"
	"google.golang.org/protobuf/proto"
)

type CandidateError struct {
	errorData string
}

func (v CandidateError) Error() string {
	return v.errorData
}

func NewCandidateError(data string) CandidateError {
	return CandidateError{errorData: data}
}

func GetList(c *fiber.Ctx) error {
	candidates := new([]db.Candidate)
	reply := new(CandidateList)
	if err := db.SelectCandidates(context.Background(), candidates); err != nil {
		jumper.Offer(NewCandidateError(err.Error()))
		reply.Error = err.Error()
		data, err := proto.Marshal(reply)
		if err != nil {
			jumper.Offer(NewCandidateError(err.Error()))
		}
		return c.Send(data)
	}
	for _, v := range *candidates {
		reply.Names = append(reply.Names, v.Name)
	}
	reply.Error = ""
	data, err := proto.Marshal(reply)
	if err != nil {
		jumper.Offer(NewCandidateError(err.Error()))
	}
	return c.Send(data)
}

func Get(c *fiber.Ctx) error {
	name, err := base64.RawURLEncoding.DecodeString(c.Params("name"))
	if err != nil {
		c.Send([]byte(err.Error()))
	}
	candidate := new(db.Candidate)
	candidate.Name = string(name)
	result := new(Candidate)
	if err := db.SelectCandidate(context.Background(), candidate); err != nil {
		jumper.Offer(NewCandidateError(err.Error()))
		result.Error = err.Error()
		data, err := proto.Marshal(result)
		if err != nil {
			jumper.Offer(NewCandidateError(err.Error()))
		}
		return c.Send(data)
	}
	result.Name = candidate.Name
	result.Markdown = candidate.Markdown
	data, err := proto.Marshal(result)
	if err != nil {
		jumper.Offer(NewCandidateError(err.Error()))
	}
	return c.Send(data)
}
