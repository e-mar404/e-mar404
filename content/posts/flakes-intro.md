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
- [ ] conclude with if I want to transition my nixos config to using flakes or
  not

