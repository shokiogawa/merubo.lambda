FROM --platform=linux/amd64 golang:1.21.5-alpine as builder
WORKDIR /go/merubo/
COPY ../src ./src
COPY ../go.mod go.sum ./
RUN apk update && apk --no-cache add git
RUN go mod tidy
WORKDIR /go/merubo/src
RUN CGO_ENABLE=0 GOARCH=amd64 GOOS=linux go build -o /go/merubo/binary


#production
FROM alpine as production
WORKDIR go/merubo/production
RUN apk add --no-cache ca-certificates && apk add curl
COPY --from=builder /go/merubo/binary /go/merubo/production
ENV PORT=${PORT}
CMD ["/go/merubo/production/binary"]