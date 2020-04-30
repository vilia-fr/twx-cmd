**THIS IS THE EARLY WORK-IN-PROGRESS, NOTHING IS FUNCTIONAL HERE YET**

# Unofficial ThingWorx CLI

Command-line utilities to simplify your daily ThingWorx ops

Twx-cmd is a set of Bash scripts that work on Linux (obviously), Windows Subsystem for Linux (WSL) and MinGW (e.g. git-bash on Windows).

They communicate with ThingWorx over HTTP, so you can use them with your local installation, Docker containers, cloud instances and whatnot. This implies
that your ThingWorx server is online.

## Installation

There are three prerequisites for twx-cmd: [bash(1)](https://linux.die.net/man/1/bash), [zip(1)](https://linux.die.net/man/1/zip) and [curl(1)](https://linux.die.net/man/1/curl).
Zip is only required if you are going to deploy ThingWorx sources on a running server.

Twx-cmd shell scripts call each other, so you shouldn't try to delete some of them. The most reliable way to run them is by adding them to your $PATH, or executing
directly from the directory where you clone / unzip this repository.

### Installation on Windows

There are different ways to execute Bash scripts on Windows, but two of the most common ones are:

1. Enable [Windows Subsystem for Linux (WSL)](https://docs.microsoft.com/en-us/windows/wsl/install-win10), install Linux distro and follow instructions below;
2. Install MinGW and add zip to your PATH. Command-line Git client already includes MinGW, so you can simply use %GIT_HOME%\bin\bash.exe

### Installation on Linux (RPM)

Although you most probably already have those tools, try this: `yum install bash zip curl`

### Installation on Linux (DEB)

Although you most probably already have those tools, try this: `apt install bash zip curl`

## Usage

### Common parameters and environment variables

All twx-cmd scripts share some subset of parameters. Those can be set either as environment variables, or specified directly on the command line:

1. $TWXURL / --twx-url=URL
2. $TWXAPPKEY / --twx-appkey=APPKEY
3. $TWXUSER / --twx-user=USERNAME
4. $TWXPASSWORD / --twx-password=PASSWORD

### Execute services

twx-exec --service-url= [--timeout=] [[p1=v1]]

twx-exec-foreach FILENAME.csv SERVICEURL [--delimiter=,]

twx-exec-infotable FILENAME.csv SERVICEURL [--param=input] [--delimiter=,]

twx-exec-json FILENAME.csv SERVICEURL [--param=input] [--element=array] [--delimiter=,]

### Import / export sources

twx-src-export PROJECT [--strip] [--start-date=`date`]

twx-src-download PROJECT [--strip] [--start-date=`date`]

twx-src-backup PROJECT [--strip] [--start-date=`date`]

twx-src-import

### Working with extensions

twx-ext-import

twx-ext-list

### Working with File Repositories

twx-files-mkdir

twx-files-rm

twx-files-ls

twx-files-download [--zip]

twx-files-upload [--unzip]

### Querying data

twx-query-things [--entity-type=Thing] [--properties=*] [--tags=] [--thing-template=] [--thing-shape=] [--project=]

twx-query-users [--entity-type=Thing] [--properties=*] [--tags=] [--thing-template=] [--thing-shape=] [--project=]

twx-query-data --name [--filter=p1:v1,p2:v2,p3:v3]

twx-query-stream --name [--source=] [--source-type=] [--tags=] [--date-from=] [--date-to=] [--fields=*]

### Monitoring platform state and health

twx-state-ping [--timeout=]

twx-state-waitboot [--timeout=]

twx-state-waitshutdown [--timeout=]

twx-state-summary [--timeout=]

### Working with logs

twx-log-configure [--type=ApplicationLog] [--level=INFO]

twx-log-query [--type=ApplicationLog] [--level=INFO] [--date-from=] [--date-to=] [--origin=] [--substr=]

### Working with users

twx-users-active

twx-users-disable

twx-users-enable

### External integrations

twx-snapshot

twx-diff

twx-patch

twx-docs --project=

## Building

Building twx-cmd is trivial. Firstly you need to install the latest version of [Go](https://golang.org/), then build `main.go`:

```bash
cd src
go build main.go
```

## License