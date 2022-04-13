FROM golang:1.17.9-bullseye as builder
WORKDIR /src
RUN apt update && apt install -y upx
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o flights cmd/api/*.go 
RUN upx flights
 
FROM scratch
COPY --from=builder /src/flights /app/flights
EXPOSE 8080
ENTRYPOINT ["/app/flights"]