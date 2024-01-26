package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicWriteUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug: "slug",
		// Title: "title",
		HtmlContent: "html content",
		// AuthorId: 1,
	}

	WriteProtoToFile(&uc, "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Read User Content V1")
	uc := basic.UserContent{}

	ReadProtoFromFile("user_content_v1.bin", &uc)
	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
	log.Println()
}

func BasicWriteUserContentV2() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug: "slug",
		// Title: "title",
		HtmlContent: "<p> html content </p>",
		// AuthorId: 1,
		// Category: "category",
	}

	WriteProtoToFile(&uc, "user_content_v2.bin")
}

func BasicReadUserContentV2() {
	log.Println("Read User Content V2")
	uc := basic.UserContent{}

	ReadProtoFromFile("user_content_v2.bin", &uc)
	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
	log.Println()
}


func BasicWriteUserContentV3() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug: "slug",
		// Title: "title",
		HtmlContent: "<p> html content </p>",
		// AuthorId: 1,
		// Category: "category",
		// SubCategory: "sub category",
	}

	WriteProtoToFile(&uc, "user_content_v3.bin")
}

func BasicReadUserContentV3() {
	log.Println("Read User Content V3")
	uc := basic.UserContent{}

	ReadProtoFromFile("user_content_v3.bin", &uc)
	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
	log.Println()
}
