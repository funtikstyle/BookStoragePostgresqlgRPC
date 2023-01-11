FROM golang:alpine

RUN mkdir /app
ADD . /app/

# RUN mkdir /app
# RUN mkdir $GOPATH/github.com
# RUN mkdir $GOPATH/github.com/jackc
# RUN mkdir $GOPATH/github.com/jackc/pgx/
# COPY /server/pgx $GOPATH/github.com/jackc/pgx/

WORKDIR /app 
# COPY . . 
# ENTRYPOINT [" go env -w GO111MODULE=off"]

RUN go mod init server
RUN go mod tidy
RUN go build -o main /app/server
CMD ["/app/main"]