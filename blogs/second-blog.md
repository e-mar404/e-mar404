# This is the second blog

## How am i going to support multiple blogs

I am thinking of instead of storing the blog md files in a db to just use the file system since that seems more fit for the scale of 1 user (me) a month

### Naming

The title of the blog will be the file name. E.x.

/blogs/second-blog.md

Will be

Title: Second Blog

Potential of having "collections/folders" of blogs via nested dirs (not important rn only requirement rn is to display a list of blogs and be able to read them)

## Getting the blogs

- Have a function that gets the file names of all the md files under /blogs/* 
- Each file will have and id of its name without the .md extension (naming-convention-is-this)
    - this will be for the list of items on the front end to be able to get the data of an individual file
