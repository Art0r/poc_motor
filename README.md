# POC Motor

## Running App

```sh
docker rm -f $(docker ps -a | grep 'poc_motor' | awk '{print $1}')
docker compose down && docker compose up -d --build
```

## Monitor Consumer

```sh
docker logs $(docker ps -a | grep 'consumer' | awk '{print $1}') --follow
```
