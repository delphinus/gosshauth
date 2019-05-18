# `gosshauth`

A tiny hook tool for bash/zsh to re-authenticate with ssh-agent.

## What's this?

This small command manages sockets for SSH authentication created by
`ssh-agent`.  It automatically detects the valid socket and rewrite the symlink
to avoid disconnecting.

## When do I need this?

`ssh-agent` stores the path for sockets into `$SSH_AUTH_SOCK`.  It is a path
below, for instance.

```sh
# example in macOS
$ echo $SSH_AUTH_SOCK
/private/tmp/com.apple.launchd.sa197Z7kVN/Listeners
```

This changes every time you login and `ssh-agent` detects the changes and makes
`$SSH_AUTH_SOCK` indicate validly.  Usually that's enough.

But when you use terminal multiplexers -- `tmux`, `screen`, or so --, it breaks
this.

## Reproducible way

1. Login your Mac.
  - `$SSH_AUTH_SOCK` is `/some/path/to/Listeners`.
2. SSH into a Linux box.
  - `$SSH_AUTH_SOCK` will be changed into `/tmp/path/to/agent.foo`.
3. Launch `tmux` in the box.  Also you can use SSH authentication in `tmux` by
   `ssh-agent`.
4. Detach `tmux` and logout the box.
5. SSH into the box again.
  - `$SSH_AUTH_SOCK` will be changed into `/tmp/path/to/agent.bar`.
6. Attach the existent `tmux` session.
7. `$SSH_AUTH_SOCK` will be `/tmp/path/to/ssh-agent.foo`, not `...bar`.
   You can NOT use SSH authentcation in it.

## Solution

Here is `gosshauth`.  You can install this from [release page][] or a command
below.

[release page]: https://github.com/delphinus/gosshauth/releases

```sh
go get github.com/delphinus/gosshauth
```

And you should set the hook for zsh/bash.

```zsh
# for bash
if which gosshauth > /dev/null 2>&1; then
  eval "$(gosshauth hook bash)"
fi

# for zsh
if (( $+commands[gosshauth] )); then
  eval "$(gosshauth hook zsh)"
fi
```

Now you can use SSH authentication even if in the way above.

## How do `gosshauth` work for this?

1. Check `$SSH_AUTH_SOCK`.
2. `gosshauth` checks the existence.
  - If exists, that's all, done.
3. `gosshauth` globs all socket-like files for `/tmp/**/Listeners` and
   `/tmp/ssh*/agent.*`.
4. The candidate that has the latest timestamp is the goal, maybe.  `gosshauth`
   rewrite the `~/.ssh/auth_sock` symlink to target the goal, and set
   `$SSH_AUTH_SOCK` to use it.

## Author

JINNOUCHI Yasushi &lt;me@delphinus.dev&gt;

## License

The MIT License
