docker run -itd \
  --name redis \
  --network mirro \
  --expose=6379 \
  redis
