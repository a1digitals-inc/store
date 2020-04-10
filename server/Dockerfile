FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o api.o ./cmd/main.go" --command="./api.o"
