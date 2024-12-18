package main

import "fmt"

func main() {

	m := map[string]string{
		"a": "A",
	}

	for k, v := range m {
		fmt.Println(k, v)
		m["b"] = "B"
	}

}
