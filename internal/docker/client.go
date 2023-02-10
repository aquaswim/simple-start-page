package docker

import (
	"github.com/valyala/fasthttp"
	"net"
)

type client struct {
	httpClient *fasthttp.Client
}

type Client interface {
}

func NewDockerClient() Client {
	httpClient := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return net.Dial("unix", "")
		},
	}
	return &client{httpClient: httpClient}
}
