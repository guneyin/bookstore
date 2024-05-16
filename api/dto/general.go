package dto

import "github.com/guneyin/bookstore/service/general"

type StatusResponse struct {
	Status      string `json:"status"`
	VersionInfo struct {
		Version    string `json:"version"`
		CommitHash string `json:"commitHash"`
		BuildTime  string `json:"buildTime"`
	} `json:"version"`
	Env    string `json:"env"`
	Uptime string `json:"uptime"`
}

func (sr *StatusResponse) FromEntity(e general.Status) {
	sr.Status = string(e.Status)
	sr.VersionInfo.Version = e.Version.Version
	sr.VersionInfo.CommitHash = e.Version.CommitHash
	sr.VersionInfo.BuildTime = e.Version.BuildTime
	sr.Env = string(e.Env)
	sr.Uptime = e.Uptime
}
