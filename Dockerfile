FROM golang:latest AS go_builder
ADD . /source
RUN cd /source && go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o store.o ./cmd/main.go

FROM alpine:latest
WORKDIR /store
RUN apk --no-cache add ca-certificates
COPY --from=go_builder /source/static/images /store/static/images/
COPY --from=go_builder /source/views /store/views/
COPY --from=go_builder /source/store.o /store/
RUN echo "$(ls .)"
RUN ["chmod", "+x", "./store.o"]
EXPOSE 80
ENTRYPOINT ["./store.o"]
