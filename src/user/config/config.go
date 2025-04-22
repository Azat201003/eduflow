package config

type Service struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	ID          int    `yaml:"id"`
}

type Config struct {
	Services []*Service `yaml:"services"`
}
