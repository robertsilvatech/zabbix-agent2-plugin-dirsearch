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