# www.jacksorrell.com
[![js-happiness-style](https://img.shields.io/badge/code%20style-happiness-brightgreen.svg)](https://github.com/JedWatson/happiness)

My personal website at <https://www.jacksorrell.com>

## Setup
1.  Clone repository or download branch.

1.  Ensure npm is v5+ and node is v8+.

1.  `npm install`

1.  Create a `.env` file with desired [configuration](#configuration)
(or use environment variables).

1.  Build with `gulp`.

1.  Start with `gulp run`.

## Configuration
All configuration is taken from environment variables.
A `.env` file can be used in root.

### _Options_
`NODE_ENV`: The server mode. Either `production` or `development`. Required.

`PORT`: The port for the server to listen on. Defaults to `3000`.

`GRECAPTCHA_SITE_KEY` and `GRECAPTCHA_SECRET_KEY`:
The site and secret keys for Google's ReCaptcha. Required.

`MAILGUN_API_KEY`: The api key for Mailgun. Required.

`MAILGUN_DOMAIN`: The domain for mailgun to send to. Required.

`MAILGUN_TO_ADDRESS`: The address to send contact messages to. Required.

`MAILGUN_SUBJECT`: The subject line of the contact messages. Required.

`LOG_LEVEL`: The level of logging.
Can be `error`, `warn`, `info`, `verbose`, `debug`, or `silly`.
Defaults to `warn`.

## Systemd
### _Example Configuration_
```systemd
[Unit]
Description=www.jacksorrell.com server
After=network.target

[Service]
User=example_user
Group=example_group
WorkingDirectory=/var/www/www.jacksorrell.com
ExecStart= start
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

## Design Choices
-   Original favicon created using font [The Wastes of Space](http://www.fontspace.com/chequered-ink/the-wastes-of-space).

-   Favicons from <https://realfavicongenerator.net>.

-   Use contact form to protect email from spambots.

-   Contact form sends email via Mailgun and is protected from bots
by Google's invisible recaptcha.

-   Both `www.jacksorrell.com/resume` and
`www.jacksorrell.com/resume/` are valid,
but the latter is the canonical version and is defined by a http link header.

_Updated August 2018_
