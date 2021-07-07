protoc -I src/ --go_out=src/ src/simplepb/simple.proto
protoc -I src/ --go_out=src/ src/enumpb/enum.proto
protoc -I src/ --go_out=src/ src/complexpb/complex.proto