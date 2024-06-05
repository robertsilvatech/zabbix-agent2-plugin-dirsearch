# README.md

- [README.md](#readmemd)
  - [Plugin description](#plugin-description)
    - [Supported keys](#supported-keys)
  - [How to use](#how-to-use)
    - [Requirements](#requirements)
    - [Install Go on the server running Zabbix Agent 2](#install-go-on-the-server-running-zabbix-agent-2)
    - [Clone the repository](#clone-the-repository)
    - [Create a directory to store Zabbix Agent 2 plugins](#create-a-directory-to-store-zabbix-agent-2-plugins)
    - [Copy the binary according to your OS and architecture](#copy-the-binary-according-to-your-os-and-architecture)
    - [Create the plugin configuration file](#create-the-plugin-configuration-file)
    - [Test the item key](#test-the-item-key)
    - [Restart the Agent](#restart-the-agent)
    - [Validate the plugin against Zabbix Agent 2 metrics](#validate-the-plugin-against-zabbix-agent-2-metrics)
  - [Development Guide](#development-guide)


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

Steps:
- Install Go on the server running Zabbix Agent 2
- Clone the repository
- Create a directory to store Zabbix Agent 2 plugins
- Copy the binary according to your OS and architecture
- Create the plugin configuration file
- Test the item key
- Restart the Agent
- Validate the plugin against Zabbix Agent 2 metrics

### Requirements

- Linux operating system
- Go lang installed on the server running zabbix agent 2

### Install Go on the server running Zabbix Agent 2

Example for RHEL-based distributions

```bash
dnf install golang-go -y
```

### Clone the repository

```bash
dnf install -y git
git clone https://github.com/robertsilvatech/zabbix-agent2-plugin-dirsearch.git
```

### Create a directory to store Zabbix Agent 2 plugins

```bash
mkdir -p /etc/zabbix/external_plugins
```

### Copy the binary according to your OS and architecture

```bash
ls zabbix-agent2-plugin-dirsearch/build/bin/
# Output
dirsearch  dirsearch-linux-amd64  dirsearch-linux-arm64
```

```bash
cp zabbix-agent2-plugin-dirsearch/build/bin/dirsearch-linux-amd64 /etc/zabbix/external_plugins
ls -l /etc/zabbix/external_plugins
```

### Create the plugin configuration file

```bash
echo 'Plugins.DirSearch.System.Path=/etc/zabbix/external_plugins/dirsearch-linux-amd64' > /etc/zabbix/zabbix_agent2.d/plugins.d/dirsearch.conf
```

### Test the item key

```bash
zabbix_agent2 -t dir.search["/var/log","zabbix$"]
```

Output

```bash
dir.search[/var/log,zabbix$]                  [s|[{"name":"/var/log/zabbix"}]]
```

### Restart the Agent

```bash
systemctl restart zabbix-agent2
```

### Validate the plugin against Zabbix Agent 2 metrics

```bash
zabbix_agent2 -R metrics | grep DirSearch -A 6
```

Output
```bash
[DirSearch]
active: false
path: /etc/zabbix/external_plugins/dirsearch-linux-amd64
capacity: 0/1000
check on start: 0
tasks: 0
dir.search: Returns a json with the list of directories.
```

Ready, create a master item and LLD with dependent item to monitor your logs

## Development Guide

- Development Guide: [click here](docs/development_guide.md)