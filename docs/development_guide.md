# Development guide

References:
- https://www.zabbix.com/documentation/guidelines/en/plugins/loadable_plugins

## VS Code - Developing inside a Container

```
mkdir -p zabbix-agent2-plugins/.devcontainer
cd zabbix-agent2-plugins
vim .devcontainer/devcontainer.json
```

## Desenvolvendo o plugin

Criando diretório do plugin
```bash
mkdir -p cmd/dirsearch
mkdir -p build/bin
```

Criando o arquivo main.go

```
touch cmd/dirsearch/main.go
```

Inicializando o módulo

```bash
# cd cmd/dirsearch
go mod init github.com/zabbix-agent2-plugin-dirsearch
GOPROXY=direct go get golang.zabbix.com/sdk/plugin@master
go mod tidy
go build -o build/bin/dirsearch cmd/dirsearch/main.go
GOOS=linux GOARCH=arm64 go build -o build/bin/dirsearch-linux-arm64 cmd/dirsearch/main.go
GOOS=linux GOARCH=amd64 go build -o build/bin/dirsearch-linux-amd64 cmd/dirsearch/main.go
```