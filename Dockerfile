FROM golang:1.21

WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o delivery-api ./cmd/api

WORKDIR /app

EXPOSE 8080
EXPOSE 2345

CMD ["./delivery-api"]