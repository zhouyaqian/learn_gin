FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /data/learn

COPY . /data/learn

RUN go build -o learnGin

EXPOSE 8080

ENTRYPOINT ["./learnGin"]


