package http

import "time"

type Config struct {
	Port    string        `json:"port"`
	Timeout time.Duration `json:"timeout"`
	Swagger bool          `json:"swagger"`
}
