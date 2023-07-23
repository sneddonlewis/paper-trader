FROM golang:1.20-alpine3.16 AS GO_BUILD
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.16.3
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
