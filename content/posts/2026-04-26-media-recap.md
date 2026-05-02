+++
date = '2026-05-02'
title = 'GitHub On The Spotlight'
categories = ['weekly-recap']
tags = ['github', 'neovim']
+++

### Week of [Apr 26 - May 02]

Recently I saw a LinkedIn post by
[@noelcodes](https://www.linkedin.com/in/noelcodes/) where he mentioned that he
gathers a bunch of articles on Saturday and then reads them all on Sunday.
Seeing him share all his thoughts on the articles he read gave me the motivation
to start writing weekly recaps on the articles and videos I watch. 

I find myself struggling to recall significant points from media a few days
after  consuming it. This way I am able to read and watch with the purpose of
writing  about it later on and actively think a bit more about what is being
said.

I'm not going to recap the weekly reads that he posted since that would just be
a recap about a recap, but I encourage you to check out his
[site](https://www.noelcodes.dev/). Now, let's dig into the media that stood out
to me the most this week ...

## Should Neovim support transitive plugin dependencies?  
[Blog by Justin M. Keys](https://sink.io/jmk/neovim-plugin-deps/)

I came across this blog post while transferring my lazynvim config to just  use
`vim.pack`. 

There were 3 key points that I took away from this read:

1. Dependencies *WILL* introduce risks, but there will be different risks
depending on how you connect the source to those dependencies. If you vendor in
dependencies you risk zero days. If you use a dependency manager like a package
manager, you risk supply chain attacks

2. There are ways to mitigate the risks that dependencies bring. One such way is
by including a great and well rounded standard library, examples given were
golang and the goal of stdlib in neovim. Another is to make dependencies "hard"
to manage, either artificially or by design. This way there will be more thought
put into having dependencies

3. This was not an argument in the blog, but I am starting to believe that we
should roll our own *x* a lot more than we do now. Especially if it is the core
of what makes your product, project, or company stand out. I am introducing this
mindset into my own work by rewriting my neovim config to use as few plugins and
3rd party configurations as possible.  I took the first step by not using a
plugin manager

---

This next section will be all about ✨GitHub✨. Mainly due to how much it came
up in my feed, and the cultural significance that all the issues they are
causing have on the developer community.

---

## Ghostty Is Leaving GitHub  
[Blog by Mitchell Hashimoto](https://mitchellh.com/writing/ghostty-leaving-github)

This was a short read on how Mitchell feels about GitHub and how he is taking
Ghostty out of GitHub. He doesn't mention where the project will go, but if I
were to guess I would say Codeberg, similar to the  Zig project[^1].

Quick takeaways:

- I did not know that Mitchell was so nerdy. It makes me feel like nerdiness is
  making a massive comeback
- He is not the only one that I saw this week complaining and mentioning how
  disappointed they were in GitHub. In most of the posts I've read, there is
  this underlying tone that they still want GitHub to succeed. This is the place
  where Hashimoto and so many others found the wonders of open source, but even
  their day 1 supporters are turning their backs towards GitHub. 
- Him mentioning that his personal projects are going to stay on GitHub makes me
  think of this weird alternative timeline future. A future where GitHub is the
  coding playground where you point all your AI projects and automated AI bots
  to save their work. And we use some other service for real work that  has a
  clear vision and direction which takes performance and uptime seriously[^2]

Since this was a very heavy GitHub week I'll put here other resources  that I
found interesting on the topic. [^3] [^4]

---

## Open Source Friday with Mitchell Hashimoto  
[Interview hosted by GitHub](https://www.youtube.com/watch?v=0npMvuh7qNc)

It is kind of ironic that I found this the same week as his article mentioned
above. Throughout the interview, which was only 2 months ago, you can tell there
are no bad feelings from Hashimoto toward GitHub. It seems to be well understood
that it is not the fault of the employees. It is mostly on the company as its
own  entity. But who can we blame? Probably the CEO should take responsibility,
oh wait ... there is no CEO. Well,  anyways that was not the point of this video
and there were some really good  takeaways.

A huge part of the talk, if not all of it, was about how a project as big and as
popular as Ghostty[^5] is dealing with the new age of AI contributions. He
brought up an interesting point on how AI is little by little getting rid of his
humanity. There were a few specific situations where he was harsher than he
would have been prior to AI and closed down pull requests that seemed purely AI
generated. It turned out to be new developers who were using AI to help with the
PR and genuinely wanted to help. 

This could 100% be the new harsh and unwelcoming environment for new developers
that early Stack Overflow used to be. Instead of closing down questions with
"duplicate", "off topic", or the well known "RTFM", PRs will now be closed  with
"AI slop". 

I am a believer that a slight barrier to entry hugely increases the quality of
the help, contributions, AND learning. I think Hashimoto's vouch [^6] project
that he mentions in the YouTube video is bringing back the old barrier to entry
that a pull request used to bring in a good way.


[^1]: https://codeberg.org/ziglang/zig
[^2]: https://mrshu.github.io/github-statuses/
[^3]: https://www.youtube.com/watch?v=R7ex-Gt8dtw&t=1828s
[^4]: https://www.bookstackapp.com/blog/project-migrated-to-codeberg/
[^5]: https://ghostty.org/
[^6]: https://github.com/mitchellh/vouch
