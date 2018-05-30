FROM golang:1.8.3 as builder
WORKDIR /src/main
RUN go get github.com/gorilla/mux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /src/main/main .
CMD ["./main"]


# iron/go:dev is the alpine image with the go tools added
# FROM iron/go:dev
# RUN mkdir /app 
# ADD . /app/ 
# WORKDIR /app 
# RUN go get github.com/gorilla/mux
# RUN go build ./src/main/main.go 
# EXPOSE 8000 8000
# CMD ["/app/main"]
