package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Secret struct {
	ApiVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   Metadata          `yaml:"metadata"`
	Type       string            `yaml:"type"`
	Data       map[string]string `yaml:"data"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

func main() {
	var env map[string]string
	env, _ = godotenv.Read()

	metadata := Metadata{Name: "env"}
	y := Secret{ApiVersion: "v1", Kind: "Secret", Metadata: metadata, Type: "Opaque", Data: env}
	d, err := yaml.Marshal(&y)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("---\n%s\n", string(d))
}
