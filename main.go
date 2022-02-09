package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"

	"time"

	"github.com/valyala/fasthttp"
)

type Config struct {
	ServerURL  string
	ListenPort string
	Host       string
}

var cfg Config

func handler(ctx *fasthttp.RequestCtx) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	ctx.Request.CopyTo(req)
	ctx.PostArgs().CopyTo(req.PostArgs())
	ctx.Request.Header.CopyTo(&req.Header)

	req.SetHost(cfg.Host)

	err := client.Do(req, resp)
	if err != nil {
		log.Println(err.Error())
		return
	}

	resp.CopyTo(&ctx.Response)
	ctx.SetStatusCode(resp.StatusCode())
}

func main() {
	buf, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = json.Unmarshal(buf, &cfg)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Fatal(fasthttp.ListenAndServe(cfg.ListenPort, handler))
}
