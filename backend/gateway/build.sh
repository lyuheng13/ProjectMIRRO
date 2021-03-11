GOOS=linux go build
docker build -t niupi/gateway:latest .
docker push niupi/gateway:latest
go clean