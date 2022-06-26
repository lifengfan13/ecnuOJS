#!/bin/bash
# docker rmi hxc/mysql:5.7
# docker buildx build --platform=linux/amd64 . -t hxc/mysql:5.7

docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=paper123456 hxc/mysql:5.7

# http://blog.java110.com/docker-mysql-start-init-table/ 参考
