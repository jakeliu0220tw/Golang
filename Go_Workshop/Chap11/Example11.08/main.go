package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	byteStr := []byte(`
		{
			"checknum":123,
			"amount":200,
			"catetory":["gift","clothing"]
		}	
	`)

	var m map[string]interface{}

	if err := json.Unmarshal(byteStr, &m); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("%s=>%v, Type:%s\n", k, v, findType(v))
	}
}

func findType(v interface{}) string {
	// switch v.(type) {
	// case string:
	// 	return "string"
	// case int:
	// 	return "int"
	// case float64:
	// 	return "float64"
	// case bool:
	// 	return "bool"
	// default:
	// 	return fmt.Sprintf("%T", v)
	// }

	return fmt.Sprintf("%T", v)
}
