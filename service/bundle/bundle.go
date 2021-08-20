package bundle

import (
	"alcor/db"
	"alcor/model/info"
	"alcor/worker/bundle_cache"
	"context"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func GetLast(c *fiber.Ctx) error {
	data, err := bundle_cache.Get("last")
	if err != nil {
		return c.SendStatus(fasthttp.StatusNotFound)
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}

func Get(c *fiber.Ctx) error {
	data, err := base64.RawURLEncoding.DecodeString(c.Params("id", ""))
	if err != nil {
		return c.SendStatus(fasthttp.StatusBadRequest)
	}

	bundle := new(db.Bundle)
	bundle.Hash = data
	if err != db.SelectBundle(context.Background(), bundle) {
		return c.SendStatus(fasthttp.StatusNotFound)
	}

	bundlePB := new(info.Bundle)
	bundlePB.Hash = bundle.Hash
	bundlePB.Prev = bundle.Prev
	bundlePB.SubHashes = bundle.SubHashes
	data, err = proto.Marshal(bundlePB)
	if err != nil {
		return c.SendStatus(fasthttp.StatusConflict)
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}
