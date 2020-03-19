//Service 1. Additional, send JSON.
//1. Create data array.
//2. Print data on the screen.
//2. Encode data array into JSON.
//3. Send JSON to service 2 on port http :8082

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Item struct {
	ID             uint    `json:"id"`   //Employees_ID
	Type           string  `json:"type"` //developer, designer, manager
	First_Name     string  `json:"first_name"`
	Second_Name    string  `json:"second_name"`
	Salary_Default int     `json:"salary_default"`
	Experience     uint    `json:"experience"`
	Coefficient    float32 `json:"coefficient"`
}

func main() {
	MakeRequest()
}

func MakeRequest() {

	message := map[string]interface{}{
		"ID":         1,
		"Type":       "designer",
		"First_Name": "Picasso",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://127.0.0.1:8082", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
}
