FROM golang:1.20

ADD . /dockerdev
WORKDIR /dockerdev

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify

#COPY . .
RUN go build -o /api
EXPOSE 8000

CMD ["/api"]