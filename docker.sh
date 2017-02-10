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
