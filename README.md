# README.md

- Development Guide: [click here](docs/development_guide.md)

## Plugin description

This plugin was developed with the aim of allowing you to search directories based on a regular expression.
The initial motivation was to resolve a known issue regarding log monitoring.

To contextualize the problem, I will give the example of two item keys:
- log
    ```
    log[file,<regexp>,<encoding>,<maxlines>,<mode>,<output>,<maxdelay>,<options>,<persistent dir>]
    ```
- logrt
    ```
    logrt[file regexp,<regexp>,<encoding>,<maxlines>,<mode>,<output>,<maxdelay>,<options>,<persistent dir>]
    ```

Regarding the use of regex, we can use regex only on the content received by the log in both keys or only on the file name according to the documentation

`file regexp - the absolute path to file and the file name described by a regular expression. Note that only the file name is a regular expression.`

Reference: https://www.zabbix.com/documentation/current/en/manual/config/items/itemtypes/zabbix_agent#log

### Supported keys

```
dir.search[dir_scan, regexp]
```

Parameters:
- dir_scan: The absolute path of the directory you want to search for subdirectories
- regexp: The regular expression that describes the required pattern

Examples

- dir.search["/var","zabbix$"]

In this example we have the directory tree in /var and inside /var/log we have zabbix

```bash
├── local
├── lock -> ../run/lock
├── log
│   ├── anaconda
│   ├── audit
│   ├── btmp
│   ├── chrony
│   ├── cloud-init.log
│   ├── cloud-init-output.log
│   ├── cron
│   ├── dnf.librepo.log
│   ├── dnf.log
│   ├── dnf.rpm.log
│   ├── droplet-agent.update.log
│   ├── hawkey.log
│   ├── kdump.log
│   ├── lastlog
│   ├── messages
│   ├── private
│   ├── qemu-ga
│   ├── README -> ../../usr/share/doc/systemd/README.logs
│   ├── secure
│   ├── sssd
│   ├── tallylog
│   ├── wtmp
│   └── zabbix
```

The regex match will occur if it is a directory and if you include zabbix in the name, in this case **/var/log/zabbix**

For more examples visit [the examples folder](examples)

## How to use

### Instale o Go no servidor

RHEL based

```bash
dnf install golang-go -y
```

### Clone o repositório

```bash
dnf install -y git
git clone git@github.com:robertsilvatech/zabbix-agent2-plugin-dirsearch.git
```

### Crie um diretório para armazenar plugins do Zabbix Agent 2

```bash
mkdir -p /etc/zabbix/external_plugins
```

### Copie o binario de acordo com o seu s.o e sua arquitetura

```bash
ls zabbix-agent2-plugin-dirsearch/build/bin/
# Output
dirsearch  dirsearch-linux-amd64  dirsearch-linux-arm64
```

```bash
cp zabbix-agent2-plugin-dirsearch/build/bin/dirsearch-linux-amd64 /etc/zabbix/external_plugins
ls -l /etc/zabbix/external_plugins
```

### Crie o arquivo de configuração do plugin

```bash
echo 'Plugins.DirSearch.System.Path=/etc/zabbix/external_plugins/dirsearch-linux-amd64' > /etc/zabbix/zabbix_agent2.d/plugins.d/dirsearch.conf
```

### Teste a chave do item

```bash
zabbix_agent2 -t dir.search["/var/log","zabbix$"]
```

### Reinicie o Agent 

```bash
systemctl restart zabbix-agent2
```