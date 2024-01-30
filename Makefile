
stack:
	go build -gcflags="-m -m" reader.go node.go decoder.go deserializer.go errors.go 2>&1 | grep "reader.go" | grep "escapes to heap"
	

