FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o fcc-file-metadata-ms

EXPOSE 3000

CMD ./fcc-file-metadata-ms