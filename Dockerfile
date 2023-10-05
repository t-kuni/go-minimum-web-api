FROM golang:1.21.1-alpine
COPY main.go .
RUN go build -o server main.go
CMD ["./server"]
