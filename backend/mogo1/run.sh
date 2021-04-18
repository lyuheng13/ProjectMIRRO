#docker kill mongoone
docker rm mongoone
docker pull niupi/mongoone

docker run -d \
-p 27137:27137 \
-e MONGO_INITDB_ROOT_USERNAME=niupi \
-e MONGO_INITDB_ROOT_PASSWORD=@NIUPI123 \
--name mongoone \
--network niupi \
mongo
