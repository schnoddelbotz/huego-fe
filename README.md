# huego-fe

A more or less thin cross-platform frontend / wrapper around [huego](https://github.com/amimof/huego).
huego-fe only provides basic lights operations to the user (on, off, brightness, ... more to come):

- [x] using [Cobra](https://cobra.dev/) for a sleek CLI interface 
- [x] using a web browser
- [ ] using a native UI (tray icon?)

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

It might be handy to assign a Keyboard shortcut to start huego-fe for later use. 
Would let you switch a light with 2 keystrokes. For Gnome:

- Go to settings, Keyboard shortcuts
- Assign key to launch `huego-fe`

- [ ] UI is tiny and optimized for quick keyboard control. Controls:
    - [ ] Up/Dn: Select Light
    - [ ] Enter: Toggle selected light's state and quit (default-selected button)
    - [ ] Left/Right: -/+ brightness
    - [ ] Delayed / Dimmed On/Off?

# bugs

Plenty for sure - have you seen a single test in here?
Lame excuse: Toy project. Still, issues / PRs very welcome.
