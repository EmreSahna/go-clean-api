package responses

import (
	"time"
)

type HealthCheckResponse struct {
	Status  string    `json:"status"`
	Env     string    `json:"env"`
	Version string    `json:"version"`
	Time    time.Time `json:"time"`
}
