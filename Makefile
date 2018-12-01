build:
	mkdir -p functions
	go get ./source/hello
	go build -o ./functions/hello ./source/hello
	go get ./source/json
	go build -o ./functions/json ./source/json