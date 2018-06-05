FROM golang:latest
EXPOSE 8080
RUN mkdir /app
ADD src/github.com/patrickz98/GoGo/ /app/
WORKDIR /app
RUN go build -o main
CMD ["/app/main"]
