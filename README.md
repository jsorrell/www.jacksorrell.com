# www.jacksorrell.com
[![js-happiness-style](https://img.shields.io/badge/code%20style-happiness-brightgreen.svg)](https://github.com/JedWatson/happiness)

My personal website at <https://www.jacksorrell.com>

## Configuration
Run `./www.jacksorrell.com genconfig` to generate a sample configuration yaml file.
Modify this config file as needed.

All config options can also be passed as flags. Run `./www.jacksorrell.com help` for more information.

### _Options_
```yaml
# The commented out fields are reqired

# General
logLevel: warn # Valid Levels: panic, fatal, error, warn, warning, info, debug

# Server
server:
  port: 3000

# Contact Form
contact:
  mailgun:
    # publicValidationKey: pubkey-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    # privateAPIKey: key-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
  email:
    # domain: mg.example.com
    # toAddress: contact@example.com
    subject: Contact Form Message
  # Maximum input lengths
  maxLengths:
    name: 70
    email: 254
    message: 10000
```

## Building from Source

Node.js and NPM required to generate assets.

### Generate Assets

`npx gulp && go generate`

### Compile

`go build`

All assets will be embedded in the executable (except for the configuration).

## Development

Node.js and NPM required to generate assets.

### Development Mode

To automatically recompile assets and reload the browser
on changes using browsersync, run `npx gulp dev`

Compile using `go install -tags 'dev'`. This will not cache any assets. They will be reread from disk.

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

-   Favicons automatically generated.

-   Use contact form to protect email from spambots.

-   Contact form sends email via Mailgun and is protected from bots
by a honeypot input.

-   `www.jacksorrell.com/resume` permanently redirects to
`www.jacksorrell.com/resume/`.

-   Embedding assets allows for portability.

_Updated September 2018_
