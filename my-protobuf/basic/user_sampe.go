package basic

import (
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func BasicUser() {
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

	u := basic.User{
		Id:                   99,
		Username:             "John Doe",
		IsActive:             true,
		Password:             []byte("secret"),
		Emails:               []string{"bayu@gmail.com", "bayu2@gmail.com"},
		Gender:               basic.Gender_GENDER_MALE,
		Address:              &addr,
		CommunicationChannel: &comm,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       99,
		Username: "John Doe",
		IsActive: true,
		Password: []byte("secret"),
		Emails:   []string{"bayu3@gmail.com", "bayu4@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
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

func randomCommunicationChannel() anypb.Any {
	rand.Seed(time.Now().UnixNano())

	paperMail := basic.PaperMail{
		PaperMailAddress: "Some paper mail address",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "byteMe",
		SocialMediaUsername: "social media username",
	}

	instantMessagaging := basic.InstantMessaging{
		InstantMessagingProduct:  "instant messaging product",
		InstantMessagingUsername: "instant messaging username",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessagaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnMarshalAnyKnow() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "my-social-media-platform",
		SocialMediaUsername: "my-social-media-username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	sm := &basic.SocialMedia{}

	if err := a.UnmarshalTo(sm); err != nil {
		log.Fatal(err)
	}

	jsonBytes, _ := protojson.Marshal(sm)
	log.Println(string(jsonBytes))
}

func BasicUnMarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarsheled protoreflect.ProtoMessage

	unmarsheled, err := a.UnmarshalNew()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Unmarshaled: ", unmarsheled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarsheled)
	log.Println(string(jsonBytes))

}

func BasicUnMarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {

		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatal(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Message is not PaperMail", a.TypeUrl)
	}
}

func BasicOneOf() {
	// option 1
	// socialMedia := basic.SocialMedia{
	// 	SocialMediaPlatform: "byteMe",
	// 	SocialMediaUsername: "username bayu",
	// }

	// ecomm := basic.User_SocialMedia{
	// 	SocialMedia: &socialMedia,
	// }


	// option 2
	instanMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "instant messaging product",
		InstantMessagingUsername: "instant messaging username",
	}

	ecomm := basic.User_InstantMessaging{
		InstantMessaging: &instanMessaging,
	}

	u := basic.User{
		Id:                    99,
		Username:              "John Doe",
		IsActive:              true,
		ElectronicCommChannel: &ecomm,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}
