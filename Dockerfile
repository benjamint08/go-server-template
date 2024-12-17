FROM golang:1.23.4-alpine AS go-builder
WORKDIR /goapp
COPY go.mod ./
COPY main.go ./
COPY handlers/ ./handlers/
COPY helpers/ ./helpers/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /goserver .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=go-builder /goserver /app/
CMD ["/app/goserver"]