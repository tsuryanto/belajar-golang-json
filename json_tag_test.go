package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	products := Product{
		Id:       "P0001",
		Name:     "Apple Macbook Pro",
		ImageURL: "http://apple.id",
	}

	bytes, _ := json.Marshal(products)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Macbook Pro","image_url":"http://apple.id"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
