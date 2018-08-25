FROM library/golang

COPY main.go /go/src/github.com/rayz0620/GolangHttpServiceDemo/main.go
RUN cd /go/src/github.com/rayz0620/GolangHttpServiceDemo && go build

CMD /go/src/github.com/rayz0620/GolangHttpServiceDemo/GolangHttpServiceDemo
