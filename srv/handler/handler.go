package handler

import (
	"github.com/kazouphq/crawler/srv/proto"
	"golang.org/x/net/context"
)

type Crawler struct{}

func (c *Crawler) Start(ctx context.Context, req *go_micro_srv_crawler.StartRequest, res *go_micro_srv_crawler.StartResponse) error {
	return nil
}
