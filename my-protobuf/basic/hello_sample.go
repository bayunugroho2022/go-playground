package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{
		Name: "John Doe",
	}

	log.Println(&h)
}