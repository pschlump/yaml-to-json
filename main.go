package main

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

	if *out == "" {
		fmt.Printf("%s\n", UnmarshalIndentJSON(body))
	} else {
		ioutil.WriteFile(*out, []byte(UnmarshalIndentJSON(body)), 0644)
	}

	//	if b, err := json.Marshal(body); err != nil {
	//		panic(err)
	//	} else {
	//		fmt.Printf("Output: %s\n", b)
	//	}

}

func UnmarshalIndentJSON(v interface{}) string {
	// s, err := json.Marshal ( v )
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	} else {
		return string(s)
	}
}

const db1 = false
