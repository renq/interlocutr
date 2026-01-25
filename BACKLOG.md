# The backlog

1. Improve auth mechamism and create something more sophisticated than hardcoded login and password. dd possibility to define users. User should have an ID because I'll need this later.
1. Create add/get sites API for a user.
1. Figure out how to calculate coverage. The current way I wrote tests shows almost 0 where it should be almost 100% (as I use TDD).


Plan:
- POST /api/admin/sites + pyload. Should should create a new site.
- GET /api/admin/sites. Should require valid authentication. User should see only their websites.


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
