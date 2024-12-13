FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum app ./ 

COPY . .

RUN go build -o main .

FROM scratch

COPY --from=builder /app/configApp.yaml /configApp.yaml

COPY --from=builder /app/main /main

EXPOSE 8080

ENTRYPOINT ["./app/cmd/main"]