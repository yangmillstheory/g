## g

> Google from the command line

![g](https://user-images.githubusercontent.com/2729079/32401909-6683f82e-c0d6-11e7-9d29-7301f84c0ef8.gif)

### Requirements

Python 3. Well, not really, but why not?

```
$ brew install python3
$ which python3
```

### Install

Clone the repository and put `bin/g` somewhere in your `$PATH`.

### Usage

From the command line:

```
# preserve double-quotes and run in verbose mode
$ g -v '"Black Mirror" review'

# issue a standard search query in non-verbose mode
$ g Stranger Things 2
```

Use in Vim (Visual, Normal, or Operator-pending modes supported):

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

### FAQ

> Will this ever be published on PIP, Homebrew, or as a Vim plugin?

If there's interest or I'm bored enough, sure.

For now I just sync this into `$HOME/.bin` - which lies in my `$PATH` - manually using a script in my dotfiles repository.

> What about [`googler`](https://github.com/jarun/googler)?

It's nice, but did too much for what I wanted.
