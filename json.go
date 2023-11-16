package main

import (
	"encoding/json"
	"fmt"
	"time"
	"log"
	"os"
)


type MyJSON struct {
	IntValue        int       `json:"intValue"`
	BoolValue       bool      `json:"boolValue"`
	StringValue     string    `json:"stringValue"`
	DateValue       time.Time `json:"dateValue"`
	ObjectValue     *myObject `json:"objectValue"`
	NullStringValue *string   `json:"nullStringValue"`
	NullIntValue    *int      `json:"nullIntValue"`
}

type myObject struct {
	ArrayValue []int `json:"arrayValue"`
}

func main() {
	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"dateValue":   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
		"nullStringValue": nil,
		"nullIntValue":    nil,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	
	otherInt := 4321
	data1 := &MyJSON{
		IntValue:    1234,
		BoolValue:   true,
		StringValue: "hello!",
		DateValue:   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		ObjectValue: &myObject{
			ArrayValue: []int{1, 2, 3, 4},
		},
		NullStringValue: nil,
		NullIntValue:    &otherInt,
	}
	out, err := json.MarshalIndent(data1, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("json data: %s\n", jsonData)
	fmt.Println(string(out))
	//  parsing data  using a map 
	jsonData1 := `
		{
			"intValue":1234,
			"boolValue":true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null,
			"extraValue":4321
		}
	`

	var data2 map[string]interface{}
	err = json.Unmarshal([]byte(jsonData1), &data2)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json map: %v\n", data2)

	rawDateValue,ok := data2["dateValue"]
	if !ok {
		fmt.Printf("dateValue does not exist\n")
		return
	}
	dateValue, ok := rawDateValue.(string)
	if !ok {
		fmt.Printf("dateValue is not a string\n")
		return
	}
	fmt.Printf("date value: %s\n", dateValue)
	// parsing data using  a struct

	var data3 *MyJSON
	err1 := json.Unmarshal([]byte(out), &data3)
	if err1 != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}
	fmt.Printf("json struct: %#v\n", data3)
	fmt.Printf("dateValue: %#v\n", data3.DateValue)
	fmt.Printf("objectValue: %#v\n", data3.ObjectValue)

}