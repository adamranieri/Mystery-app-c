FROM golang:1.17-alpine

COPY . /workspace

WORKDIR /workspace

RUN go mod download

RUN go build -o app main.go models.go
 
EXPOSE 8080

CMD [ "./app" ]