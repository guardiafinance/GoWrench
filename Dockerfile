FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o whrenchapp ./app/cmd
#####

FROM alpine
COPY --from=build /app/whrenchapp /

ENTRYPOINT [ "/whrenchapp" ]