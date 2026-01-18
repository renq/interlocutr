# The backlog

1. Install cobra and create a simple command (help)
1. Create command line command to generate a token
1. Install JWT and check the token in the API
1. POST /api/admin/sites returns 401 Unautorized if user token is not present in the request.


Plan:
- POST /api/admin/sites + pyload. Should require valid authentication.
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
