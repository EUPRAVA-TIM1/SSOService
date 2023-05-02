FROM golang:alpine as build_container
WORKDIR /app
COPY ./SSOService/go.mod ./SSOService/go.sum ./
RUN go mod download
COPY ./SSOService/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine
WORKDIR /root/
COPY --from=build_container /app/main .

EXPOSE 8000
ENTRYPOINT ["./main"]