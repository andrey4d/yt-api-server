FROM golang:1.22-alpine as builder

RUN apk add make

COPY /  /yt-api-server
WORKDIR /yt-api-server

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN /go/bin/templ generate 
RUN go build -ldflags="-s -w" -o bin/yt-api-server cmd/main.go


FROM alpine

COPY --from=builder /yt-api-server/bin/yt-api-server /

ENTRYPOINT [ "/yt-api-server" ]