package basic

import (
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"
	"google.golang.org/protobuf/proto"
)

func WriteProtoToFile(msg proto.Message, fname string) {
	b, err := proto.Marshal(msg)

	if err != nil {
		log.Fatal("Can not marshal message : ", err)
	}

	if err := ioutil.WriteFile(fname, b, 0644); err != nil {
		log.Fatal("Can not write to file : ", err)
	}
}

func ReadProtoFromFile(fname string, msg proto.Message) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatal("Can not read file : ", err)
	}

	if err := proto.Unmarshal(in, msg); err != nil {
		log.Fatal("Can not unmarshal data : ", err)
	}
}

func dummyUser() basic.User {
	addr := basic.Address{
		Street:  "street name",
		City:    "City Name",
		Country: "Country Name",
		Cordinate: &basic.Address_Cordinate{
			Latitude:  1.123123,
			Longitude: 2.123123,
		},
	}

	comm := randomCommunicationChannel()
	sr := map[string]uint32{"fly": 1, "strength": 100, "speed": 100}

	return basic.User{
		Id:                   99,
		Username:             "John Doe",
		IsActive:             true,
		Password:             []byte("secret"),
		Emails:               []string{"bayu@gmail.com", "bayu2@gmail.com"},
		Gender:               basic.Gender_GENDER_MALE,
		Address:              &addr,
		CommunicationChannel: &comm,
		SkillRating: 		sr,
	}
}

func WriteToFileSample() {
	u := dummyUser()
	WriteProtoToFile(&u, "superman_file.bin")
}

func ReadFromFileSample() {
	u := basic.User{}

	ReadProtoFromFile("superman_file.bin", &u)
	log.Println(&u)
}