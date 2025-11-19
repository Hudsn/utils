package main

import (
	"fmt"
	"log"

	"github.com/hudsn/utils/uuid"
)

func main() {
	for range 10 {

		u, err := uuid.New()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u.String())
	}
}
