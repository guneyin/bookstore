package dto

import "github.com/guneyin/bookstore/service/general"

type StatusResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Env     string `json:"env"`
	Uptime  string `json:"uptime"`
}

func (sr *StatusResponse) FromEntity(e general.Status) {
	sr.Status = string(e.Status)
	sr.Version = e.Version
	sr.Env = string(e.Env)
	sr.Uptime = e.Uptime
}
