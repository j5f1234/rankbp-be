FROM golang:latest
COPY ./ /runtime/rankbp/
WORKDIR /runtime/rankbp/
RUN go env -w GOPROXY=goproxy.cn,direct
RUN go mod tidy
RUN go build -o goapp
ENTRYPOINT [ "/runtime/rankbp/goapp" ]
EXPOSE 8088
