# s

> Maps, videos, images, and YouTube from the command line

## Install

Put the program `bin/s` somewhere in your `$PATH`, e.g.

```
$ curl -s https://raw.githubusercontent.com/yangmillstheory/s/master/bin/s -o /usr/local/bin/s
$ chmod +x /usr/local/bin/s
```
## Usage

```
👌 ~/c/g [cc7a0af] (master)
(i) s -help
Usage of s:
  -d    Enable debug logging.
  -i    Search Google Images.
  -m    Search Google Maps.
  -v    Search Google Videos.
  -w    Search Google. (default true)
  -yt
        Search YouTube.
```

### Shell

```
# images, videos, youtube, web
$ s -i -v -y '"Black Mirror Season 4" review'

# maps, no web search
$ s -m -w=false pizza
```

### Vim

Put [this](s.vim) in your `.vimrc`.

Now you can search from `vim` in visual, normal, and operator-pending modes :).

## Development

1. Edit `s.go`.
1. `make build`.
1. Play with `bin/s`.

## FAQ

> What about [`googler`](https://github.com/jarun/googler)?

It's nice, but did too much for what I wanted.

## Limitations

The program is Darwin-only. This can be fixed by using [pkg/browser](https://godoc.org/github.com/pkg/browser).
