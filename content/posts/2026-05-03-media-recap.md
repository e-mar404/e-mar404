+++
date = '2026-05-08T20:42:00-05:00'
draft = true
title = 'Educational Week'
categories = ['weekly-recap']
+++

### Week of May 03 - May 09 

Most of this week was spent researching things and listening to podcasts related
to the process of learning and academic-adjacent topics. I am not going to recap
too much since most of it was practical knowledge and I am not looking at
explaining it just yet. It would be cool to start making videos where I am
explaining concepts I've learned and how I applied them to projects I am working
on. But this will be a project for another day. 

---

## Nginx: Beginner's Guide
[Docs by Nginx](https://nginx.org/en/docs/beginners_guide.html)

Currently working on a project at work in which I am hosting Bookstack, an open
source wiki software[^1]. In order to bring that self hosted instance of
Bookstack I want to install a reverse proxy and since I have not done that
before I am reading through their docs now. In theory I know what a reverse
proxy is and what it does, but I have never had to implement or set one up so
this is going to be a cool little learning opportunity. 

---

## Go by Example: Command-Line Flags
[Post by Mark McGranaghan and Eli Bendersky](https://gobyexample.com/command-line-flags)

This was part of my research on a personal project I am doing. I've written cli
programs in go before, but usually I have needed a bit more control of the
arguments, flags, and commands so I end up using an external library. Usually I
use viper[^2] and cobra[^3]. But on this project, a neovim plugin using a go rpc
server, I am only needing very simple and controlled use of flags so I opted for
the built in library `flag`.

---

## THIS is Android 17
[YouTube video by 9to5Google](https://www.youtube.com/watch?v=C8bZmQ2YXGM)

Long ago, when I was very into jailbreaking my iPhone and installing developer
betas I used to be super into all the new features and into the channels
covering what to expect for the next release of iOS. This interest has dropped a
lot since then, but now that switched to an android the itch to customize and
see what is coming soon is starting to come back. 

From the sound of it this might not be a crazy overhaul, not like the update
from 15 -> 16 on Material 3 Expressive[^4]. Still some pretty interesting
updates, from this video the things that were the most interesting to me where:

- the multitasking floating apps
- per app dark mode controls
- the possibility of having more blurred surfaces (really hope this is not like
  liquid glass[^5])
- native app blocking

---

## Why does Cambridge teach OCaml as the first programming language?
[YouTube video by Frank  Stajano Explains](youtube.com/watch?v=6APBx0WsgeQ )

I love hearing about different colleges and their takes on why they teach what
they do and the way they do that. I did not learn about functional programming 
languages until my last year of university and I remember having quite a hard
time with the entire paradigm of functional programming. To hear that this is
the first language introduced at Cambridge was very shocking. 

The main arguments as to why they use OCaml as the first language were:

1. functional programming gives a solid mathematical foundation
2. using functional languages help with proving programs are correct (ok formal
verification[^6])
3. good challenge for already high achieving scholars while still being beginner
friendly (It is almost everyone's first time using OCaml, so a good equalizer)

All of these are good points, and thinking deeper about it, I was probably just
having a hard time with the functional paradigm because of my bad habits of
using non functional languages. Starting out with a functional language might
not be so bad after all, specially for a theoretical degree. I do keep
forgetting that the Computer Science degree is a very theory heavy degree. Its
roots start with research, not as a pathway to software engineering, which I 
once erroneously though it as. Thinking of the CS degree this way makes so much
more sense to teach a functional language (and probably C since it is the
ultimate 'high level' language that expresses how a computer works at the most
human-friendly way.)

There is also a section about him answering the question of if being good at
Mathematics is required to be a good computer scientist/programmer. This was
interesting but not why I clicked on the vid :P, it was honestly purely because
of OCaml. That does not mean it was not a good section though, it was
interesting and worth a listen to hear about the similarities between
mathematicians and computer scientists.

## Grant Sanderson (@3Blue1Brown): The High Cost of Being a Second-Hand Thinker
[YouTube video by Life of Luba](https://www.youtube.com/watch?v=Rtkac4WHC1o)

> "I actually know the solution for education ... it is neither scalable nor
> ethical" - 3Blue1Brown

This was the highlight of the podcast. Grant Sanderson brought up a humorous non
real solution to the problem of education. This was during a section in the
video that they were talking about how he sees education evolving with LLMs.
Grant brings up the argument that in education there is not a problem of lack of
explanation, it is a problem of motivation. And to solve that you need to hire
actors, that are the same age as the students you are teaching, assign a mark to
each actor and have them get close/flirt with each other into learning about the
topic at hand. The though process behind that is that there is not a single
stronger motivation than social peer motivation.

I remember being in high school and watching his videos, and every time I
finished them I would get this delusional thought that I wanted to get a math
phd and just study/teach the intricacies in numbers and the world around us.
Right before university Math was a very strong contender for my major, but
ironically I ultimately ended up going to Comp Sci for the job security 🤡.

All throughout the interview I got a similar feeling that I used to when I
watched his "essence of" videos[^7] [^8]. The only difference is that now I got
that feeling with education, I felt so inspired and made me want to keep
learning as much as possible. I have been deviating so much from the path of
engaging in what I am curious about, and instead I am just learning things to
better prepare for the job market. I think this has been very detrimental to my
motivation and in turn my learning. Hearing them talk about what success means,
the difference between being a source vs a relay, and the story about his recent
speech at a wedding makes me feel very inspired. This is the type of content that
heals the AI trauma. 

[^1]: https://bookstackapp.com
[^2]: https://github.com/spf13/viper
[^3]: https://github.com/spf13/cobra
[^4]: https://9to5google.com/2025/05/13/android-16-material-3-expressive-redesign/
[^5]: https://www.phonearena.com/news/liquid-glass-was-so-bad-that-apple-will-give-you-another-way-to-tone-it-down_id178795
[^6]: https://en.wikipedia.org/wiki/Formal_verification
[^7]: https://www.youtube.com/playlist?list=PLZHQObOWTQDMsr9K-rj53DwVRMYO3t5Yr
[^8]: https://www.youtube.com/watch?v=fNk_zzaMoSs&list=PLZHQObOWTQDPD3MizzM2xVFitgF8hE_ab
