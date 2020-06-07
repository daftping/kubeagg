package kubeagg

import (
	"encoding/json"
	"fmt"
	"log"
)

func (objects AllObjects) PrintTable() {

	fmt.Println("CONTEXT\tNAME")
	for _, list := range objects.Lists {
		for _, item := range list.Items {
			fmt.Printf("%v\t%v\n", list.Context, item.Metadata.Name)
		}
	}
}

func (objects AllObjects) PrintJSON() {
	json, err := json.MarshalIndent(objects, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
