## g

> Search Google Maps, Videos, Images, or YouTube from the command line

### Install

Clone the repository and put `bin/g` somewhere in your `$PATH`, e.g.

```
$ curl -s https://raw.githubusercontent.com/yangmillstheory/g/master/bin/g -o /usr/local/bin/g
$ chmod +x /usr/local/bin/g
```
### Usage

```
👌 ~/c/g [65012fd] (master⚡)
(i) g -help
Usage of ./bin/g:
  -d    Enable debug logging.
  -i    Search Google Images.
  -m    Search Google Maps.
  -v    Search Google Videos.
  -w    Search Google. (default true)
  -yt
        Search YouTube.
```

#### Shell

```
# images, videos, youtube, web
$ g -i -v -y '"Black Mirror Season 4" review'

# maps, no web search
$ g -m -w=false pizza
```

#### Vim

Put [this](g.vim) in your `.vimrc`.

Now you can search from `vim` in visual, normal, and operator-pending modes :).

### Development

1. Edit `g.go`.
1. `make build`.
1. Play with `bin/g`.

### FAQ

> What about [`googler`](https://github.com/jarun/googler)?

It's nice, but did too much for what I wanted.
