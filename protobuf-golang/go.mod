module github.com/nhatvu148/protobuf

go 1.16

require (
	github.com/nhatvu148/protobuf/src/complexpb v0.0.0-00010101000000-000000000000 // indirect
	github.com/nhatvu148/protobuf/src/enumpb v0.0.0-00010101000000-000000000000
	github.com/nhatvu148/protobuf/src/simplepb v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.27.1
)

replace github.com/nhatvu148/protobuf/src/simplepb => ./src/simplepb

replace github.com/nhatvu148/protobuf/src/enumpb => ./src/enumpb

replace github.com/nhatvu148/protobuf/src/complexpb => ./src/complexpb
