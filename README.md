# aw-timer ⏱️

Simple command line timer app that uses [ActivityWatch](https://github.com/activitywatch/activitywatch) for data persistence.

## Usage

```bash
Usage: aw-timer [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  current     Print specified fields
  help        Help about any command
  start       Start a task
  stop        stop a task
  summery     Print specified fields

Flags:
  -h, --help   help for aw-timer

Use "aw-timer [command] --help" for more information about a command.
```

## Installation

```bash
git clone https://github.com/mosaleh52/aw-timer
cd aw-timer
go build
sudo ln -s $(pwd)/aw-timer /usr/local/bin/aw-timer
```

## Commands

### `start`

Start a task.

```bash
Usage: aw-timer start [flags]

Flags:
  -u, --api-url string       specify the api url  (default "http://127.0.0.1:5600/api/0/")
  -b, --bucket-id string     specify the bucket-id (default "aw-stopwatch")
  -d, --date-Layout string   specify the dateLayout used for formatting in aw server (default "2006-01-02T15:04:05.999999-07:00")
  -h, --help                 help for start
```

### `stop`

Stop a task.

```bash
Usage: aw-timer stop [flags]

Flags:
  -u, --api-url string       specify the api url  (default "http://127.0.0.1:5600/api/0/")
  -b, --bucket-id string     specify the bucket-id (default "aw-stopwatch")
  -d, --date-Layout string   specify the dateLayout used for formatting in aw server (default "2006-01-02T15:04:05.999999-07:00")
  -h, --help                 help for stop
```

### `current`

Print specified fields.

```bash
Usage: aw-timer current [flags]

Flags:
  -u, --api-url string       specify the api url  (default "http://127.0.0.1:5600/api/0/")
  -b, --bucket-id string     specify the bucket-id (default "aw-stopwatch")
  -c, --color string         specify the coloring method form [normal , i3 , none] default to normal  (default "term")
  -d, --date-Layout string   specify the dateLayout used for formatting in aw server (default "2006-01-02T15:04:05.999Z")
  -h, --help                 help for current
  -r, --require string       specifie the required atter from [uuid,id,label] (default "all")
```

## i3 blocks

```
[CurrentTodoBlock]
command=aw-timer current -c "i3" -r label
interval=15
```

## Todo

- [ ] handle stop while toggle
- [ ] refactor settings to diffrent file allowing config file
- [ ] allow for toggling for period
- [ ] add option for demanding uuid check or generate one random
- [ ] add tray to control from gui with timer on it
