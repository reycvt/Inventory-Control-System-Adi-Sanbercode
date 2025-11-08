FROM golang:1.24.8-alpine

RUN apk add --no-cache bash
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000

CMD ["go", "run", "main.go", "start"]
