# This is the first stage, for building things that will be required by the
# final stage (notably the binary)

FROM airdb/beego
MAINTAINER Dean dean@airdb.com

WORKDIR  /go/src/github.com/bbhj/bbac
COPY go.mod go.sum ./
RUN go mod download

ADD . $WORKDIR

RUN go build -o main
# RUN pwd
# The second and final stage
# FROM scratch
FROM centos

# Copy the binary from the builder stage
COPY --from=0 /go/src/github.com/bbhj/bbac/main /srv/

EXPOSE 8080

CMD ["/srv/main"]
