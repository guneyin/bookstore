package general

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
	Version string
	Env     Env
	Uptime  string
}
