# The backlog

1. Install golang-migrate and sqlx and create repository implementations using SQLite. Try to use exactly the same tests cases for all repository implementations.
1. Improve auth mechamism and create something more sophisticated than hardcoded login and password. dd possibility to define users. User should have an ID because I'll need this later.
1. Create two entrypoints for use cases: web and command line
1. Rethink directory structure.
1. Create JS code to render comments. Maybe we can skip JS and use something like https://github.com/donseba/go-htmx?


## High priority

1. Create sites module to keep data about sites for which we keep the comments.
1. Create another implementation of comments storage (in files or sqlite)

## Less important ideas

1. Go back to https://github.com/krzysztofreczek/go-structurizr and try to make better diagrams
1. Create dockerfiles, makefile, docker compose
1. Create github actions pipeline
1. Verify that referrer header matches to site settings
1. Create comment confirmation email
1. Create comment confirmation API
