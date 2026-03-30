# The backlog

1. Add backend validation. Author and text cannot be empty!
1. Verify headers and allow adding comments only when referrer matches!
1. Stop using json in app layer.
1. Add logs (for echo errors)
1. CQRS. What about creating a separate domain and query structures and removing getXXX from repositories?
1. Create moderation tool. Start from remove comment feature. We can try using LLM for spam classification. CLI or Web with something like https://github.com/donseba/go-htmx?
1. Improve auth mechamism and create something more sophisticated than hardcoded login and password. dd possibility to define users. User should have an ID because I'll need this later.
1. Create two entrypoints for use cases: web and command line
1. Rethink directory structure.
1. Use ROLLBACK transaction in tests instead of DELETE FROM... [not easy with sqlx... deprioritized]


## Less important ideas

1. Go back to https://github.com/krzysztofreczek/go-structurizr and try to make better diagrams
1. Verify that referrer header matches to site settings
1. Create comment confirmation email
1. Create comment confirmation API
1. Handle spam and moderation.
