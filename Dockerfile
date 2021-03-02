FROM golang:1.13 as builder
ENV GO111MODULE=on
WORKDIR /go/src/kamgo
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/kamgo/app .
COPY --from=builder /go/src/kamgo/config /config

# Expose the application on port 8080.
# This should be the same as in the app.conf file
EXPOSE 8080

# Set the entry point of the container to the application executable
CMD ["/app"]
