+++
date = '2026-01-25T21:02:15-06:00'
draft = true
title = 'microdosing on nix flakes knowledge'
categories = ['nix', 'learn']
+++

I have heard a lot about how nix flakes are so goated for development and how
devs really like it when a repo has its `flake.nix` file but I never really got
the appeal for it. I have been using Nix OS for about 6 months now and I have 
abstained from them since I did not want to add too much to what I already had 
to learn to be able to use and configure Nix OS. But I think that it is about 
time that I start incorporating flakes into my configurations and new repos.

I will start by making a flake to let me develop this site. I think it should be
fairly straight forward since it only requires 3 packages to be able to use it,
git, hugo, and stdenv.cc (for hugo extended). There is no other configuration to
the repo, no other coworkers I have to collaborate with, or huge multi language 
monorepos.

I want what I learn to stick and what better way to make it stick than to write
about my process in learning it.

### prior knowledge

I do have some consequential knowledge about nix flakes due to using NixOS. I
came across flakes when reading about configuring my system and getting
applications on it. At the time of me setting up my first NixOS machine I just 
found them a bit confusing and did not think were that beneficial.

I want to write my thoughts prior to researching that I formed from things I 
have heard online and the brief wiki article I have read about them. 

My beginning impression of flakes is that flakes are, in web dev terms, an 
approach of using the nix package manager as a `package.json` and
`package-lock.json` but for your system. This way it uses the declarative nature
of nix but for an operating system configuration that will always have the 
specific packages pinned to a specific version.

## what even is a flake

The actual [definition](https://wiki.nixos.org/wiki/Flakes#:~:text=A%20flake%20is%20a%20directory%20which%20directly%20contains%20a%20Nix%20file%20called%20flake.nix%2C%20that%20follows%20a%20very%20specific%20structure.) of a flake is super straight forward. It is basically 
just saying how a flake takes form. Which by itself is not bad but it still is
not very descriptive and left me with more questions rather than answers.

Looking a bit more into it you learn that a flake follows the same schema by
convention and with this known [schema](#flake-schema). 

Just learning what a flake is was not very helpful. This was not necessarily as
easy as just learning what a `package.json` file was. For a `package.json` file
you can see pretty much all the [attributes](https://docs.npmjs.com/cli/v11/configuring-npm/package-json) that it creates making it really easy
to know what it does without much understanding. A flake on the other hand is
completely different, it does have fewer root attributes but the fact that one
of them is a function means that by design it can do anything you want it to do.

## flake schema

Now that I had a very shallow understanding of a flake I needed to understand
the structure of it.

Turns out that all flakes can only have 4 top level/root attributes:

- `description`
- `inputs`
- `outputs`
- `nixConfig`

That's easy enough to understand but actually understanding how to write a 
flake.nix file was something else. It requires at least some knowledge of the
nix programming language with functional programming as a plus.

---

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
