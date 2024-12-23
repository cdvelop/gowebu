package main

import (
	"log"

	"github.com/cdvelop/pwago"
)

func main() {

	if err := pwago.Server(); err != nil {
		log.Fatal(err)
	}

}
