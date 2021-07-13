# Notice: Open the correct root folder of Go projects to run protoc properly

protoc -I src/ --go_out=src/ src/simplepb/simple.proto
protoc -I src/ --go_out=src/ src/enumpb/enum.proto
protoc -I src/ --go_out=src/ src/complexpb/complex.proto
protoc -I src/ --go_out=src/ src/addressbookpb/addressbook.proto
