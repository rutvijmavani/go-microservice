package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"google.golang.org/protobuf/proto"
	"movieexample.com/src/gen"
	"movieexample.com/src/metadata/pkg/model"
)

var metadata = &model.Metadata{
	ID:          "123",
	Title:       "The Movie 2",
	Description: "Sequel of the legendary The Movie",
	Director:    "Foo Bars",
}

var genMetadata = &gen.Metadata{
	Id:          "123",
	Title:       "The Movie 2",
	Description: "Sequel of the legendary The Movie",
	Director:    "Foo Bars",
}

// serializeToJSON serializes the given Metadata to JSON and returns the encoded bytes.
// It returns an error if encoding fails.
func serializeToJSON(m *model.Metadata) ([]byte, error) {
	return json.Marshal(m)
}

// serializeToXML marshals the provided Metadata value to XML and returns the resulting bytes.
// It returns an error if the marshaling fails.
func serializeToXML(m *model.Metadata) ([]byte, error) {
	return xml.Marshal(m)
}

// serializeToProto serializes the provided gen.Metadata into protobuf-encoded bytes.
// It returns the encoded bytes or an error if marshaling fails.
func serializeToProto(m *gen.Metadata) ([]byte, error) {
	return proto.Marshal(m)
}

// main constructs example metadata, serializes it to JSON, XML, and Protobuf,
// and prints the resulting byte sizes. It panics if any serialization fails.
func main() {
	jsonBytes, err := serializeToJSON(metadata)
	if err != nil {
		panic(err)
	}

	xmlBytes, err := serializeToXML(metadata)
	if err != nil {
		panic(err)
	}

	protoBytes, err := serializeToProto(genMetadata)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON size:\t%dB\n", len(jsonBytes))
	fmt.Printf("XML size:\t%dB\n", len(xmlBytes))
	fmt.Printf("Proto size:\t%dB\n", len(protoBytes))
}