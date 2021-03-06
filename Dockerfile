FROM golang:1.9-alpine as builder

RUN apk add --update git
RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/envy

COPY . .

RUN dep ensure
RUN go build

FROM scratch

COPY --from=builder /go/src/envy/envy /envy
