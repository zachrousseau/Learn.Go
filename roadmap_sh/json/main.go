package main

import (
	"encoding/json"
	"fmt"
) 

func main(){

	// IF YOU KNOW THE STRUCTURE OF JSON BEFOREHAND THIS IS GREAT 
	type FunTown struct { // Define what FunTowns should look like
		Name string 
		Body string 
		Time int64 
	}

	boston := FunTown{"Boston", "Dirty Water", 123456789} // create a FunTown boston 

	bostonJson, err := json.Marshal(boston)


	if err == nil {
		fmt.Printf("Encoded bostonJson is %v", bostonJson) // It is an array of bytes. Not really helpful.
	} else {
		fmt.Printf("%v", err)
	}


	var bostonDecoded FunTown

	err = json.Unmarshal(bostonJson, &bostonDecoded)

	if err == nil {
		fmt.Printf("\nDecoded bostonJson is %v", bostonDecoded) // It is an array of bytes. Not really helpful.
	} else {
		fmt.Printf("%v", err)
	}


	/// IF YOU DON'T KNOW JSON STRUCTURE WE CAN DO A GENERIC 

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`) // Mystery value 

	var j interface{} // Generic interface 

	err = json.Unmarshal(b, &j)

	fmt.Printf("\n\nGeneric JSON interface is %v", j)

	// we can unravel that 

	unraveledInterface := j.(map[string]interface{})

	fmt.Printf("\nCleaned up %v", unraveledInterface)

}