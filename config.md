



```go
➜  giligili git:(main) ✗ go build main.go

➜  giligili git:(main) ✗ ./main 



3. 将镜像推送到Registry

$ docker login --username=你那个面试咋样啊 registry.ap-southeast-1.aliyuncs.com
$ docker tag [ImageId] registry.ap-southeast-1.aliyuncs.com/hfbpw/giligili-vue:[镜像版本号]
$ docker push registry.ap-southeast-1.aliyuncs.com/hfbpw/giligili-vue:[镜像版本号]

docker build -t registry.ap-southeast-1.aliyuncs.com/hfbpw/giligili:v0.0.3 ./

docker push registry.ap-southeast-1.aliyuncs.com/hfbpw/giligili:v0.0.3


snap install docker



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



TCP	接受	9000	9000	36.142.146.43/24


docker volume rm portainer_data

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



# 安装 nginx

sudo apt update

root@hfbpw:~# sudo apt install nginx -y


Failure
failed to deploy a stack: 
Named volume "mysql_data:/var/lib/mysql/data:rw" is used in service "mysql" but no declaration was found in the volumes section. : exit status 1

cd /etc/nginx/sites-enabled

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


# 检查配置
root@hfbpw:/etc/nginx/sites-enabled# sudo nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
# 让配置生效
root@hfbpw:/etc/nginx/sites-enabled# sudo service nginx restart
root@hfbpw:/etc/nginx/sites-enabled# sudo nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful


# mysql

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


# redis

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



# api

version: '2'

services:

  api:
    image: registry.ap-southeast-1.aliyuncs.com/hfbpw/giligili:v0.0.3
    restart: always
    environment:
      MYSQL_DSN: "root:123456@tcp(172.16.0.2:3306)/giligili?charset=utf8mb4,utf8&parseTime=True&loc=Local"
      REDIS_ADDR: "172.16.0.2:6379"
      REDIS_DB: "0"
      SESSION_SECRET: "rBaXcd1PPrC1"
      GIN_MODE: "release"
    ports:
      - 3002:3000



# vue

version: '2'

services:
  nginx:
    image: registry.ap-southeast-1.aliyuncs.com/hfbpw/giligili-vue:v0.0.3
    restart: always
    ports:
      - 3001:80

# 查看服务器内存
root@hfbpw:~# apt install htop




➜  ~ curl -voa 'oss-cn-hongkong.aliyuncs.com' -H 'Origin:http://www.hfbpw.top'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying 47.75.19.144:80...
* Connected to oss-cn-hongkong.aliyuncs.com (47.75.19.144) port 80 (#0)
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0> GET / HTTP/1.1
> Host: oss-cn-hongkong.aliyuncs.com
> User-Agent: curl/7.77.0
> Accept: */*
> Origin:http://www.hfbpw.top
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 403 Forbidden
< Server: AliyunOSS
< Date: Wed, 17 Nov 2021 11:22:43 GMT
< Content-Type: application/xml
< Content-Length: 253
< Connection: keep-alive
< x-oss-request-id: 6194E603D0409B36334D3CD8
< x-oss-server-time: 0
<
{ [253 bytes data]
100   253  100   253    0     0    928      0 --:--:-- --:--:-- --:--:--   973
* Connection #0 to host oss-cn-hongkong.aliyuncs.com left intact




```