package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	// fmt.Println(`here is a quote: \n \"` + quote.Go() + `\"`)
	fmt.Println("here is a quote: \n \"", quote.Go(), "\"")
}
