FROM golang:alpine3.16
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o myapp .
CMD ["/app/myapp"]
EXPOSE 3000
