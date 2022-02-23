package main

import "fmt"

func main() {

	m := map[string]string{"key1": "val1", "key2": "val2"}

	for k, v := range m {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}

	for k := range m {
		fmt.Printf("key[%s] value[%s]\n", k, m[k])
	}

}
