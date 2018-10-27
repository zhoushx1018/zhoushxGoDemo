package main

// 本demo 参考 gopkg.in/yaml.v2/README.md

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

import (
	. "zhoushxGoDemo/GoByExample/fileline"
	"os"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
  e: [test1,test2]
  f: [test3,test4]
`

type T struct {
	X1 string `yaml:"a"`
	X2 string `yaml:"aa"`
	X3 struct {
		Y1 int      `yaml:"c"`
		Y2 []string `yaml:"f,flow"`
		Y3 []int    `yaml:"d,flow"`
		Y4 []string `yaml:"e,flow"`
	}  `yaml:"b"`
}

func main() {
	t := T{}

	fmt.Printf("%s-->t_1=%v\n", FileLine(), t)

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%s-->t_2=%v\n", FileLine(), t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s-->string(d)=%s\n", FileLine(), string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s-->m=%v\n", FileLine(), m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("%s-->string(d)=%s\n", FileLine(), string(d))
}
