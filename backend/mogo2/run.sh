#docker kill mongoone
docker rm mongotwo
docker pull niupi/mongotwo

docker run -d \
-p 27138:27138 \
-e MONGO_INITDB_ROOT_USERNAME=niupi \
-e MONGO_INITDB_ROOT_PASSWORD=@NIUPI123 \
--name mongotwo \
--network niupi \
niupi/mongotwo
