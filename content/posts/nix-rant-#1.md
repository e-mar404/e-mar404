+++
date = '2026-01-27T11:13:42-06:00'
draft = true
title = 'nix rant #1'
+++

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
Turns out that my nix store was corrupted ... this fixed my issues:

```bash
sudo nix-store --verify --check-contents --repair
```

My battery is very bad so it might have died during a `sudo nixos-rebuild switch`
or something.

But shout out to the goated comment in a similar open gh [issue](https://github.com/NixOS/nix/issues/11148) by
[@n8henrie](https://github.com/n8henrie).

Ultimately this is what fixed my issues:

```bash
nix-store --query --referrers-closure \
    $(find /nix/store -maxdepth 1 -type f -name '*.drv' -size 0) |
    xargs sudo nix-store --delete --ignore-liveness
sudo nix-store --gc
sudo nix-store --verify --check-contents --repair
sudo nixos-rebuild switch
```

I finally rebooted for good measure, should've tried to at least get a package
with `nix-shell` to see if that had worked. But I assumed that since the nix
store repair did not output any errors I was confident enough that it had
worked.
