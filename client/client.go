package client

import "github.com/go-resty/resty/v2"

type Client struct {
	httpClient resty.Client
}
