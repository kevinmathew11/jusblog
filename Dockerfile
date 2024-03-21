FROM golang:1.19.0


WORKDIR /Users/kevinmathew/blog

COPY . .

RUN go mod tidy
