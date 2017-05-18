// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "cellar": Client
//
// Command:
// $ goagen
// --design=github.com/goadesign/goa-cellar/design
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the cellar service client.
type Client struct {
	*goaclient.Client
	Encoder *goa.HTTPEncoder
	Decoder *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	client.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")
	client.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	client.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}
