FROM golang:1.20

WORKDIR /usr/local/exporter

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#RUN mkdir exporter
COPY go.mod go.sum ./
RUN go mod download && go mod verify

##COPY --chown=.:. . .
RUN chmod 777 /usr

RUN go build -o main.go

#RUN chmod -x /
CMD ["/usr/local/exporter"]

EXPOSE 5123