package jobsearch

import (
	"log"
	"my-protobuf/protogen/basic"
	"my-protobuf/protogen/dummy"
	"my-protobuf/protogen/jobsearch"

	"google.golang.org/protobuf/encoding/protojson"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 1,
		Application: &basic.Application{
			Version: "1.0.0",
			Name: "The Amazing Proto",
			Platforms: []string{"Mac", "Linux", "Windows"},
		},
	} 

	jsonBytes, _ := protojson.Marshal(&js)
	log.Println(string(jsonBytes))
}

func JobSearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 1,
		Application: &dummy.Application{
			ApplicationId:		 	1,
			ApplicationFullName: 	"The Amazing Proto",
			Phone: 				 	"08123456789",
			Email: 					"bayu@gmail.com",
		},
	} 

	jsonBytes, _ := protojson.Marshal(&jc)
	log.Println(string(jsonBytes))
}