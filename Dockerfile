FROM golang:1.16 as builder

WORKDIR /go/src/assignment

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

ENV CGO_ENABLED=0
RUN go build -o assignment

FROM alpine:latest
WORKDIR /assignment

COPY --from=builder /go/src/assignment/assignment .

CMD ["./assignment"]
