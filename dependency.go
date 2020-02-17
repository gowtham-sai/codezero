package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Method string
type Header map[string][]string
type Query map[string]string

const (
	Get     Method = "GET"
	Head    Method = "HEAD"
	Post    Method = "POST"
	Put     Method = "PUT"
	Patch   Method = "PATCH" // RFC 5789
	Delete  Method = "DELETE"
	Connect Method = "CONNECT"
	Options Method = "OPTIONS"
	Trace   Method = "TRACE"
)

type Request struct {
	Method  Method `yaml:"method"`
	Path    string `yaml:"path"`
	Query   Query  `yaml:"query"`
	Headers Header `yaml:"headers"`
}

type Response struct {
	Body    string `yaml:"body"`
	Headers Header `yaml:"headers"`
}

type Dependency struct {
	Req Request  `yaml:"req"`
	Res Response `yaml:"res"`
}

func ParseDependency(filepath string) (dep *Dependency, err error) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, &dep)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return
	}
	return
}
