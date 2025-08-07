package config

import "time"

type Orm struct {
	DSN       string `yaml:"DSN" json:"dsn"`              // database source name
	Debug     bool   `yaml:"Debug" json:"debug"`          // open or close debug log
	MaxConns  int    `yaml:"MaxConns" json:"max_conns"`   // max database connections
	IdleConns int    `yaml:"IdleConns" json:"idle_conns"` // idle database connections
	NodeId    int64  `yaml:"NodeId" json:"node_id"`       // snowflake node id
}

type Config struct {
	Name    string        `yaml:"Name" json:"name"`       // service name
	Host    string        `yaml:"Host" json:"host"`       // service host
	Port    string        `yaml:"Port" json:"port"`       // service port
	Mode    string        `yaml:"Mode" json:"mode"`       // service mode (dev/test/prod)
	Timeout time.Duration `yaml:"Timeout" json:"timeout"` // service timeout duration
	Orm     Orm           `yaml:"Orm" json:"orm"`         // database ORM config
}
