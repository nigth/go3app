//SERVICE. ENCODING JSON->YAML.
//1. Listen incoming traffic on port http :8082
//2. Decode structure from JSON into map.
//3. Print mapped data on the screen.
//4. Encode data map into YAML.
//5. Send YAML to service 3 on port http :8083

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ghodss/yaml"
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

	item := Item{ID: 2, Type: "developer", Second_Name: "Gates"}
	jitem, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(jitem))

	y, err := yaml.JSONToYAML(jitem)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//var item1 Item
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		y1, err := yaml.JSONToYAML(body)

		//if err := json.Unmarshal(body, &item1); err != nil {
		//	w.WriteHeader(400)
		//}

		resp, err := http.Post("http://127.0.0.1:8083", "application/yaml", bytes.NewBuffer(y1))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintf(w, string(resp.Status))
	})
	http.ListenAndServe(":8082", nil)
}
