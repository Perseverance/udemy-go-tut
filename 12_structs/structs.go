package main

import (
	"encoding/json"
	"fmt"

	"github.com/Perseverance/udemy-go-tut/12_structs/person"
)

func main() {
	ivan := person.NewPerson()
	fmt.Println(ivan.SayYo())
	bs, _ := json.Marshal(ivan)

	fmt.Println(string(bs))
}
