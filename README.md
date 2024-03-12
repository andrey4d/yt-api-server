### Build
```shell
docker build -t registry.home.local/goytdl .
```

### RUN
```shell
docker run -it --rm -p 8081:8081 -v ./configs/apiserver.yaml:/configs/apiserver.yaml registry.home.local/goytdl
```