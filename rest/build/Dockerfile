FROM golang:1.8
WORKDIR  $GOPATH/src/github.com/alemjc/gophercises/cyoa/rest
COPY ./rest  $GOPATH/src/github.com/alemjc/gophercises/cyoa/rest
RUN go get -u github.com/gorilla/mux
RUN go install -v github.com/alemjc/gophercises/cyoa/rest/cmd

EXPOSE $PORT

CMD ["cmd"]

