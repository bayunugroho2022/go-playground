package basic

import (
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtoToJSON(msg proto.Message, fname string) {
	b, err := protojson.Marshal(msg)

	if err != nil {
		log.Fatal("Can not marshal message : ", err)
	}

	if err := ioutil.WriteFile(fname, b, 0644); err != nil {
		log.Fatal("Can not write to file : ", err)
	}
}

func ReadProtoFromJSON(fname string, msg proto.Message) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatal("Can not read file : ", err)
	}

	if err := protojson.Unmarshal(in, msg); err != nil {
		log.Fatal("Can not unmarshal data : ", err)
	}
}

func WriteToJSONSample() {
	u := dummyUser()
	WriteProtoToJSON(&u, "superman_file.json")
}

func ReadFromJSONSample() {
	u := basic.User{}

	ReadProtoFromJSON("superman_file.json", &u)
	log.Println(&u)
}