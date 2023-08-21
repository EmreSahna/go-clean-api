package service

import (
	"go-clean-api/config"
	"go-clean-api/responses"
	"time"
)

type HealthCheckServiceInterface interface {
	Healthcheck(*responses.Response)
}

type HealthCheckService struct {
	cfg *config.Config
}

func (s *HealthCheckService) Healthcheck(response *responses.Response) {
	result := &responses.HealthCheckResponse{
		Status:  "OK",
		Env:     s.cfg.Env,
		Version: s.cfg.Version,
		Time:    time.Now(),
	}

	response.SetStatusCode(200).SetBody(result)
	return
}
