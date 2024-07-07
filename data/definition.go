package data

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Service struct {
	Name             string            `yaml:"name"`
	HighAvailability bool              `yaml:"highAvailability"`
	Image            string            `yaml:"image"`
	Labels           map[string]Labels `yaml:"labels,omitempty"`
}

type Labels struct {
	Application string `yaml:"application,omitempty"`
	Service     string `yaml:"service-type,omitempty"`
}

func Decode(data string) *Service {
	s := &Service{}
	e := yaml.Unmarshal([]byte(data), s)
	if e != nil {
		log.Println(e)
	}
	return s
}
