#!/bin/sh

### 此脚本仅限Ubuntu系统使用 ###

# 执行nvm安装脚本
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
# 导入环境变量
echo "export NVM_NODEJS_ORG_MIRROR=https://npmmirror.com/mirrors/node" >> ~/.bashrc
# 使环境变量生效
source ~/.bashrc
# 安装低版本python3
sudo add-apt-repository ppa:deadsnakes/ppa && sudo apt update && sudo apt install -y python3.10
# 安装 Node.js v16 的最新版
PYTHON=python3.10 nvm install v16.20.2 && nvm install v22.16.0
# 查看node/npm版本
node -v && npm -v
# 设置默认node版本16
nvm use 16 && nvm alias default 16
# 安装@vue/cli 5.0.8
yarn global add @vue/cli@5.0.8 && vue --version
# 设置cms项目默认使用yarn v1.22.x版本(node 16使用)
cd ${GOPATH}/src/gitlab.goiot.net/chargingc/cms-web2 && corepack enable && corepack prepare yarn@v1.22.22 --activate && yarn install && yarn serve

