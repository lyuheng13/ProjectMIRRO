GOOS=linux go build
docker build -t niupi/gateway:latest .
go clean