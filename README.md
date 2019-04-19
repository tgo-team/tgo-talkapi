# tgo-talkapi
tgo talk对应的api

protoc --go_out=. *.proto

## 运行数据库
docker run --name mysql -p 3306:3306 -v /work/tmp/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=a123456 -e MYSQL_DATABASE=test -d mysql
##  redis
docker run -p 6379:6379 -d redis