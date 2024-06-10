FROM golang:1.22

WORKDIR /my-app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /my-app/todo .

CMD ["/my-app/todo", "https://jsonplaceholder.typicode.com/todos", "20"]