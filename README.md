# Icepid

Icepid is a very simple dashboard. It should be good to use on a small personal
server.

# Screenshots

![Login page](screenshots/icepid1.png?raw=true)
![Index page](screenshots/icepid2.png?raw=true)

# Build

Run this in `src` folder:

```sh
go build -o icepid
```

# Setup

Drop the icepid binary somewhere on your server, edit the icepid.json settings
file and let it run via cron/systemd/etc.

# License

Licensed under terms of GNU General Public License version 3.
