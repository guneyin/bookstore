package general

import "github.com/guneyin/bookstore/common"

type (
	ServiceStatus string
	Env           string
)

const (
	ServiceStatusOnline      ServiceStatus = "online"
	ServiceStatusMaintenance ServiceStatus = "maintenance"

	EnvProduction Env = "production"
	EnvStaging    Env = "staging"
)

type Status struct {
	Status  ServiceStatus
	Version *common.VersionInfo
	Env     Env
	Uptime  string
}
