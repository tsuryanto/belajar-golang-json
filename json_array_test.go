package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Taufiq",
		LastName:  "Suryanto",
		Hobbies:   []string{"Learning", "Coding", "Design"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	customerString := `{"FirstName":"Taufiq","MiddleName":"","LastName":"Suryanto","Age":0,"Married":false,"Hobbies":["Learning","Coding","Design"]}`
	customerBytes := []byte(customerString)

	customer := &Customer{}
	err := json.Unmarshal(customerBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Taufiq",
		LastName:  "Suryanto",
		Addresses: []Address{
			{
				Street:     "Jl. Teluk Gong Raya",
				Country:    "Indonesia",
				PostalCode: "14450",
			},
			{
				Street:     "Jl. Sungai Kampar",
				Country:    "Indonesia",
				PostalCode: "13960",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	JSONString := `{"FirstName":"Taufiq","MiddleName":"","LastName":"Suryanto","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jl. Teluk Gong Raya","Country":"Indonesia","PostalCode":"14450"},{"Street":"Jl. Sungai Kampar","Country":"Indonesia","PostalCode":"13960"}]}`
	JSONBytes := []byte(JSONString)

	customer := &Customer{}
	err := json.Unmarshal(JSONBytes, customer)
	if err != nil {
		panic(err)
	}

	// fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	JSONString := `[{"Street":"Jl. Teluk Gong Raya","Country":"Indonesia","PostalCode":"14450"},{"Street":"Jl. Sungai Kampar","Country":"Indonesia","PostalCode":"13960"}]`
	JSONBytes := []byte(JSONString)

	addresses := &[]Address{}
	err := json.Unmarshal(JSONBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jl. Teluk Gong Raya",
			Country:    "Indonesia",
			PostalCode: "14450",
		},
		{
			Street:     "Jl. Sungai Kampar",
			Country:    "Indonesia",
			PostalCode: "13960",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
