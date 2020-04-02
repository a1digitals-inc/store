FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o api.o ./cmd/main.go

EXPOSE 80
ENTRYPOINT ["./api.o"]
