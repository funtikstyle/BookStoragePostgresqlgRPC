FROM golang:alpine

RUN mkdir /app
ADD .. /app/

# RUN mkdir /app
# RUN mkdir $GOPATH/github.com
# RUN mkdir $GOPATH/github.com/jackc
# RUN mkdir $GOPATH/github.com/jackc/pgx/
# COPY /server/pgx $GOPATH/github.com/jackc/pgx/

WORKDIR /app
# COPY . .
# ENTRYPOINT [" go env -w GO111MODULE=off"]

#RUN go mod init client
#RUN go mod tidy


#WORKDIR ..

#RUN go build

#RUN file="$(ls -1 /app)" && echo $file
#RUN mkdir -p /app && touch /app/file1 /app/file2
#RUN echo $(ls -1 /app)
#RUN cp -R /app/client /app/client1
#RUN cp -R /app/client/config /app/client/config1
#RUN go build -o main /
#RUN ls
#RUN go build -o main
#RUN go mod init client
RUN go mod tidy
RUN go build -o client /app
CMD ["/app/client"]
#ENTRYPOINT ["*.sh"]