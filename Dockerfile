FROM library/golang

COPY main.go /go/src/github.com/rayz0620/simplehttpservice/main.go
RUN cd /go/src/github.com/rayz0620/simplehttpservice && go build

CMD /go/src/github.com/rayz0620/simplehttpservice/simplehttpservice
