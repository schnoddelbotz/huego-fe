# huego-fe

A cross-platform CLI, GUI and Web - frontend for Philips Hue bridges, based on [huego](https://github.com/amimof/huego).

`huego-fe` focuses on basic lights operations (on, off, brightness, ... more to come potentially).

## CLI usage

```bash
$ huego-fe
huego-fe can control your philips hue stuff

Usage:
  huego-fe [command]

Available Commands:
  brightness  control gravity
  color       language agnostic eye pleasures
  help        Help about any command
  list        A brief description of your command
  login       Discover Hue bridge and log in -- press link button first!
  off         fusion reactor control plane
  on          engage rocket launcher
  serve       runs the thing that philps frogot on the Hoe
  version     prints a bestseller novel on-demand

Flags:
      --config string     config file (default is $HOME/.huego-fe.yaml)
  -h, --help              help for huego-fe
  -i, --hue-ip string     Hue bridge IP [$HUE_IP] , see: huego-fe login -h
  -l, --hue-light int     Hue light No.# [$HUE_LIGHT], see: huego-fe list (default 1)
  -u, --hue-user string   Hue bridge user/token [$HUE_USER], see: huego-fe login -h

Use "huego-fe [command] --help" for more information about a command.
```


## Installation

If you have [Go](https://golang.org/doc/install) installed and want to build from current master:

```bash
go get github.com/schnoddelbotz/huego-fe
```

Otherwise, [download a binary release](./../../releases) and put the binary somewhere on your `PATH`.

## Setup / Usage / Examples

At first run after installation, link your Hue. Web and CLI can currently be used to login / link.
Hue address and login data will be stored in `~/.huego-fe.yml`. Should you ever want to re-link,
delete the file.

### Web

- Run `huego-fe serve --open-browser` (or `huego-fe s -o` for short). 

Your browser should open, showing huego-fe web UI, asking you to push link button. Once pressed, 
you should be warped into control UI.

### CLI

- Press Hue's link button to enable login
- Run `huego-fe login` once
- Try `huego-fe list; huego-fe on; huego-fe b 64; huego-fe 0` etc.

### GUI

- just run `huego-fe`

UI WIP/working status / keyboard shortcuts:

- [ ] Up/Dn: Select Light
- [x] PgUp/Home: Power on
- [x] PgDn/End: Power off
- [x] Enter/Return: Toggle selected light's state
- [x] Left/Right: -/+ brightness 20
- [x] Ctrl-Left/Right: -/+ brightness 10
- [x] Shift-Left/Right: -/+ brightness 1
- [x] Alt-Left/Right: brightness jump min/max
- [x] Space: Toggle selected light's state and quit
- [x] ESC: Quit
- [ ] Delayed / Dimmed On/Off?

As long as GUI light selection is missing, start `huego-fe` with `-l ...` to override default from `~/.huego-fe.yml`.

#### Desktop integration

It might be handy to assign a Keyboard shortcut to start `huego-fe` GUI for regular use. 
Example setup for Gnome / Ubuntu 20.04:

- Go to settings > Keyboard shortcuts, scroll to bottom, hit `+`
- Given you put `huego-fe` into `$PATH` during installation, just use `huego-fe` here as Name and Command
- Click `Set Shortcut` and e.g. choose/press Ctrl-F12

Pressing Ctrl-F12 will now bring up `huego-fe` with default `hue-light` as set in `~/.huego-fe.yml`!

You may want to additionally assign `huego-fe toggle` (to e.g. Ctrl-Shift-F12), permitting direct toggling
of your default lamp.

# todo

- pairing via GUI
- enable cobra shell auto-completion on commands / lights
- changing brightness via slider does not update brightness (only kbd works)
- add a cmd/install_linux.go that permits simple installation of systemd socket-activated `huego-fe serve`?
- use index.tpl.html for link process, too
- numeric lamp id vs name ... usage issues? id stability?
- split gui and cli/web binaries? build time for CLI/web only usage concerns + mousetrap breaks cli on win 

# kudos to ...

- [huego](https://github.com/amimof/huego) -- for making building `huego-fe` on top of it a simple joy
- [Gio](https://gioui.org/) -- for enabling `huego-fe` GUI
- [Cobra](https://cobra.dev/) -- for rocking `huego-fe` CLI

# bugs

Plenty for sure - have you seen a single test in here? Lame excuse: Toy project. Still, issues / PRs very welcome.
