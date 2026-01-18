+++
date = 2024-10-17T10:32:00-06:00
title = 'Code output in neovim'
categories = ['Showcase']
draft = false
+++

While taking software design and programming languages courses with Dr. Venkat
Subramaniam, I was inspired by how he used TextMate’s tooltip feature during
lectures. He skillfully displayed the output of code snippets directly in his
editor without needing to switch to a terminal. This seamless experience caught
my attention as I was just beginning to explore text editors like Neovim
instead of traditional IDEs. However, upon learning that this tooltip feature
was only found on TextMate, I felt discouraged and decided to focus on
mastering Neovim instead.

That was back in the Spring 2024 semester. But come Fall 2024, my interest in
creating custom Neovim plugins resurfaced, and I decided to build my own
version of the tooltip functionality for Neovim. Here’s a look at my journey
and the plugin I developed.

You can find the repo for my plug in with demo, installation guide and more
[here](https://github.com/e-mar404/tooltip.nvim).

## Building the Neovim Plugin

I’ve been using Neovim as my primary editor for around six months, but when I
started, I barely knew how to install plugins without extensive googling—let
alone create my own. YouTube was a huge help, especially tutorials from
@teej_dv. These videos taught me the basics of writing Lua plugins and made the
Neovim documentation much more approachable. With the help of telescope.nvim, I
could quickly look up help tags and dig deeper into the docs.

After experimenting with opening windows and displaying buffers, I was ready to
tackle my own solution.

## How the Plugin Works

The initial plan for my plugin was simple:

1.	get the file path
2.	execute the file using a command based on its file extension
3.	display the output in a floating window

I wasn’t sure of the exact implementation at first, but I figured I could work
out the details as I went. Surprisingly, the final version remained pretty
close to my initial concept, showing that a bit of upfront planning went a long
way.

### Managing Commands

To handle the different commands needed for various file types, I created a
default file pattern table. This table maps file extensions to commands, using
%s as a placeholder for the file path.

One challenge I hadn’t anticipated was allowing users to customize these
commands. Initially, my plugin only supported the default
file-extension-command pairs I set up. But I realized users might want to tweak
these or add their own. To address this, I set up a user-defined file pattern
table that merges with the default patterns, allowing users to specify commands
through `tooltip.setup()` without losing the default options. If a user removes
a custom command that belongs to one of my file extensions already set up on my
default list, then it automatically falls back to the default.

There are two ways to add custom commands:

1.	temporarily add commands for the current Neovim session (they’ll disappear
    once Neovim is closed)
2.	add persistent commands through the setup function

I considered adding a third option: allowing users to add persistent commands
directly through a user-command. However, this would require storing those
commands somewhere (like a text file) and managing their removal. It seemed
like too much complexity for a minor convenience that could be achieved through
the setup function.

## Unexpected Challenges and Solutions

A surprising hurdle was handling asynchronous output. I couldn’t open a buffer
that was actively being written to, so I had to wait for the command to finish
executing before writing the output to the buffer and displaying it in a
floating window.

Another unexpected challenge involved handling terminal colors. Initially, I
was focused on simply printing to standard output, so I didn’t consider color
codes. However, after encountering colorful compilation error messages in
Scala, I realized the escape sequences made the output hard to read. Instead of
stripping the color codes, I opted to display the output in a terminal window,
preserving the colors. Now, I appreciate how much easier it is to parse error
messages with proper highlighting.

Documenting the plugin was another interesting aspect. It was fascinating to
see how formatting a doc file correctly in the right directory made it
immediately accessible through telescope and other help tag searches in Neovim.
This process really highlighted how much thought went into Neovim’s
documentation standards, and it inspired me to aim for that level of polish in
my own work.

## Biggest Lesson Learned: Read the Docs

If there’s one lesson I learned through this process, it’s the value of reading
the documentation. While online tutorials and community forums can be helpful,
nothing beats the depth and detail of the official Neovim docs. Being able to
quickly reference these using telescope.nvim made my development process
significantly smoother.

I’ve been using this plugin frequently to quickly view program outputs during
code testing, and it’s been a reliable tool. While there’s still the wait time
associated with compiling languages like Scala, it’s a time-saver for many
tasks and fits seamlessly into my workflow.
