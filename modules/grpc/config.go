package grpc

import "time"

type Config struct {
	Port    string        `json:"port"`
	Timeout time.Duration `json:"timeout"`
}
