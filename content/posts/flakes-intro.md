+++
date = '2026-01-25T21:02:15-06:00'
draft = false 
title = 'microdosing on nix flakes knowledge'
categories = ['nix', 'learning']
+++

![nix-lambda-linux](/images/nix-lambda-linux.webp)

I have heard a lot about how nix flakes are so goated for development and how
devs really like it when a repo has a `flake.nix` file but I never really got
the appeal for it. I have been using Nix OS for about 6 months now and I have
abstained from them since I did not want to add too much to what I already had
to learn to be able to use and configure Nix OS. But I think that it is about
time that I start incorporating flakes into my configurations and new repos.

I will start by making a flake to let me develop this site. I think it should be
fairly straight forward since it only requires ~3 packages to be able to use it,
go, hugo, and stdenv.cc (for hugo extended). There is no other configuration to
the repo, no other coworkers I have to collaborate with, or huge multi language
monorepos.

I want what I learn to stick and what better way to make it stick than to write
about my process in learning it.

### prior knowledge

My beginning impression of flakes is that flakes are, in web dev terms, an
approach of using the nix package manager as a `package.json` and
`package-lock.json` but for your system. This way it uses the declarative nature
of nix but for an operating system configuration that will always have the
specific packages pinned to a specific version.

## what even is a flake

The actual
[definition](https://wiki.nixos.org/wiki/Flakes#:~:text=A%20flake%20is%20a%20directory%20which%20directly%20contains%20a%20Nix%20file%20called%20flake.nix%2C%20that%20follows%20a%20very%20specific%20structure.)
of a flake is super straight forward. It is basically just saying how a flake
takes form. Which by itself is not bad but it still is not very descriptive and
left me with more questions rather than answers.

Looking a bit more into it I found out that a flake simply follows a
conventional known [schema](#flake-schema). 

Just learning what a flake is and its schema was not very helpful in being able
to write my own nix flake right away I still had to research a bit more. This
was not necessarily as easy as just learning what a `package.json` file was as I
originally compared it with. For a `package.json` file you can see pretty much
all the
[attributes](https://docs.npmjs.com/cli/v11/configuring-npm/package-json) that
it creates making it really easy to know what it does without much
understanding. A flake on the other hand is completely different, it does have
fewer root attributes but the fact that one of them is a function means that by
design it can do anything you want it to do and it quite hard to grasp at first.

## flake schema

Now that I had a very shallow understanding of a flake I needed to understand
the structure of it.

Turns out that all flakes can only have 4 top level/root attributes:

- `description`
- `inputs`
- `outputs`
- `nixConfig`

That's easy enough to understand but now it is time to learn the bare minimum
fundamentals of the [nix language](https://nix.dev/tutorials/nix-language.html).
At some point I should come back and actually learn since I kind of just used
the language basics to see what a function even looks like and very basic things
that are needed for a flake.

## so how did my flake turn out

For now my `flake.nix` file looks like this:

```nix 
{
  description = "All things needed for development on this site";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };
  
  outputs = { self, nixpkgs }: 

    let 
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
   in
    {
      devShells = nixpkgs.lib.genAttrs systems (system: 
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in {
          default = pkgs.mkShell {
            packages = with pkgs; [ go hugo stdenv.cc ];
          };
        }
      );
    };
}
```

If you check line :24 that's where I define packages needed to develop on the
repo that hosts this static site. Very few packages and I am honestly pretty
confident that I'm able to remove `go` and `stdenv.cc` but I left them there
for good measure since technically hugo needs go and a c compiler to install
hugo extended but they might actually already be packaged with the hugo flake.

This will go through all of the systems defined and set what packages the
shell needs, nothing fancy, which is why I wanted to start learning about flakes
by creating one for development of this site.

Note: I did run into some obstacles when trying to use this flake in another
machine but that was a skill issue and nothing inherently wrong with the flake
itself, which you can read about [here](/posts/debug-session-1).

There is a lot more that I still have to learn about flakes but for now I will
leave it at this. For a future time I will want to go deeper on remote flakes,
writing flakes to install applications, how to install app images through
flakes, and possibly trying to dissect a production flake to see if I can infer
the choices made and explain everything that is going on on the flake.

