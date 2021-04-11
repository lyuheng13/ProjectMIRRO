#docker network create niupi
docker rm niupigateway
docker pull niupi/niupigateway

docker run -d \
--name niupigateway \
-p 8080:80 \
--network niupi \
niupi/niupigateway

