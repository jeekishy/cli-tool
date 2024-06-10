# build and run application
run:
	docker build -f Dockerfile -t my-todo-app:latest . && docker run -it my-todo-app:latest

# test application
test:
	go test ./... -v

# run linter in application
lint:
	 golangci-lint run