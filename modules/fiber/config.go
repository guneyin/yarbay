package fiber

import "time"

type Config struct {
	AppName string        `json:"appName"`
	Port    string        `json:"port"`
	Timeout time.Duration `json:"timeout"`
	Swagger bool          `json:"swagger"`
}
