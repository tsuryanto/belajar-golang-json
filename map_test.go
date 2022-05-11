package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0002", "name":"Apple IPad Pro", "price":15000000}`
	jsonByte := []byte(jsonString)

	// product := make(map[string]interface{})
	var product map[string]interface{}

	err := json.Unmarshal(jsonByte, &product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product["id"])
	fmt.Println(product["name"])
	fmt.Println(product["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0003",
		"name":  "Iphone 13",
		"price": 12000000,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
