<div align="center">
    <img src="https://cdn.jsdelivr.net/gh/krlc/bravo@master/icon-misc/bravo.svg" alt="Icon" width="300" />
  <br><br>
  <p>System tray GUI for <a href="https://www.wireguard.com/">WireGuard</a>  <a href="https://travis-ci.com/krlc/bravo"><img alt="Build Status" src="https://travis-ci.com/krlc/bravo.svg?branch=master"></a></p>
  <p>A whiff of a normal GUI. Quite practical, though.</p>
</div>

---

## Why?

Coz it's fun!


... and, as of Feb 1, 2019, there is still no WireGuard GUI for macOS and Linux (and Windows, when the WG client will be released).

Struggling every signle time and asking myself a question: _"is WireGuard running?"_, I told myself: _"__That's enough!__ No more terminal, sudo, wg and wg-quick!"_ and created a small yet neat menu to see WireGuard connection status, control it and see the stats.

## How would I get it?

First, make sure you have a working WireGuard setup and  have `wg-quick` tool installed.

### macOS (and Windows, as soon as WireGuard team releases a client for it)

Currently, no `brew` or `apt-get` is available, so you'll be needing to compile it by yourself.

- Go get [Go](https://golang.org/dl/)
- Then, clone it: `git clone https://github.com/krlc/bravo.git bravo`
- Compile it: `make build`
- And finally, `sudo ./bravo`

Straight after you hit enter to execute the last step, `bravo` icon should appear on the system tray, as shown below:

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/krlc/bravo@readme/icon-misc/screen-1.png" alt="System tray screenshot" width="300" />
</div>

### Linux

Install this: `sudo apt-get install libgtk-3-dev libappindicator3-dev` and follow the macOS setup.

<br>

Despite the fact that it is running (hopefully), you mighta noticed that the terminal should remain opened in order for `bravo` to work, which is ~~disgusting~~ tremendously disgusting! 
Rest assured, though, the help is on its way.

## How would I use it?

## Contributing to Hoaxwagen

Contributions are always welcome. There's no code of conduct, so feel free to do whatever you want, yeah!

**But don't forget to give it a star!**
