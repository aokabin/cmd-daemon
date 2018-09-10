NAME := cmdaemon

.PHONY : run
run:
	go run main.go

.PHONY : test
test:
	go test ./...

.PHONY : build
build:
	go build -o $(NAME)

.PHONY : ensure
ensure:
	dep ensure

