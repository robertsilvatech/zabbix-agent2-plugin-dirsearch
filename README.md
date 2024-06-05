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

Parametros:
- dir_scan: The absolute path of the directory you want to search for subdirectories
- regexp: The regular expression that describes the required pattern

Examples

```
di
```