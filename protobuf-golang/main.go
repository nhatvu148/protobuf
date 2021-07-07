package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/nhatvu148/protobuf/src/addressbookpb"
	"github.com/nhatvu148/protobuf/src/complexpb"
	"github.com/nhatvu148/protobuf/src/enumpb"
	"github.com/nhatvu148/protobuf/src/simplepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()

	// ReadAndWriteDemo(sm)
	jsonDemo(sm)

	doEnum()

	doComplex()

	doAddressbok()
}

func doAddressbok() {
	p := addressbookpb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*addressbookpb.Person_PhoneNumber{
			{Number: "555-4321", Type: addressbookpb.Person_HOME},
		},
	}

	ab := addressbookpb.AddressBook{
		People: []*addressbookpb.Person{
			&p,
			{
				Id:    5435,
				Name:  "Will Doe",
				Email: "wdoe@example.com",
				Phones: []*addressbookpb.Person_PhoneNumber{
					{Number: "222-3245", Type: addressbookpb.Person_HOME},
				},
			},
		},
	}

	fmt.Println(ab)

	fname := "address.bin"
	// Write the new address book back to disk.
	out, err := proto.Marshal(&ab)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &addressbookpb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println("New book:", book)
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   2,
				Name: "Second message",
			},
			{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_TUESDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY

	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON([]byte(smAsString), sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func fromJSON(in []byte, pb proto.Message) {
	err := protojson.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func toJSON(pb proto.Message) string {
	out, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return string(out)
}

func ReadAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
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

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
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
