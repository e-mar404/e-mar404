+++
date = '2026-01-27T11:13:42-06:00'
draft = false
title = 'debug session #1'
categories = ['nix', 'debug-session'] 
+++

If I really want to put out what I do I think I should not only be including my
good project showcases but also what I learn along the way. So I will be adding
a tag for [debug sessions](/categories/debug-session).

For today we have an issue I ran into while working on writing an introduction
to flakes. I have both a desktop and a laptop running NixOS and I tend to do
work on both and silly me thought nix and flakes were supposed to just work, 
what happened to ending the "it works on my machine" syndrome, both machines are
literally mine and got imaged with the same nix os iso, why did I run into
issues?!?

I expected to be able to run `nix develop` and be able to use the same flake
file with literally 3 pkgs on both my laptop and desktop. Why am I having to 
debug, is this just some massive skill issue on my side (edit: kinda)? 

For context I wrote a simple
[flake](https://github.com/e-mar404/e-mar404/commit/fc6e5b0c64da04d7560274004636a88ea69fb687) for this site on my desktop and it seemed to
work with no issues. I go to my laptop to pull the changes to the repo and I get
the following error:

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

I thought it was related to the channel I was using for nixpks so I changed it
to the unstable channel, ran `nix develop` one more time and then I got this:

```bash
error: builder for '/nix/store/h414vl12bmgrf03asyslmw4yx0dx242w-nix-shell-env.drv' failed to produce output path for output 'out' at '/nix/store/h414vl12bmgrf03asyslmw4yx0dx242w-nix-shell-env.drv.chroot/root/nix/store/kxwmaaskhx4cb5l3jvc1yh8njdwq865b-nix-shell-env'
```

![yay-new-error](/images/yay-new-error.webp)

After a lot of research I found that a possible cause of this is a corrupted nix 
store. I read a bit more about the nix store and the `nix-store` cmd and it 
turns out there is a way to repair it, great! ... Wait why did this `sudo 
nix-store --verify --check-contents --repair` not work. Why does it say that 
the hashes not match and that it cant repair it. 

Ok fine back to research. This is when I found this [github issue](https://github.com/NixOS/nix/issues/11148). 
I went through it and concluded that my nix store most likely got corrupted 
during a nix os rebuild and then my laptop "died" midway through since the 
battery is pretty bad and it will turn off even if it says 30% battery 
remaining. I don't remember that happening but I also do not use it all the time
so it might've been a while ago.

Anyways after reading through the issue comments I found a comment by [@n8henrie](https://github.com/n8henrie).
And that did the trick. I did have to change it a bit since my config was a bit
weird and it would not take the `nix store` command so I just used the old 
`nix-store` command.

Ultimately this is what fixed my issues:

```bash
nix-store --query --referrers-closure \
    $(find /nix/store -maxdepth 1 -type f -name '*.drv' -size 0) |
    xargs sudo nix-store --delete --ignore-liveness
sudo nix-store --gc
sudo nix-store --verify --check-contents --repair
sudo nixos-rebuild switch
```

Finally, I rebooted for good measure, should've tried to at least get a package
with `nix-shell` to see if that had worked. But I assumed that since the nix
store repair did not output any errors I was confident enough that it had
worked.

End of debug session #1.
