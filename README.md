# Bingo talks

该 slide 由 `go present` 驱动。

### go present安装方法：
```
go get -v golang.org/x/tools/cmd/present
```

### go present说明文档：
```
https://godoc.org/golang.org/x/tools/present
```

### go present使用简介：
1. 配置好 go 环境，安装好 go present
2. 克隆当前工程： git clone https://github.com/bingohuang/go-talks
3. 在工程目录下执行：present
4. 输出正常，打开浏览器访问： http://127.0.0.1:3999

### docker 打包
支持 docker 打包和分发，可修改成自己需要的脚本

Dockerfile如下：
```dockerfile
FROM hub.c.163.com/bingohuang/go-present:1.7.3
MAINTAINER Bingo Huang<me@bingohaung.com>
COPY . /slides
WORKDIR /slides
EXPOSE 1989
ENTRYPOINT ["/go/bin/present"]
CMD ["-http", "0.0.0.0:1989"]
```

docker.sh脚本如下：
```shell
#!/usr/bin/env bash

function docker-build() {
    if [ -n "$1" ]; then
        echo "build $1"
        docker build -t hub.c.163.com/bingohuang/talks:$1 .
    else
        echo "build latest"
        docker build -t hub.c.163.com/bingohuang/talks:latest .
    fi
}

function docker-push() {
    if [ -n "$1" ]; then
        echo "push $1"
        docker push hub.c.163.com/bingohuang/talks:$1
    else
        echo "push latest"
        docker push hub.c.163.com/bingohuang/talks:latest
    fi
}

function help() {
    echo "$0 build|push"
}

if [ "$1" == "" ]; then
    help
elif [ "$1" == "build" ];then
    docker-build $2
elif [ "$1" == "push" ];then
    docker-push $2
else
    help
fi
```
