package main

import (
	"fmt"

	"github.com/Perseverance/udemy-go-tut/12_structs/person"
)

func main() {
	Djena := person.NewWoman()
	// encoder := json.NewEncoder(os.Stdout)
	// encoder.Encode(&Djena)

	fmt.Println(Djena)

	// fmt.Println(Djena.SayYo())
	//
	// decoder := json.NewDecoder(os.Stdin)
	//
	// decoder.Decode(&Djena)
	//
	// fmt.Println(Djena)
	//
	// fmt.Println(Djena.LastName)

}
