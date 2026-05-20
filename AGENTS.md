# Agent Guide for PBD Blog

This is a Hugo-based static site for "Publicly Bad Dev" (PBD) - a personal developer blog with a terminal/file-tree aesthetic.

## Project Type & Stack

- **Static Site Generator**: Hugo (v0.154.4+)
- **Theme**: Custom (no external theme used, despite gitmodules reference)
- **Styling**: Plain CSS (no SASS/SCSS compilation needed locally)
- **Deployment**: GitHub Pages via GitHub Actions
- **Dev Environment**: Nix flake available (optional)

## Essential Commands

| Command | Purpose |
|---------|---------|
| `hugo server -D` | Start development server with drafts enabled |
| `hugo server` | Start production-like dev server (no drafts) |
| `hugo --gc --minify` | Build for production with garbage collection |
| `hugo new content posts/my-post.md` | Create new post from archetype |

### CI Build Command
The GitHub Actions workflow uses:
```bash
hugo --buildFuture --gc --minify --cacheDir "${{ runner.temp }}/hugo_cache"
```

## Project Structure

```
.
├── archetypes/default.md       # Template for new content files
├── assets/
│   ├── css/                    # Stylesheets (concatenated in head.html)
│   ├── images/                 # Site images (logo, favicon, title card)
│   ├── js/index.js             # Mobile nav toggle logic
│   └── files/                  # Downloadable assets (resume.pdf)
├── content/
│   ├── _index.md               # Homepage content
│   ├── about-me.md             # About page
│   └── posts/                  # Blog posts (with categories frontmatter)
├── layouts/
│   ├── baseof.html             # Base HTML template
│   ├── home.html               # Homepage layout
│   ├── single.html             # Single post layout
│   ├── list.html               # List pages (posts/, categories/)
│   ├── _markup/                # Render hooks for headings/links
│   └── _partials/              # Reusable template components
├── static/images/              # Static images referenced in posts
└── hugo.toml                   # Hugo configuration (minimal)
```

## Architecture & Design Patterns

### Layout System
- **Base template** (`baseof.html`): Defines root structure with flexbox layout
- **Main content area**: Split into navbar (file tree) + content pane
- **Two layout modes**:
  - `home.html`: Content view for homepage
  - `single.html`: Post view with prev/next navigation
  - `list.html`: File-picker style listing for posts/categories

### Styling Architecture
- **CSS Variables** in `main.css`: Catppuccin Mocha color palette
  - `--base-color: #1e1e2e` (background)
  - `--text-color: #cdd6f4` (text)
  - `--blue-color: #89b4fa`, `--peach-color: #fab387`, etc.
- **File organization**: Separate CSS files by component:
  - `main.css`: Base styles, layout, headings
  - `file-tree.css`: Navigation/file tree styles
  - `file-picker.css`: List page styles, pagination
  - `code.css`: Inline code styling
  - `mobile.css`: Responsive breakpoint (768px)
- **Bundling**: All CSS files concatenated via `resources.Match "css/**.css" | resources.Concat`

### Content Frontmatter
Posts use TOML frontmatter:
```toml
+++
date = 2025-03-21T00:00:00-06:00
title = 'Post Title'
categories = ['Category1', 'Category2']
draft = false
+++
```

### Navigation System
The navbar (`_partials/navbar.html`) presents a **file tree metaphor**:
- `~` links to home
- `posts/`, `categories/` are directories
- Categories dynamically generated from taxonomies
- `about-me.md`, `resume.pdf` are files
- Email and GitHub shown as symlinks (`->`)
- Active state highlights current location

### Link Rendering
- All external links (`http/https`) automatically get `rel="external" target="_blank"`
- Heading anchors include hash prefixes (`## Section` renders as `## Section` with link)
- Title links in single posts prefixed with `#` to mimic markdown headers

## Content Patterns

### Creating a New Post
```bash
hugo new content posts/my-new-post.md
```
- Edit the frontmatter: set `draft = false` to publish
- Add `categories = ['category-name']` for taxonomy grouping
- Categories automatically create list pages at `/categories/{name}/`

### Categories as Taxonomy
- Categories are Hugo's default taxonomy (no config needed)
- Auto-generated pages at `/categories/{category-name}/`
- Access in templates via `site.Taxonomies.categories.Alphabetical`
- Categories use `file-tree` styling in sidebar

### Post Navigation
Single post pages include automatic prev/next navigation when in a section:
- Sorted by `.ByDate`
- "Previous" = chronologically earlier
- "Next" = chronologically later
- "./.." returns to posts list

## Gotchas & Non-Obvious Patterns

### Draft Handling
- `draft = true` in frontmatter excludes from build (unless `-D` flag used)
- Drafts ARE included in dev server by default (`hugo server` includes drafts)

### Image Handling
- Site images (logo, favicon) go in `assets/images/` (pipelined)
- Post content images go in `static/images/` (served as-is)
- Reference with standard markdown: `![alt](/images/filename.png)`

### Date Formatting
List template uses Go's date format string:
```go
{{ $format := "2006 01 02" }}
{{ .Date | time.Format $format }}
```
Result: "2026 05 20"

### Responsive Behavior
- Mobile breakpoint at 768px
- File tree collapses to `<details>` element on small screens
- JavaScript auto-closes nav on mobile, opens on desktop

### CSS Specificity Issue
- `.file-tree-link` sets `color: var(--text-color)` directly
- Active state uses `.file-tree-item.active` with `background-color` only
- Hover states use `.file-tree-item:hover > *` to target children specifically

### Header Partial is Empty
The `header.html` partial exists but contains only commented-out logo code. The actual logo moved to the navbar (`_partials/navbar.html`). This is intentional per the redesign TODOs in README.

## Dependencies for Development

If not using Nix, you need:
- **Hugo Extended** v0.154.4+ (required for asset pipeline)
- **Go** 1.25.5+ (Hugo dependency)
- **Dart Sass** 1.97.2+ (only if using SCSS - currently plain CSS)

Nix flake handles all dependencies automatically.

## Deployment

- **Trigger**: Push to `main` branch
- **Workflow**: `.github/workflows/hugo.yaml`
- **Target**: GitHub Pages
- **Cache**: Hugo cache preserved between builds at `${{ runner.temp }}/hugo_cache`
- **Build flags**: `--buildFuture --gc --minify`

## Common Tasks

### Update Site Colors
Edit CSS variables in `assets/css/main.css` (lines 1-10).

### Add New Navigation Item
Edit `layouts/_partials/navbar.html`:
- Add new `<a>` with class `file-tree-link`
- Use `.file-tree-item` class on `<li>`
- Add conditional `active` class based on `.LinkTitle` or `.RelPermalink`

### Add Social Link
Add in `layouts/_partials/navbar.html` in the file tree list. Use `file-tree-symlink` class for styling.

### New Category
Just add to a post's `categories` frontmatter. Hugo auto-generates the page.

### Modify Post Archetype
Edit `archetypes/default.md` - templates for `hugo new` command.
