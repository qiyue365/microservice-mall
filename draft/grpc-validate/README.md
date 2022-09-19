## 安装插件

```bash
go install github.com/envoyproxy/protoc-gen-validate@latest
```

## 下载 proto 文件

```bash
mkdir validate && \
cd validate && \
wget -4 https://raw.githubusercontent.com/envoyproxy/protoc-gen-validate/main/validate/validate.proto
```

## 生成代码

```bash
protoc -I . --go_out=./pb --go-grpc_out=./pb --validate_out="lang=go:./pb" *.proto
```
