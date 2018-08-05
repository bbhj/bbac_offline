FROM airdb/beego:1.8.4
MAINTAINER Dean dean@airdb.com

# RUN go get github.com/tools/godep
# RUN go get github.com/astaxie/beego
# RUN go get github.com/beego/bee
# RUN go get github.com/bbhj/baobeihuijia
# RUN godep get

WORKDIR  /go/src/github.com/bbhj/baobeihuijia
ADD . $WORKDIR
RUN godep get -v

EXPOSE 8080

CMD ["bee run -downdoc=true -gendoc=true"]
