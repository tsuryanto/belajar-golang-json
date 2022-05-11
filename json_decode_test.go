package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Taufiq","MiddleName":"Suryanto","LastName":"S.Kom","Age":24,"Married":false}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Age)
}
