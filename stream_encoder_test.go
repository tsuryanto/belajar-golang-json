package belajar_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Taufiq",
		MiddleName: "Suryanto",
		LastName:   "S.Kom",
	}

	encoder.Encode(customer)
}
