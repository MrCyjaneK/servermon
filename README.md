# ServerMon

Run a commands and display them in table...

create a file named `~/.config/servermon.json` (or `config.json` in working directory)

```json
{
    "entries": [
        [
            ["echo", "Name"],
            ["echo", "IP"],
            ["echo", "Uptime"],
            ["echo", "Disk"]
        ],
        [
            ["echo", "server.example"],
            ["ssh", "server.example", "curl", "--silent", "ifconfig.co"],
            ["ssh", "server.example", "uptime -p | sed 's/,.*//g' | sed 's/..........up//g'"],
            ["ssh", "server.example", "df -h | grep 'G' | grep -v /run | grep -v tmpfs | grep -v udev"]
        ],
        [
            ["echo", "server.example2"],
            ["ssh", "root@server.example2", "curl", "--silent", "ifconfig.co"],
            ["ssh", "root@server.example2", "uptime -p | sed 's/,.*//g' | sed 's/..........up//g'"],
            ["ssh","root@server.example2","df -h | grep 'G' | grep -v /run | grep -v tmpfs | grep -v udev"]
        ]
    ]
}
```

That will display a table like this:

| Name | IP | Uptime | Disk |
| ---- | -- | ------ | ---- |
| server.example  | 127.0.0.1 | up 7 weeks | *output from df in multiline* |
| server.example2 | 127.0.0.2 | up 2 weeks | *output from df in multiline* |