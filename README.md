# www.jacksorrell.com
My personal website at https://www.jacksorrell.com

## Setup
1. Clone repository or download branch.
1. Ensure npm is v5+ and node is v8+.
1. `npm install`
1. Add config files `lib/config/captcha.js` and `lib/config/mailgun.js` using the examples.
1. Start with (remember to set port)
	* **Production:** `NODE_ENV="production" PORT=#### npm start`
	* **Development:** `NODE_ENV="development" DEBUG="jacksorrell.com:*" PORT=#### npm start`
	* **Systemd:**

		```
		[Unit]
		Description=www.jacksorrell.com server
		After=network.target

		[Service]
		User=example_user
		Group=example_group
		WorkingDirectory=/var/www/www.jacksorrell.com
		ExecStart=/usr/bin/npm start
		Restart=on-failure
		Environment=NODE_ENV=production
		Environment=PORT=####

		[Install]
		WantedBy=multi-user.target
		```

## Design Choices
* Original favicon created using font [The Wastes of Space](http://www.fontspace.com/chequered-ink/the-wastes-of-space).
* Favicons from https://realfavicongenerator.net.
* Use contact form to protect email from spambots.
* Contact form sends email via Mailgun and is protected from bots by Google's invisible recaptcha.
* Both `www.jacksorrell.com/resume` and `www.jacksorrell.com/resume/` are valid, but the latter is the canonical version and is defined by a http link header.

*Updated February 2018*
