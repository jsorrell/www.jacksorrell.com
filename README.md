# www.jacksorrell.com
My personal website at jacksorrell.com

## Setup
1. Clone repository or download branch
2. Ensure npm in v5+ and node is v9+
3. npm install
4. Add config files
```javascript
// lib/config/captcha.js

const personalKeys = {
	siteKey: 'PUT_SITE_KEY_HERE',
	secretKey: 'PUT_SECRET_KEY_HERE'
}

const testingKeys = {
	siteKey: '6LeIxAcTAAAAAJcZVRqyHh71UMIEGNQ_MXjiZKhI',
	secretKey: '6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe'
}

module.exports = process.env.NODE_ENV == 'development' ? testingKeys : personalKeys;
```
```javascript
// lib/config/mailgun.js

module.exports = {
	privateApiKey: 'PUT_PRIVATE_API_KEY_HERE',
	domain: 'PUT_DOMAIN_HERE',
	toAddress: 'PUT_MAILTO_ADDRESS_HERE',
	subject: 'Contact Form Message'
}
```
5. Start with (remember to set port)  
**Production:** `NODE_ENV="production" PORT=#### npm start`  
**Development:** `NODE_ENV="development" DEBUG="jacksorrell.com:*" PORT=#### npm start`


## Design Choices
1. Favicons from https://realfavicongenerator.net using font "The Wastes of Space".
2. Mail on contact form sent via Mailgun and protected using Google's invisible recaptcha.



*Updated February 2018*
