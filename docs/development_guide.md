# Development guide

References:
- https://www.zabbix.com/documentation/guidelines/en/plugins/loadable_plugins

## VS Code - Developing inside a Container

```bash
mkdir -p zabbix-agent2-plugins/.devcontainer
cd zabbix-agent2-plugins
vim .devcontainer/devcontainer.json
```

Adicionar conteúdo

```json
{
    "name": "My dev environment",
    "image": "robertsilvatech/my-dev-container:1.0.3",
	"containerEnv": {
		"GITHUB_TOKEN": "${localEnv:GITHUB_TOKEN}",
		"GITHUB_USER": "${localEnv:GITHUB_USER}",
		"GOPROXY": "${localEnv:GOPROXY}" , 
		"HTTP_PROXY": "${localEnv:HTTP_PROXY}" , 
		"HTTPS_PROXY": "${localEnv:HTTPS_PROXY}"
	},	
	"remoteEnv": {
		"GITHUB_TOKEN": "${localEnv:GITHUB_TOKEN}",
		"GITHUB_USER": "${localEnv:GITHUB_USER}",
		"GOPROXY": "${localEnv:GOPROXY}" , 
		"HTTP_PROXY": "${localEnv:HTTP_PROXY}" , 
		"HTTPS_PROXY": "${localEnv:HTTPS_PROXY}"
	}	
}
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

Construindo o módulo do plugin

```bash
go mod init github.com/zabbix-agent2-plugin-dirsearch
GOPROXY=direct go get golang.zabbix.com/sdk/plugin@master
go mod tidy
go build -o build/bin/dirsearch cmd/dirsearch/main.go
GOOS=linux GOARCH=arm64 go build -o build/bin/dirsearch-linux-arm64 cmd/dirsearch/main.go
GOOS=linux GOARCH=amd64 go build -o build/bin/dirsearch-linux-amd64 cmd/dirsearch/main.go
```

## Testando o plugin no dev container com Zabbix Agent 2

> Attention, check the branch you are using in the dev container

```bash
apt update
apt install -y git
git branch
```

### Remover versao antiga do go

```bash
rm $(which go version)
```

### Instalar nersao mais recente do go

```bash
wget https://go.dev/dl/go1.22.3.linux-arm64.tar.gz && tar -C /usr/local -xzf go1.22.3.linux-arm64.tar.gz && export PATH=$PATH:/usr/local/go/bin
rm go1.22.3.linux-arm64.tar.gz
go version
```

### Instalar o Zabbix agent 2

```bash
apt install procps -y && \
wget https://repo.zabbix.com/zabbix/7.0/ubuntu-arm64/pool/main/z/zabbix/zabbix-agent2_7.0.0-1%2Bubuntu22.04_arm64.deb && \
dpkg -i zabbix-agent2_7.0.0-1+ubuntu22.04_arm64.deb && \
/etc/init.d/zabbix-agent2 restart && \
/etc/init.d/zabbix-agent2 status
rm *.deb
```

### Carregar o plugin para o Zabbix agent 2

```bash
echo 'Plugins.DirSearch.System.Path=/workspaces/zabbix-agent2-plugins/build/bin/dirsearch-linux-arm64' > /etc/zabbix/zabbix_agent2.d/plugins.d/dirsearch.conf
```

### Testar a chave do plugin

```bash
zabbix_agent2 -t dir.search["/var","zabbix$"]
```

Output

```bash
dir.search[/var,zabbix$]                      [s|[{"name":"/var/log/zabbix"}]]
```