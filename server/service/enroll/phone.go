package enroll

import (
	"alcor/db"
	"alcor/service/enroll/certpb"
	"alcor/service/wst"
	"bytes"
	"context"
	"crypto/sha512"
	"fmt"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/snowmerak/jumper"
	"github.com/snowmerak/wrapper/buffer"
	"google.golang.org/protobuf/proto"
)

type CertError struct {
	errorData string
}

func (c CertError) Error() string {
	return c.errorData
}

func NewCertError(data string) CertError {
	return CertError{errorData: data}
}

/*
<- phone info
-> OK
<- cert code
-> OK
<- public key
-> ID
*/
func PhoneCertSocket(c *fiber.Ctx) error {
	return wst.Listen(c.Context(), func(conn *websocket.Conn) {
		buf := bytes.NewBuffer(nil)
		reply := new(certpb.Result)

		aead, err := wst.Handshake(conn, buf)
		if err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}
		secret := make([]byte, buf.Len())
		if err := buffer.BytesBuffer2Bytes(secret, buf); err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}

		buf.Reset()
		if err := wst.ReceiveBinaryMessage(conn, aead, secret, buf); err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}

		errmsg := ""

		phone := new(certpb.Phone)
		if errmsg == "" {
			if err := proto.Unmarshal(buf.Bytes(), phone); err != nil {
				jumper.Offer(NewCertError(err.Error()))
				errmsg = err.Error()
			}
		}

		enrolled := new(db.Enrolled)
		if errmsg == "" {
			enrolled.Name = phone.Name
			enrolled.PhoneNumber = phone.Number
			enrolled.PrivateNumber = phone.RRN
			if !db.NotExistEnrolled(context.Background(), enrolled) {
				jumper.Offer(NewCertError("already enrolled"))
				errmsg = "already enrolled"
			}
		}

		//TODO::connect phone auth
		reply.Ok = true
		reply.Error = errmsg
		if errmsg != "" {
			reply.Ok = false
		}
		data, err := proto.Marshal(reply)
		if err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}
		if err := wst.SendBinaryMessage(conn, aead, secret, data); err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}
		errmsg = ""

		buf.Reset()
		if err := wst.ReceiveBinaryMessage(conn, aead, secret, buf); err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}

		//TODO::check cert code
		reply.Ok = true
		reply.Error = ""
		data, err = proto.Marshal(reply)
		if err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}
		if err := wst.SendBinaryMessage(conn, aead, secret, data); err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}

		buf.Reset()
		if err := wst.ReceiveBinaryMessage(conn, aead, secret, buf); err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}
		publicKey := make([]byte, len(buf.Bytes()))
		if err := buffer.BytesBuffer2Bytes(publicKey, buf); err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}

		ctx := context.Background()
		voter := new(db.Voter)
		voter.PublicKey = publicKey
		hashed := sha512.Sum512(publicKey)
		voter.HashID = hashed[:]
		voter.Voted = false
		if err := db.InsertVoter(ctx, voter); err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}
		if err := db.InsertEnrolled(ctx, enrolled); err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}

		result := new(certpb.ID)
		result.ID = voter.HashID
		result.Error = ""
		data, err = proto.Marshal(result)
		if err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}
		if err := wst.SendBinaryMessage(conn, aead, secret, data); err != nil {
			jumper.Offer(NewCertError(err.Error()))
			return
		}

		jumper.Offer(NewCertError(fmt.Sprintf("%v is added", voter.HashID)))
	})
}
