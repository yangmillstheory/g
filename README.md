> Google from the command line

<p align="center">
  <img src='https://user-images.githubusercontent.com/2729079/32401909-6683f82e-c0d6-11e7-9d29-7301f84c0ef8.gif'>
  <img src='https://user-images.githubusercontent.com/2729079/32401911-70d8b954-c0d6-11e7-81ad-bba7f38b3d97.gif'>
</p>

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
