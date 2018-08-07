FROM golang
RUN go get "gopkg.in/mgo.v2"
RUN go get "github.com/juju/mgosession"
ADD . /code
WORKDIR /code
RUN go build
CMD ["./code"]