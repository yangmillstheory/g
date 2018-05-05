## g

> Search Google Maps, Videos, Images, or YouTube from the command line

### Install

Clone the repository and put `bin/g` somewhere in your `$PATH`, e.g.

```
$ curl -s https://raw.githubusercontent.com/yangmillstheory/g/master/bin/g -o /usr/local/bin/g
$ chmod +x /usr/local/bin/g
```
### Usage

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

#### Shell

```
$ g -images -videos -youtube '"Black Mirror Season 4" review'
```

```
$ g -maps -noweb pizza
```

#### Vim

Put the following in your `~/.vimrc`.

```vim
function! s:Google(type)
  let saved = @@
  if a:type ==# 'V'
    normal! `<V`>y
  elseif a:type ==# 'v'
    normal! `<v`>y
  elseif a:type ==# 'char'
    normal! `[v`]y
  else
    return
  endif
  " remove trailing newline
  "
  "   https://stackoverflow.com/a/6228454
  let @@ = substitute(strtrans(@@), '\^@', '', 'g')
  call system('g ' . shellescape(@@))
  let @@ = saved
endfunction

vnoremap <F2> :<c-u>call <SID>Google(visualmode())<cr>
nnoremap <F2> :set operatorfunc=<SID>Google<cr>g@
nnoremap <F2><F2> :silent !g <c-r><c-w><cr>
```

Now you can search from `vim` in visual, normal, and operator-pending modes :).

### FAQ

> What about [`googler`](https://github.com/jarun/googler)?

It's nice, but did too much for what I wanted.
