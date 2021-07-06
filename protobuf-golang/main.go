package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/nhatvu148/protobuf/src/simple"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()

	ReadAndWriteDemo(sm)
}

func ReadAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simple.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content:", sm2)
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into e protocol buffer struct", err2)
		return err2
	}

	return nil
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func doSimple() *simple.SimpleMessage {
	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 9},
	}

	fmt.Println(sm)
	sm.Name = "I renamed you"
	fmt.Println(sm)
	fmt.Println("The ID is: ", sm.GetId())

	return &sm
}
