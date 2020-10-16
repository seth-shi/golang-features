## Golang 功能模块站点


### 部署
* 前提条件, docker, docker-compose
* 克隆代码
* `git clone https://github.com/seth-shi/golang-features.git && cd golang-features`
* 复制`.env.example`为`.env`按需配置
* `docker-compose up -d app`

****

* 开发环境请使用`dev`容器
* `docker-compose up -d dev`
* `docker-compose exec dev bash`
* 使用热加载代码
* `air`