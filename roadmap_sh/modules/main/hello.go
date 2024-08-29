package hello 

import (
	"rsc.io/quote/v3" // Give an alternate name of quotev3
) 

func Hello() string { 
	return quote.HelloV3()
}

func Proverb() string { 
	return quote.Concurrency()
}

