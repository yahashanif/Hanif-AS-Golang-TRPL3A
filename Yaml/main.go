package main

import (
	"fmt"
	"io/ioutil"

	//using  go get gopkg.in/yaml.v2
	"gopkg.in/yaml.v2"
)

// YamlConfig is exported.
type YamlConfig struct {
	Connection struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		User     string `yaml:"user"`
	}
}

func main() {

	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	var yamlConfig YamlConfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("Host: %v\n", yamlConfig.Connection.Host)
	fmt.Printf("Password: %v\n", yamlConfig.Connection.Password)
	fmt.Printf("Port: %v\n", yamlConfig.Connection.Port)
	fmt.Printf("User: %v\n", yamlConfig.Connection.User)

}
