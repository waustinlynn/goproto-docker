# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get github.com/gorilla/mux
RUN go build ./src/main/main.go 
CMD ["/app/main"]
