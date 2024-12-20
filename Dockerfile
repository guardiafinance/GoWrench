FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o whrenchapp ./app/cmd
#####

FROM ubunto

# Comentado devido a necessidade de execucao com o usuario root
# RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
# USER nonroot

COPY --from=build /app/whrenchapp /

ENTRYPOINT [ "/whrenchapp" ]