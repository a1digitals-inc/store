FROM golang:latest AS go_builder
ADD . /source
RUN cd /source && go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o store.o ./cmd/main.go

FROM node:latest AS vue_builder
COPY client/package*.json ./
COPY client/babel.config.js ./
RUN npm install
COPY client/public/ ./public/
COPY client/src/ ./src/
RUN npm run build


FROM alpine:latest
WORKDIR /store
RUN apk --no-cache add ca-certificates
COPY --from=go_builder /source/static/images /store/static/images/
COPY --from=go_builder /source/store.o /store/
COPY --from=vue_builder /dist/ /store/client/dist/
RUN echo "$(ls .)"
RUN ["chmod", "+x", "./store.o"]
EXPOSE 80
ENTRYPOINT ["./store.o"]
