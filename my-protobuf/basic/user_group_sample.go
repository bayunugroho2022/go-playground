package basic

import (
	"log"
	"my-protobuf/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	batman := basic.User{
		Id: 97,
		Username: "John Doe",
		IsActive: true,
		Password: []byte("secret"),
		Gender: basic.Gender_GENDER_MALE,
	}

	nightwing := basic.User{
		Id: 96,
		Username: "John Doe",
		IsActive: true,
		Password: []byte("secret"),
		Gender: basic.Gender_GENDER_MALE,
	}

	robin := basic.User{
		Id: 95,
		Username: "John Doe",
		IsActive: true,
		Password: []byte("secret"),
		Gender: basic.Gender_GENDER_MALE,
	}

	batFamily := basic.UserGroup{
		GroupId: 999,
		GroupName: "Bat Family",
		Users: []*basic.User{&batman, &nightwing, &robin},
		Description: "Bat Family is a group of vigilantes who pledge to fight crime together.",
	}

	jsonBytes, _ := protojson.Marshal(&batFamily)
	log.Println(string(jsonBytes))
}