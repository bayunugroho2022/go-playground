package main

import (
	"fmt"
	"log"
	"time"
	"my-protobuf/basic"
	// "my-protobuf/jobsearch"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes));
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// basic.BasicOneOf()
	// basic.BasicUser()
	// basic.WriteToJSONSample()
	basic.ReadFromJSONSample()
	// basic.WriteToFileSample()
	// basic.ReadFromFileSample()
	// basic.BasicUnMarshalAnyKnow()
	// basic.BasicUnMarshalAnyNotKnown()
	// basic.BasicUnMarshalAnyIs()
	// basic.ProtoToJsonUser()
	// basic.JsonToProtoUser() 
	// basic.BasicUserGroup()
	// jobsearch.JobSearchCandidate()
	// jobsearch.JobSearchSoftware()

}