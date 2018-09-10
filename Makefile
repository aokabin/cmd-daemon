NAME := cmdaemon


.PHONY : test
test:
	go test ./...

.PHONY : build
build:
	go build -o $(NAME)

.PHONY : ensure
ensure:
	dep ensure

