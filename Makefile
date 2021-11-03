build:
	CGO_ENABLED=0 GOOS=linux 
	swag init --parseDependency --parseInternal
	# go build afeilulu.com/example
	go build -a -installsuffix cgo -o main .

run: build
	# go run afeilulu.com/example
	./main

docker:
	docker build -t example-scratch -f Dockerfile.scratch .
