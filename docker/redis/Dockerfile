FROM redis:latest
LABEL authors="27678"

VOLUME /var/lib/redis

EXPOSE 6379

CMD ["redis-server"]

# docker run -d \
  #  -v D:/github_project/bluebell/docker/redis/redis-data:/var/lib/redis \
  #  --name my-redis-container \
  #  -p 6379:6379 \
  #  bluebell-redis
#  docker run -d -v D:/github_project/bluebell/docker/redis/redis-data:/var/lib/redis --name my-redis-container -p 6379:6379 bluebell-redis