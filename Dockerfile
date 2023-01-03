FROM golang:1.19.4

RUN rm -rf /app
RUN mkdir /app

COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

WORKDIR /app
RUN ls
RUN go mod download

COPY . /app

EXPOSE 8081

CMD go run main.go
