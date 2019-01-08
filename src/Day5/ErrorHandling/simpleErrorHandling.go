package main

import (
	"errors"
	"log"
)

func main() {
	err := ProcessData(0)

	if err != nil {
		log.Print(err)
	}
}

func ProcessData(id int) (err error) {

	if id == 0 {
		err = errors.New("something went wrong....")
	}

	return

}
