package general

import (
	"github.com/guneyin/bookstore/common"
	"time"
)

type Service struct{}

func New() *Service {
	return &Service{}
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
