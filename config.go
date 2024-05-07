package main

type Config struct {
	BindAddress string  `yaml:"bind_address"`
	Routes      []Route `yaml:"routes"`
}

type Route struct {
	Path       string `yaml:"path"`
	SourceUrl  string `yaml:"sourceUrl"`
	UrlPattern string `yaml:"urlPattern"`
	RewriteUrl string `yaml:"rewriteUrl"`
}
