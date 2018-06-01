FROM golang:latest
RUN mkdir /app
ADD src/ /app/
WORKDIR /app
RUN go build main.go Simple.go
CMD ["/app/main"]
