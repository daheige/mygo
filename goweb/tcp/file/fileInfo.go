package main

import (
	"log"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		log.Println("useage: xxx file")
		return
	}

	filename := list[1]
	info, err := os.Stat(filename)
	log.Println(info, err)
	log.Println(info.Name(), info.Size())
}
