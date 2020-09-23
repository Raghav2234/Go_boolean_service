FROM golang:1.15.1
RUN mkdir -p /Go_boolean_service
WORKDIR /Go_boolean_service
ADD . /Go_boolean_service
RUN go build .
CMD ["./Go_boolean_service"]



