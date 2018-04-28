FROM golang:latest
ENV WORKDIR $GOPATH/src/github.com/j0nk0/ransomware
ENV GOROOT /usr/local/go
ADD . $WORKDIR
WORKDIR $WORKDIR
RUN make deps
VOLUME ["$GOPATH/src/github.com/j0nk0/ransomware/bin"]