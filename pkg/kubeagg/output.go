package kubeagg

import (
	"encoding/json"
	"fmt"
	"log"
)

func (objects AllObjects) PrintWide() {
	fmt.Println("CONTEXT\t\tTYPE\tNAME")
	for _, list := range objects.Lists {
		for _, item := range list.Items {
			fmt.Printf("%v\t%v\t%v\n", list.Context, objects.Type, item.Metadata.Name)
		}
	}
}

func (objects AllObjects) PrintTable() {

	fmt.Println("CONTEXT\t\tNAME")
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
