package main

// Convert from YAML to JSON

// Copyright (C) Philip Schlump, 2023.
// MIT Licensed.

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

var in = flag.String("input", "", "Input yaml file to convert")
var out = flag.String("output", "", "Output json file to convert")
var indent = flag.Bool("no-indent", false, "Output JSON with indentation")

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "yaml-to-json: Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse() // Parse CLI arguments

	buf, err := ioutil.ReadFile(*in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening %s for input: %s\n", *in, err)
		os.Exit(1)
	}

	s := string(buf)

	if db1 {
		fmt.Printf("Input: %s\n", s)
	}

	var body interface{}
	if err := yaml.Unmarshal([]byte(s), &body); err != nil {
		panic(err)
	}

	body = convert(body)

	fx := UnmarshalIndentJSON
	if *indent {
		fx = UnmarshalJSON
	}

	if *out == "" {
		fmt.Printf("%s\n", fx(body))
	} else {
		ioutil.WriteFile(*out, []byte(fx(body)), 0644)
	}

}

func UnmarshalIndentJSON(v interface{}) string {
	s, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	} else {
		return string(s)
	}
}

func UnmarshalJSON(v interface{}) string {
	s, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	} else {
		return string(s)
	}
}

const db1 = false
