package pkg

import (
	"encoding/json"
	"fmt"
	"log"
)

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	log.Print(obj)
	data, err := json.Marshal(obj)

	if err != nil {
		return
	}

	fmt.Println("JSON data:", string(data))

	err = json.Unmarshal(data, &newMap) // Convert to a map

	fmt.Println(newMap)
	return
}
