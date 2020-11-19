# huego-fe

A cross-platform CLI, GUI and Web - frontend for [huego](https://github.com/amimof/huego).
huego-fe only provides basic lights operations to the user (on, off, brightness, ... more to come):

- [x] using [Cobra](https://cobra.dev/) for a sleek CLI interface 
- [x] using a web browser
- [x] using [Gio](https://gioui.org/) for a native UI (inclomplete/wip)

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

If you have Go installed and want a build from current master:

```bash
go get github.com/schnoddelbotz/huego-fe
```

Otherwise, [download a binary release](./../../releases) and put the binary somewhere on your `PATH`.

## Setup / Usage / Examples

At first run after installation, link your Hue. UI, Web and CLI -- all can be used to Login / Link.
Hue address and login data will be stored in `~/.huego-fe.yml`. Should you ever want to re-link,
delete the file.

### Web

- Run `huego-fe serve --open-browser` (or `huego-fe s -o` for short). 

Your browser should open, showing huego-fe web UI, asking you to push link button. Once pressed, 
you should be warped into control UI.

<!-- TODO:
It's well imaginable to start webserver at boot, login ... or via socket activation.
Examples might follow here.
Just remember it does zero authentication ... [yet?] - anybody on your network will have full lights control :scream:!
-->

### CLI

- Press Hue's link button to enable login
- Run `huego-fe login` once
- Try `huego-fe list; huego-fe on; huego-fe b 64; huego-fe 0` etc.

### UI -- WIP / NOT-HERE-YET / Idea ...

- just run `huego-fe`

It might be handy to assign a Keyboard shortcut to start huego-fe for regular use. 
Example setup for Gnome / Ubuntu 20.04:

- Go to settings > Keyboard shortcuts, scroll to bottom, hit `+`
- Given you put `huego-fe` into `$PATH` during installation, just use `huego-fe` here as Name and Command
- Click `Set Shortcut` and e.g. choose/press Ctrl-F12
- Pressing Ctrl-F12 will now bring up `huego-fe` with default `hue-light` as set in `~/.huego-fe.yml`  

UI WIP/working status:
- [x] UI is tiny and optimized for quick keyboard control
- [ ] Up/Dn: Select Light
- [ ] Enter: Toggle selected light's state and quit (default-selected button)
- [x] Left/Right: -/+ brightness
- [x] ESC: Quit
- [ ] Delayed / Dimmed On/Off?

# bugs

Plenty for sure - have you seen a single test in here?
Lame excuse: Toy project. Still, issues / PRs very welcome.
