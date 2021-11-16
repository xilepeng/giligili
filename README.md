

```
➜  giligili git:(main) ✗ go build main.go

➜  giligili git:(main) ✗ ./main 



3. 将镜像推送到Registry

$ docker login --username=你那个面试咋样啊 registry.ap-southeast-1.aliyuncs.com
$ docker tag [ImageId] registry.ap-southeast-1.aliyuncs.com/hfbpw/giligili-vue:[镜像版本号]
$ docker push registry.ap-southeast-1.aliyuncs.com/hfbpw/giligili-vue:[镜像版本号]

docker build -t registry.ap-southeast-1.aliyuncs.com/hfbpw/giligili:v0.0.1 ./

docker push registry.ap-southeast-1.aliyuncs.com/hfbpw/giligili:v0.0.1




官方：
docker run -d -p 8000:8000 -p 9443:9443 --name portainer \
    --restart=always \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v portainer_data:/data \
    portainer/portainer-ce:latest


使用：
docker volume create portainer_data

docker run -d -p 9000:9000 --name portainer \
    --restart=always \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v portainer_data:/data \
    portainer/portainer-ce:latest




忘记密码：

root@hfbpw:~# docker stop portainer
portainer

root@hfbpw:/# docker run --rm -v portainer_data:/data portainer/helper-reset-password
Unable to find image 'portainer/helper-reset-password:latest' locally
latest: Pulling from portainer/helper-reset-password
79916c70cb9e: Pull complete
93e26fa95550: Pull complete
Digest: sha256:735a809b1bfe14b5fae340d4b350bae97c2016371c47fb6e34d71a45e4512f79
Status: Downloaded newer image for portainer/helper-reset-password:latest
2021/11/12 12:04:42 Password succesfully updated for user: admin
2021/11/12 12:04:42 Use the following password to login: y3F`x%52oBX_k\1a4/HI:f9rKGU]=d76

root@hfbpw:~# docker start portainer

root@hfbpw:~# sudo apt install nginx -y


Failure
failed to deploy a stack: 
Named volume "mysql_data:/var/lib/mysql/data:rw" is used in service "mysql" but no declaration was found in the volumes section. : exit status 1


root@hfbpw:/etc/nginx/sites-enabled# vim go.conf
root@hfbpw:/etc/nginx/sites-enabled# cat go.conf
server {
    listen 80;
    server_name www.hfbpw.top;

    location / {
      proxy_set_header X-Real-Ip $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_pass http://127.0.0.1:3001;
    }

    location /api {
      proxy_set_header X-Real-Ip $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_pass http://127.0.0.1:3002;
    }
}
root@hfbpw:/etc/nginx/sites-enabled# sudo nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
root@hfbpw:/etc/nginx/sites-enabled# sudo service nginx restart
root@hfbpw:/etc/nginx/sites-enabled# sudo nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful




version: '2'

services:
  mysql:
    image: mysql:5.6
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: 123456
    volumes:
      - data:/var/lib/mysql/data
    ports:
      - 3306:3306
      
volumes:
  data: {}




version: '2'

services:
  redis:
    image: redis
    restart: always
    volumes:
      - data:/data
    ports:
      - 6379:6379
      
volumes:
  data: {}




# 查看服务器内存
root@hfbpw:~# apt install htop





```


## giligili视频网站

## 项目地址[giligili](http://www.hfbpw.top)

## 目的

本项目采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建Restful Web API

## 特色

本项目已经整合了许多开发API所必要的组件：

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](http://gorm.io/docs/index.html): ORM工具。本项目需要配合Mysql使用 
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
7. 自行实现了国际化i18n的一些基本功能
8. 本项目是使用基于cookie实现的session来保存登录状态的，如果需要可以自行修改为token验证

本项目已经实现了一些功能:
1. 用户系统
2. 视频上传
3. 视频播放（防盗链）
4. 图片上传
5. 图片下载（防盗链）
6. 视频评论
7. 播放量记录
8. 热点视频排行榜
9. 简陋聊天室


本项目已经预先创建了一系列文件夹划分出下列模块:

1. api文件夹就是MVC框架的controller，负责协调各部件完成任务
2. model文件夹负责存储数据库模型和数据库操作相关的代码
3. service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象
5. cache负责redis缓存相关的代码
6. auth权限控制文件夹
7. util一些通用的小工具
8. conf放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRE="" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

## Go 依赖

本项目使用[go mod]管理依赖，可自己运行项目，go mod会自动下载依赖
如果包下载失败，请运行
```.env
go env -w GOPROXY=https://goproxy.io,direct
go mod download
```

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)
