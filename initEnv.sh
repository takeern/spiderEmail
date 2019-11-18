#!/bin/bash
# Program:
#	handle project in docker env error

echo "export GOPROXY=https://goproxy.cn" >> ~/.zshrc
echo "export GO111MODULE=on" >> ~/.zshrc
source ~/.zshrc