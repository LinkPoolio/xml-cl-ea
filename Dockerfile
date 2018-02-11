FROM golang:1.8

WORKDIR /go/src/xml-cl-ea
COPY adaptor/* /go/src/xml-cl-ea/

RUN go get
RUN go install

CMD ["xml-cl-ea"]
