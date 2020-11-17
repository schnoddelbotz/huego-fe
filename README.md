# huego-fe

A more or less thin frontend / wrapper around [huego](https://github.com/amimof/huego).
huego-fe only provides basic lights operations to the user (on, off, brightness, ...):

- [x] using [Cobra](https://cobra.dev/) for a sleek CLI interface 
- [ ] using a web browser
- [ ] using a native UI (tray icon?)

WIP. Just for fun.

## Usage:

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

### Setup/Example

#### CLI

- Run `touch ~/.huego-fe.yml` or `touch ~/.huego-fe.json` once; the file will store Hue login data 
- Press Hue's link button to enable login
- Run `huego-fe login` once
- Try `huego-fe list; huego-fe on; huego-fe b 64; huego-fe off` etc.
