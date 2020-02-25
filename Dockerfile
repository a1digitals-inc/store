FROM golang:latest AS go_builder
ADD . /source
RUN cd /source && go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o api.o ./cmd/main.go

FROM node:alpine as node_builder
COPY --from=go_builder /source/src ./src
COPY --from=go_builder /source/package.json ./
COPY --from=go_builder /source/webpack.common.js ./
COPY --from=go_builder /source/webpack.prod.js ./
COPY --from=go_builder /source/.babelrc ./
RUN npm install
RUN npx webpack --config webpack.prod.js

FROM alpine:latest
WORKDIR /calida
RUN apk --no-cache add ca-certificates
COPY --from=go_builder /source/views /calida/views/
COPY --from=go_builder /source/api.o /calida/
COPY --from=node_builder /static /calida/static/
RUN echo "$(ls .)"
RUN ["chmod", "+x", "./api.o"]
EXPOSE 80
ENTRYPOINT ["./api.o"]
