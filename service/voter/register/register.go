package register

import (
	"alcor/db"
	"alcor/model/register"
	"alcor/service/security"
	"alcor/worker/certcodes"
	"context"
	"crypto/sha512"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lemon-mint/libuseful"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func PhoneService(c *fiber.Ctx) error {
	id, body, err := security.Decapsulate(c.Body())
	if err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusBadRequest)
	}

	phone := new(register.Phone)
	if err := proto.Unmarshal(body, phone); err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusBadRequest)
	}

	e := new(db.Enrolled)
	e.Name = phone.Name
	e.PhoneNumber = phone.Number
	e.PrivateNumber = phone.RRN
	if !db.NotExistEnrolled(context.Background(), e) {
		return c.SendStatus(fasthttp.StatusAlreadyReported)
	}

	// TODO: request phone auth
	code := []byte(fmt.Sprintf("%05d", libuseful.FastRand32n(99998)+1))
	fmt.Println(string(code))

	cache := new(register.Cache)
	cache.Code = code
	cache.Gender = phone.Gender
	cache.Name = phone.Name
	cache.Number = phone.Number
	cache.RRN = phone.RRN
	data, err := proto.Marshal(cache)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusConflict)
	}

	if err := certcodes.Set(id, data); err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusConflict)
	}
	return c.SendStatus(fasthttp.StatusOK)
}

func CertService(c *fiber.Ctx) error {
	id, body, err := security.Decapsulate(c.Body())
	if err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusBadRequest)
	}

	result := new(register.Result)
	cert := new(register.CertCode)
	if err := proto.Unmarshal(body, cert); err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		result.Error = err.Error()
	}

	data, err := certcodes.Get(id)
	if result.Error == "" {
		if err != nil {
			log.Println(err)
			c.SendStatus(fasthttp.StatusNotFound)
			result.Error = err.Error()
		}
	}

	cache := new(register.Cache)
	if result.Error == "" {
		if err := proto.Unmarshal(data, cache); err != nil {
			log.Println(err)
			c.SendStatus(fasthttp.StatusConflict)
			result.Error = err.Error()
		}
	}

	if result.Error == "" {
		// if !bytes.Equal(cache.Code, cert.Code) {
		// 	c.SendStatus(fasthttp.StatusBadRequest)
		// 	result.Error = "cert code is not valid"
		// } else {
		// 	hashed := sha512.Sum512(cert.PublicKey)
		// 	result.ID = hashed[:]
		// 	c.SendStatus(fasthttp.StatusOK)
		// }
		hashed := sha512.Sum512(cert.PublicKey)
		result.ID = hashed[:]
		c.SendStatus(fasthttp.StatusOK)
	}

	voter := new(db.Voter)
	if result.Error == "" {
		voter.HashID = result.ID
		voter.IsCandidate = false
		voter.PublicKey = cert.PublicKey
		voter.Voted = false
		if err := db.InsertVoter(context.Background(), voter); err != nil {
			log.Print(err)
			result.Error = err.Error()
		}
	}

	if result.Error == "" {
		enrolled := new(db.Enrolled)
		enrolled.Name = cache.Name
		enrolled.PhoneNumber = cache.Number
		enrolled.PrivateNumber = cache.RRN
		if err := db.InsertEnrolled(context.Background(), enrolled); err != nil {
			log.Print(err)
			result.Error = err.Error()
		}
	}

	data, err = proto.Marshal(result)
	if err != nil {
		log.Println(err)
	}

	data, err = security.Encapsulate(id, data)
	if err != nil {
		log.Println(err)
	}

	return c.Send(data)
}
