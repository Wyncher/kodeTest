FROM golang:alpine
WORKDIR /build
COPY go.mod ./
RUN go mod download
COPY . .
CMD ["go", "run", "main.go"]
EXPOSE 8080
