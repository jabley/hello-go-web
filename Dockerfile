FROM golang:alpine
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o hello-go-web . 
CMD ["/app/hello-go-web"]
EXPOSE 8080
