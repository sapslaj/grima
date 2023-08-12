# Grima

Kills Chrom(e)

Grima lurks in the (systemd) shadows and if detects a process that might be Chrom(e) it tries to kill it. If it fails, it brings the whole world down with it.

## Installation

### Linux

1. Download (or build) the `grima` executable and place it somewhere like `/usr/local/bin`
2. Copy `grima.service` to `/usr/lib/systemd/system/grima.service` (modifying the path to `grima` as needed)
3. Run `systemctl daemon-reload && systemctl start --now grima.service`

### macOS

idk something with plists and launchd

### Windows

lol.

## Contributing

It's not a complicated program (for now) so if you find cases where it detects a false positive or it misses a case (of which there are probably many) PRs are always welcome.

## Q and A

### Why

because I'm trying to not use Google Chrome or anything based on it as a sort of "personal challenge" and I thought the idea was funny.

### IS tHAT A FIRE EMBLEM REFERENCE

yeah but i didn't finish awakening though so the joke is ether funnier or not funny depending on your perspective

### what do you use instead of chrome

firefox and lynx (on arch btw)

### The joke isn't that funny

that's not a question but yes I know and no you're wrong.

### Why is it written in Go and not Rust

because my hair isn't blue enough

### Why is it written in Go and not `case language:`

because `${joke[language]}`

### Didn't Google orignally create Go?

![ironic.](https://i.kym-cdn.com/photos/images/newsfeed/001/578/199/124.gif)
