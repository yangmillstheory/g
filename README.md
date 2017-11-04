# g

> Google from the command line

![g](https://user-images.githubusercontent.com/2729079/32401879-9047a904-c0d5-11e7-9e27-8b51fa4d7f7b.gif)

### Install

Clone the repository and put `bin/g` somewhere in your `$PATH`.

### Usage

```
# preserve double-quotes and run in verbose mode
$ g -v '"Black Mirror" review'

# issue a standard search query in non-verbose mode
$ g Stranger Things 2
```

### FAQ

> Will this ever be published on PIP or Homebrew?

If there's interest or I'm bored enough, sure.

For now I just sync this into `$HOME/.bin` - which lies in my `$PATH` - manually using a script in my dotfiles repository.

> What about [`googler`](https://github.com/jarun/googler)?

It's nice, but did too much for what I wanted.
