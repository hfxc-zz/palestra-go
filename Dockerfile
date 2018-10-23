FROM golang

RUN go get -u -v "github.com/juju/mgosession"
RUN go get -u -v "gopkg.in/mgo.v2"

WORKDIR src

ADD . palestra-go
WORKDIR palestra-go