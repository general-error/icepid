# Icepid

Icepid is a very simple dashboard. It should be good to use on a small personal
server.

# Features

You can view the following data:

- Uptime, system time
- Available updates
- System logs
- free, w, df
- sensors
- SMART data

# Screenshots

![Login page](screenshots/icepid1.png?raw=true)
![Index page](screenshots/icepid2.png?raw=true)

# Build

Run this in `src` folder:

```sh
go build -o icepid *.go
go build icepid-smart-plugin/*.go
sudo chown root icepid-smart
sudo chmod u+s  icepid-smart
```

# Setup

Drop the icepid & icepid-smart binaries somewhere on your server, edit the
icepid.json settings file and let it run via cron/systemd/etc. `icepid.service`
file is included.

# License

Licensed under terms of GNU General Public License version 3.
