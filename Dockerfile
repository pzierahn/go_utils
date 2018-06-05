FROM ubuntu:latest
RUN mkdir /app
ADD build/main /app/
WORKDIR /app
CMD ["/app/main"]
