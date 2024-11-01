[![Build](https://github.com/dchansen06/sysapid/actions/workflows/build.yml/badge.svg)](https://github.com/dchansen06/sysapid/actions/workflows/build.yml)
# sysapid
Delivers compact JSON information so that you can integrate it into your local web apps

It is highly recommended to fork and modify the program to deliver the information you want. Installing out of box is possible but will not be super useful.

The syntax in the files is fairly straightforward, just add any new commands or other lines as you would expect. Make sure that all relevant binaries are installed as needed.

You may need to install [sensors(1)](https://linux.die.net/man/1/sensors) see [this](https://help.ubuntu.com/community/SensorInstallHowto) for more information if that is your preferred method of collecting temperature data.

Note: As of now I do not know of any reliable method to get temperature data on Windows.

Access it by navigating to [http://HOSTNAME:PORT](http://localhost:8080), add `?fast=LABEL` (such as "cpu" to only retrieve one of the elements to be slightly quicker

## Installation
Install golang-go, then run:

```$ go get github.com/gin-gonic/gin```

```$ go install sysapid.go```

It will be installed into the default directory by your GOENV

## Security
This program connects to the network. Do not use on the open internet or on a network where anyone but you can access it. Use this program at your own risk and discretion.

## systemD
Should you want to allow controlling the file with the various C programs, setup the `sysapid.service` job unit file with the correct path and move it to `/etc/systemd/system/sysapid.service`

See [systemd](https://wiki.debian.org/systemd) or [systemd(1)](https://man7.org/linux/man-pages/man1/systemd.1.html) for more information
