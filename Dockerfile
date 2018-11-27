# This is the first stage, for building things that will be required by the
# final stage (notably the binary)

FROM airdb/beego:1.8.4
MAINTAINER Dean dean@airdb.com

# RUN go get github.com/tools/godep
# RUN go get github.com/astaxie/beego
# RUN go get github.com/beego/bee
# RUN go get github.com/bbhj/baobeihuijia
# RUN godep get

WORKDIR  /go/src/github.com/bbhj/baobeihuijia
COPY go.mod go.sum ./
RUN go mod download

ADD . $WORKDIR
RUN godep get -v

# The second and final stage
FROM scratch

# Copy the binary from the builder stage
COPY --from=0 /myapp /myapp

EXPOSE 8080

CMD ["bee run -downdoc=true -gendoc=true"]
