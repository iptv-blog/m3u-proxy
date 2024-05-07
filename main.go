package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	if _, err := os.Stat("config.yaml"); errors.Is(err, os.ErrNotExist) {
		log.Fatal("config.yaml not found")
	}
	log.Println("Loading configuration from config.yaml...")
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("Could not open config.yaml: %v ", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	for _, route := range config.Routes {
		compiled, err := regexp.Compile(route.UrlPattern)
		if err != nil {
			log.Fatalf("Error while compiling url pattern: %v", err)
		}
		r.GET(route.Path, func(c *gin.Context) {
			resp, e := http.Get(route.SourceUrl)
			if e != nil {
				c.String(http.StatusInternalServerError, "Error while fetching source url")
				log.Printf("Error while fetching source url: %v", e)
				return
			}
			defer resp.Body.Close()
			body, e := io.ReadAll(resp.Body)
			if e != nil {
				c.String(http.StatusInternalServerError, "Error while reading from source url")
				log.Printf("Error while reading source url: %v", e)
				return
			}
			result := compiled.ReplaceAll(body, []byte(route.RewriteUrl))
			c.Data(http.StatusOK, resp.Header.Get("Content-Type"), result)
		})
		log.Printf("Registered route: %s -> %s", route.Path, route.SourceUrl)
	}
	log.Printf("Starting server on %s", config.BindAddress)
	err = r.Run(config.BindAddress)
	if err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}
