+++
date = 2024-09-12T23:39:00-06:00
title = 'Personal Blog with Go'
categories = ['Explanaition']
draft = false
+++


## Storing the blog content 

To start with finding a way of storing the blog content I had to set some
requirement. The only strong requirement is that the source file / the original
blog content should come from a markdown file. That's it. So now with that
chosen there is few options as to how to store the content:

1. A file and each post is organized via file system
2. A db holding a string of markdown file
3. A mix of both

I do not consider storing a file as a blob on a db as a viable options since I
think it is just bad practice to do so. Specially for an application of such
massive (1 user per month, me) scale. The option that makes the most sense
depends on one other thing though, how fast do I want to deploy this feature. 

### Why choose the second best option

I find the top option to be a mix of both, with the db holding meta-data and
the file path (or identifier of the blog post) t be the best option. But since
I just want to start getting this out there I feel that speed of implementation
is another light constraint, and for that reason I have decided to manage the
blog content via the file system, and no db (for now), of the server with each
blog entry being its own markdown file.

However the way this feature has been implemented does not close the door to
the possibility of integrating a db. Each step has been abstracted and
encapsulated into its own separate procedures that do not care about the
implementation, so as long as the next step gets the data it does not matter
how the previous step got it.

### Blog Identification

As of now there are only two identifiable fields for a blog, an `Id` and a
`Name`. The `Id` is the name of the file but without the `.md` extension and
the `Name` is the name of the file.



## Showing blog content

Now that we know how we are going to store each blog we have to worry about how
to show that data. I don't want to just have plain text with # and `[]()` with
no formatting. For this I used a go
[library](https://github.com/gomarkdown/markdown?tab=readme-ov-file) to do the
heavy lifting of the markdown parsing. 

### Showing a specific blog

The /blogs route will give a list of the published blogs. While /blogs/:blog-id
will give the contents of that that blog-id has (if it does not find one it
will show "something went wrong").

When a valid blog is requested then it will get the markdown file, parse it and
return the html for it, this is when it will be displayed on the website.
Pretty simple.

## Scaling

Now, this is just a very simple, crude way of having a "blog". Again I am just
wanting to put it out there. The uploading is bad, the editing is less than
ideal and there is no keeping track of anything. It is just a document viewer
at this point.

To increase functionality I am thinking of incorporating a db. This would
include meta-data like last time modified, tags, blog-id, if it is published or
not, etc. This can be easily added later on as checks.

One big challenge that I do see coming is uploading a new blog. Currently the
website pulls all the files under /blogs from the main branch of the GitHub
repository. Which means that if I want to post or edit something I need to
write a markdown file, with neovim btw, commit to my branch, create a pull
request and then deploy. It would be a lot easier to just have a user portal
and have the ability to edit and publish through the website. This does require
authentication of some sort, temporary storage of the file that is being
published and the finding a way to create a file on the main branch for that
entry since I would want to keep pulling the blogs from the file system.

A user portal for writing might be pushing it too much (it is) but the db
integration for keeping track of meta-data does sound very useful and will be
implemented in the near future. For now manual editing and publishing of
articles will do the trick.

## Final thoughts

Overall this was a fairly simple feature with potential of getting more
challenging, fun, useful and did complete my light constraint of speed since I
was able to get this out in a day. 

The ability to use markdown files and parse those files into html is amazing
since I really like markdown for writing. I started using markdown for note
taking in uni and fell in love with it relatively fast, it has all the features
I need and I can view it anywhere without being tied to a specific company or
product. Maybe an even more universal way of going about note taking might be
with text files, but I feel like the styling of markdown just gives a very much
needed organization to the notes overall.
