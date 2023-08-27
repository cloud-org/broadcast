# broadcast

etcd 配置变更，agent 广播器实现

## arch

![image](https://user-images.githubusercontent.com/28869910/134543364-f94cbd1d-18b5-4817-8529-eed160b58aba.png)

## run

```go
go run main.go
cd notify && go run notify.go
```

![image-20210722012303177](img/image-20210722012303177.png)

![image-20210722012307043](img/image-20210722012307043.png)

### e2e 

deploy etcd frist, and then run `make e2e`.

## acknowledgement

- etcd
