var express = require('express');
var router = express.Router();
var path = require('path');
var bodyParser = require('body-parser');
var mailgunConf = require('config/mailgun');
var mailgun = require('mailgun-js')({
	apiKey: mailgunConf.privateApiKey,
	domain: mailgunConf.domain
});
var recaptchaKeys = require('config/captcha');
var verifyRecaptcha = require('captcha/gRecaptchaVerify');
var validator = require('validator');

const canonicalPath = '/contact/';

const maxEmailLength = 254,
	maxNameLength = 70,
	maxMessageLength = 10000;

var renderVars = {
	gRecaptchaSiteKey: recaptchaKeys.siteKey,
	maxEmailLength: maxEmailLength,
	maxNameLength: maxNameLength,
	maxMessageLength: maxMessageLength
};

router.get('/', function(req, res) {
	var host = req.get('Host');
	res.setHeader('Link', '<' + req.protocol + '://' + host + canonicalPath + '>; rel="canonical"');
	res.render('contact', renderVars);
});

var parser = bodyParser.urlencoded({
	extended: false
});
router.post('/', parser, function(req, res) {
	var gRecaptchaResponse = req.body['g-recaptcha-response'];
	verifyRecaptcha(gRecaptchaResponse,
		function() { //valid
			//fix crlf counted as 2 characters TODO: do this better
			req.body.message = req.body.message.replace(/\r\n/g, '\n');

			if (!validator.isLength(req.body.email, {
					min: 1,
					max: maxEmailLength
				}) || !validator.isEmail(req.body.email)) {
				res.status(400).send('Invalid Email');
				return;
			} else if (!validator.isLength(req.body.name, {
					min: 1,
					max: maxNameLength
				})) {
				res.status(400).send('Name Too Long');
				return;
			} else if (!validator.isLength(req.body.message, {
					min: 1,
					max: maxMessageLength
				})) {
				res.status(400).send('Message Too Long');
				return;
			}

			var emailData = {
				from: req.body.name + '<' + req.body.email + '>',
				to: mailgunConf.toAddress,
				subject: mailgunConf.subject,
				text: req.body.message
			};

			mailgun.messages().send(emailData, function(error, body) {
				if (error) {
					console.log(error); //TODO: real log
					res.status(500).send('Failed to Send Message');
				} else {
					res.status(200).send('Message Received');
				}
			});

		},
		function() { //invalid
			res.status(400).send('Captcha Validation Failed');
		});
});

module.exports = router;
