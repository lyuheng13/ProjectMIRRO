docker pull niupi/mysql
docker rm mysql

docker run -d \
-p 3306:3306 \
--name mysql \
-e MYSQL_ROOT_PASSWORD=niupi123 \
-e MYSQL_DATABASE=niupiuser \
--network niupi \
niupi/mysql

#进入容器
#docker exec -it mysql bash

#登录mysql
#mysql -u root -p
#ALTER USER 'root'@'localhost' IDENTIFIED BY 'Lzslov123!';

#添加远程登录用户
#CREATE USER 'liaozesong'@'%' IDENTIFIED WITH mysql_native_password BY 'Lzslov123!';
#GRANT ALL PRIVILEGES ON *.* TO 'liaozesong'@'%';