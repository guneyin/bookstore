package general

import (
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/config"
	"time"
)

type Service struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Status() Status {
	uptime := time.Now().Sub(common.GetLastRun())

	return Status{
		Status:  ServiceStatusOnline,
		Version: common.GetVersion(),
		Env:     EnvStaging,
		Uptime:  uptime.String(),
	}
}
