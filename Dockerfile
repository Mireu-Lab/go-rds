FROM docker.io/golang:alpine

RUN mkdir app
WORKDIR app

COPY . .

RUN go get cth.release/go-rds/web
RUN go build -v ./...

CMD [ "./go-rds" ]