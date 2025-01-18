package config

import "time"

type Orm struct {
	DSN     string  `yaml:"DSN" json:"dsn"`
    Debug   bool    `yaml:"Debug" json:"debug"`
    Max     int     `yaml:"Max" json:"max"`
    Idle    int     `yaml:"Idle" json:"idle"`
}

type Config struct {
	Name    string `yaml:"Name" json:"name"`
    Host    string `yaml:"Host" json:"host"`
    Port    string `yaml:"Port" json:"port"`
    Mode    string `yaml:"Mode" json:"mode"`
    Timeout time.Duration `yaml:"Timeout" json:"timeout"`
    Orm     Orm    `yaml:"Orm" json:"orm"`
}

