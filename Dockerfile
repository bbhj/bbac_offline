FROM golang
MAINTAINER Dean dean@airdb.com

RUN go get github.com/tools/godep
RUN go get github.com/astaxie/beego
RUN go get github.com/bbhj/baobeihuijia

#RUN go build

EXPOSE 8080

CMD ["bee run -downdoc=true -gendoc=true"]
