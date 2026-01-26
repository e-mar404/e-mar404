+++
date = '2026-01-25T21:02:15-06:00'
draft = true
title = 'Flakes Intro'
+++

# Prior knowledge

I do have some consequential knowledge about nix flakes due to using NixOS so I
came across flakes when reading a out configuring my system and getting
applications on it. At the time of me setting up my first NixOS machine (around
summer 2025) I just found them a bit confusing and too much.

I want to write my thoughts prior to researching that I formed from things I 
have heard online and the brief wiki article I have read about them. 

My beginning impression of flakes is that flakes are, in web dev term, an 
approach of using the nix package manager as a `package.json` and
`package-lock.json` but for your system. This way it uses the declarative nature
of nix but for an operating system configuration that will always have the 
specific packages pinned to a specific version.

Topics that I want to address:

- [ ] high level overview of nix
- [ ] Why use flakes
- [ ] What are they
    - [ ] brief overview of inputs & outputs
- [ ] How can they help me
- [ ] should an app/cli that can be easily installed (ie go install ..., npm
  install ...) have a flake, or should I even bother creating one?
- [ ] Walk through for creating the flake that I use for this blog
    - [ ] how to keep shell env
    - [ ] what is direnv
    - [ ] added nix icon when in nix shell (tangent on why the heck was it so
      hard to find out if there is a good approach at knowing if you are within
      a nix shell)
- [ ] conclude with if I want to transition my nixos config to using flakes or
  not

### Rant

I thought this was supposed to just work, what happened to ending the "it works
on my machine" syndrome.

I expected to be able to run `nix develop` and be able to use the same flake
file with literally 3 pkgs and thats it. Why am I having to debug, is this just
some massive skill issue on my side?

Errors gotten:

This is with flake from `7417fd90`:

```bash
error:
       … while evaluating the attribute 'legacyPackages.x86_64-linux'
         at /nix/store/3p306srz83h9z9v0ma9xcxb8y8cdxkxj-source/lib/attrsets.nix:977:47:
          976|   */
          977|   nameValuePair = name: value: { inherit name value; };
             |                                               ^
          978|

       … while evaluating a branch condition
         at /nix/store/3p306srz83h9z9v0ma9xcxb8y8cdxkxj-source/pkgs/stdenv/booter.nix:107:9:
          106|       thisStage =
          107|         if args.__raw or false then
             |         ^
          108|           args'

       (stack trace truncated; use '--show-trace' to show the full, detailed trace)

       error: error parsing derivation '/nix/store/rr4cpsx1w84h2l1h71ch8spcp221kzkr-m4-1.4.19.tar.bz2.drv': error: exp
ected string 'D'
```

And this is when I changed the nixpkgs to use the unstable channel:

```bash
error: builder for '/nix/store/h414vl12bmgrf03asyslmw4yx0dx242w-nix-shell-env.drv' failed to produce output path for output 'out' at '/nix/store/h414vl12bmgrf03asyslmw4yx0dx242w-nix-shell-env.drv.chroot/root/nix/store/kxwmaaskhx4cb5l3jvc1yh8njdwq865b-nix-shell-env'
```

