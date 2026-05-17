+++
date = '2026-05-17'
title = 'media recap #3'
categories = ['weekly-recap']
+++

### Week of May 10 - May 16

It is surprisingly hard to consistently think of a common theme that is not
"programming" to call the weekly blog post, sooooo I'll just call it Media Recap
\#number unless there is a clear unique theme. My key takeaway from this week
is very profound: I love YouTube :). Ok maybe not profound in the slightest, but
I truly enjoy all the information that is out in the wild, hidden behind a
thumbnail. There are so many talented people with so much good knowledge from
recorded lectures, TED talks, interviews, and much more. Video/audio might be my
favorite way to consume information, and it shows with my shameful screen time on
YouTube...

There is a lot to cover this week and I am not that great at writing, so the
format will change a little bit. I will make my commentary into quick-fire
comments as bullet points.

---

## Can we test it? Yes, we can! - Mitchell Hashimoto
[Video recording of BugBash session @ Antithesis](https://youtu.be/MqC3tudPH6w?si=WpcjNBFU1IeXIs2M)

- Really useful information about how to test, and language-agnostic too, which
  is really nice

- He brings up a great point that usually if there are no tests in a codebase,
  the primary reason is because "it cannot be tested yet/easily." And that this
  is mostly a skill issue. If it is not easily testable or is impossible to test,
  then it is because of poor structure or because of a lack of testing knowledge
  (not everything is a unit test)

- This talk made me realize that I am thinking about tests as if they are simple
  functions that just test the output of the function, and oh boy am I so wrong
  about that

- I really hadn't ever thought about testing code that runs on the GPU, and
  apparently not a lot of people either ?? since Mitchell said there weren't
  many good resources on this

- Testing strategies covered: snapshot testing, isolating side effects, GPU
  testing, VM testing.

- Both Ghostty[^1] and Charm[^2] have a very interesting challenge. Testing things
  on the terminal feels a bit more involved than what I have experienced with
  backend/frontend related work, and all of the testing strategies that he
  mentions in the talk work really well with this type of very state-heavy
  workflow.

- Also I too believe that Nix[^3] for VM testing, or just VM work, will be the
  future if the language learning curve gets less steep (or at least close to
  what the future will look like, sorry Docker and other declarative tools)

## Designing Data-Intensive Applications with Martin Kleppmann
[Interview by The Pragmatic Engineer](https://youtu.be/SVOrURyOu_U?si=IMCPgAopCByh_ajE)

- Not sure if it is due to the times or due to me being a little further into my
  journey with programming, but it feels more important than ever to know how to
  design software, and this talk really reinforces that thought

- Academia does have the dream job for people that love programming for the love
  of it and don't really care about the business result. Academia has the
  benefit of just thinking purely about the future and implementing new
  technologies that *might* have a useful application later

- I do not know what's up, but lately I have heard formal verification[^4] more
  in the last month than in the 4 years of my university undergrad. There is
  something there that I think would be very useful, but I do not yet have the
  understanding of it to really have a strong opinion

- Distributed systems have to work so closely with DevOps and infra for everything
  to work properly. Also, the focus on distributed systems makes me miss my short
  time with Erlang and Elixir

---

## Why Programmers Adopt Bad Ideas - Casey Muratori
[Interview by NeetCode](https://youtu.be/Z8B4BSi35CI?si=HjEPkYv2abO4MIaY)

- Casey is just like me. I used to not like math and was so happy to stop doing
  it in uni, but not too long after starting uni, math grew on me a lot

- A huge chunk of the interview is Casey talking about what he thinks is the
  most interesting math/programming-related problem[^5] he's encountered in his
  years, and I love the way he explains things. Sometimes it feels a little
  slow, but there is so much passion behind his words that makes me want to go
  out and learn more about what he is talking about

- Very interesting to see Casey and NeetCode talk about their pushback to using
  AI

- Also the title is very true, programmers adopt bad ideas because we are very
  opinionated individuals who are able to make those ideas work and become a
  reality (at least a reality within our heads and immediate environment, maybe
  not reality in the global sense)

- Real, the state of current webdev sucks, srry :P

---

## The Wildlife Sanctuary You Can Visit from Anywhere | Maya Higa | TED
[Video recording of a TED talk](https://youtu.be/-YwzCOVeLEE?si=XNe_tFjy-aj-yiw4)

- So cool what the internet allows people to do

- Really like the business model that Alveus Sanctuary[^6] started for wildlife
  conservation, especially since the pushback that a lot of people have on zoos
  and other animal conservation places is that they are overall negative and not
  a net good. This can be due to a plethora of reasons like excessive
  exposure to human interaction, budgets needing to be divided between animals
  and humans (parking, concessions, etc.), etc. At some point all of this becomes
  an incentive to give the animals a worse life[^7] [^8]

- Excellent example of how everyone contributing a little bit adds up to a
  meaningful impact 

---

## 10 Online Philosophy Courses I'd Take If I Had Any F*ing Free Time
[Substack article by Meaghan Green](https://substack.com/home/post/p-196725955)

- **I LOVE FREE LEARNING RESOURCES <3**

- The summary of some courses really reminds me of math, specifically the ones
  that deal with logic. Makes me wonder how much overlap there is between math
  proofs and formal philosophy logic 

- I will for sure put the Death Yale course[^9] on my list of things to
  look at later. Anything to do with the morality and ethics of death is very
  interesting to me. I am not well versed in enough philosophy background to
  where I could go into a debate or anything like that, but that does not take
  away the interest and curiosity from the topic

- Same for the logic course. The thumbnail[^10] reminded me of discrete
  mathematics. Could be that it is just the format of a white paper with some
  symbols, but interesting nonetheless. This would be the ultimate application
  of Haskell, just imagine writing a formal philosophy logic paper in Haskell
  hehe 

---

## The Missing README (Ch. 1)
[Book by Chris Riccomini & Dmitriy Ryaboy](https://themissingreadme.com/)

- Found it through the [Pragmatic Engineer
  interview](#designing-data-intensive-applications-with-martin-kleppmann).
  There was a section where he mentioned one of the authors (Chris Riccomini)
  and how they've been coworkers in the past and invited him to co-author the
  second edition of Designing Data-Intensive Applications

- So far it seems to be setting the roadmap of a very vague, but goal-oriented,
  career of a software engineer 

- Excited to read more about this, even if I don't have a software engineering
  job yet...

---

That was quite a bit this week. I should probably stop consuming so much and
just start doing stuff.

[^1]: https://ghostty.org/
[^2]: https://charm.land
[^3]: https://nixos.org/
[^4]: https://en.wikipedia.org/wiki/Formal_verification
[^5]: https://paulbourke.net/geometry/implicitsurf/
[^6]: https://www.alveussanctuary.org/
[^7]: https://tinyurl.com/4nffc25k
[^8]: https://www.idausa.org/experts-agree-zoos-harm-good/
[^9]: https://meaghangreen.substack.com/i/196725955/9-death-phil-176-yale-open-courses
[^10]: https://meaghangreen.substack.com/i/196725955/7-logic-i-24241-mit-ocw
