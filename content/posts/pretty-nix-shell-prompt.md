+++
date = '2026-01-27T16:36:38-06:00'
draft = true
title = 'make my nix shell prompt pretty'
categories = ['nix', 'showcase']
+++

Now that I can just run `nix develop` I can start being a bit more picky. The
default nix shell prompt is low key ugly.

![nix-shell](/images/nix-shell.png)

I want it to follow my prompt styling and also the shell I use, since it
defaults to bash. Now this is not something that everyone wants, if anything I
would think that a nix user would not want this because it adds impurity and
external code that can interact with your environment (i.e. my zshrc). But this
is just for me so I cant do whatever I want. I ultimately found that I can just
pass a command flag to `nix develop` to run my current shell. Like so `nix
develop -c "$SHELL"`.

Now that I can use my shell I have no quick way of knowing if I am on a nix
shell or not. And this is where I start to modify my powerlevel10k prompt. It
was super easy and quick and it made it so much better.

- [ ] add .p10k function

- [ ] add picture of new shell prompt
