package basic

import (
	"log"
	"my-protobuf/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	addr := basic.Address{
		Street: "street name", 
		City: "City Name", 
		Country: "Country Name", 
	}

	u := basic.User{
		Id: 99,
		Username: "John Doe",
		IsActive: true,
		Password: []byte("secret"),
		Emails: []string{"bayu@gmail.com", "bayu2@gmail.com"},
		Gender: basic.Gender_GENDER_MALE,
		Address: &addr,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id: 99,
		Username: "John Doe",
		IsActive: true,
		Password: []byte("secret"),
		Emails: []string{"bayu3@gmail.com", "bayu4@gmail.com"},
		Gender: basic.Gender_GENDER_MALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `
	{
		"id": 99,
		"username": "John Doe",
		"is_active": true,
		"password": "secret",
		"emails": [
			"bayu@gmail.com",
			"bayu2@gmail.com"
			],
		"gender": "GENDER_MALE"
	}`

	var p basic.User
	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(&p)
}