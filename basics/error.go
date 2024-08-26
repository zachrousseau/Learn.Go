package main 

import ( 
	"errors"
	"fmt"
)


func DoubleTwo(i int) (j int, err error){
	
	if(i == 2){
		return i * 2, nil
	}
	return i * 2, errors.New("integer provider is not 2")
}

func main(){
	pandaWeight := 2

	twoPandaWeight, err := DoubleTwo(pandaWeight)

	if err == nil {
		fmt.Println(twoPandaWeight)
	} else {
		fmt.Println(err)
	}
	
}

