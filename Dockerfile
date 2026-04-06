FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd cmd && go build -o /app/server

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/server ./server
COPY --from=build /app/config.json ./config.json
COPY --from=build /app/client ./client

EXPOSE 8080

CMD ["./server"] 
