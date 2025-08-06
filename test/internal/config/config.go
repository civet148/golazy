package config

import "time"

type Orm struct {
	DSN   string `yaml:"DSN" json:"dsn"`     // database source name
	Debug bool   `yaml:"Debug" json:"debug"` // open or close debug log
	Max   int    `yaml:"Max" json:"max"`     // max database connections
	Idle  int    `yaml:"Idle" json:"idle"`   // idle database connections
}

type Config struct {
	Name    string        `yaml:"Name" json:"name"`       // service name
	Host    string        `yaml:"Host" json:"host"`       // service host
	Port    string        `yaml:"Port" json:"port"`       // service port
	Mode    string        `yaml:"Mode" json:"mode"`       // service mode (dev/test/prod)
	Timeout time.Duration `yaml:"Timeout" json:"timeout"` // service timeout duration
	Orm     Orm           `yaml:"Orm" json:"orm"`         // database ORM config
}
