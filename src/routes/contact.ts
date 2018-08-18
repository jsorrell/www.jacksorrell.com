import * as express from 'express';
const router = express.Router();
import * as bodyParser from 'body-parser';
import mailgun = require('mailgun-js');
import verifyRecaptcha from '../captcha/gRecaptchaVerify';
import * as validator from 'validator';
import { logger } from '../logger/logger';

if (!(process.env.MAILGUN_API_KEY && process.env.MAILGUN_DOMAIN && process.env.MAILGUN_TO_ADDRESS && process.env.MAILGUN_SUBJECT)) {
	console.error('Need MAILGUN_API_KEY, MAILGUN_DOMAIN, MAILGUN_TO_ADDRESS, and MAILGUN_SUBJECT set in environment.');
	process.exit(1);
}

const mg = mailgun({
	apiKey: process.env.MAILGUN_API_KEY as string,
	domain: process.env.MAILGUN_DOMAIN as string
});

const canonicalPath = '/contact/';

const maxEmailLength = Number(process.env.CONTACT_MAX_EMAIL_ADDRESS_LENGTH) || 254;
const maxNameLength = Number(process.env.CONTACT_MAX_NAME_LENGTH) || 70;
const maxMessageLength = Number(process.env.CONTACT_MAX_MESSAGE_LENGTH) || 10000;

if (!process.env.GRECAPTCHA_SITE_KEY) {
	console.error('Need GRECAPTCHA_SITE_KEY set in environment.');
}

const renderVars = {
	gRecaptchaSiteKey: process.env.GRECAPTCHA_SITE_KEY,
	maxEmailLength: maxEmailLength,
	maxNameLength: maxNameLength,
	maxMessageLength: maxMessageLength
};

const pushHeader = [
	{ link: '/css/style.css', type: 'style' },
	{ link: '/js/contact.js', type: 'script' }
].map((asset) => {
	return `<${asset.link}>; as=${asset.type}; rel=preload`;
}).reduce((acc, val) => {
	return acc + ', ' + val;
});

/*
 * Handle page load
 */
router.get('/', function (req, res) {
	const host = req.get('Host');
	res.setHeader('Link', `<${req.protocol}://${host + canonicalPath}>; rel="canonical, "` + pushHeader);
	res.render('contact', renderVars);
});

const parser = bodyParser.urlencoded({
	extended: false
});

/*
 * Handle message submissions.
 */
router.post('/', parser, function (req, res) {
	const gRecaptchaResponse = req.body['g-recaptcha-response'];
	verifyRecaptcha(gRecaptchaResponse,
		function valid () {
			// Ensure body exists
			if (!isValidBody(req.body)) {
				res.status(400).send('Invalid message.');
				return;
			}
			sendEmail(req.body).then((_resp) => {
				logger.info('Email sent.');
				res.status(200).send('Message Received');
			}).catch((err) => {
				logger.warn('Message send failed:', err);
				res.status(500).send('');
			});
		},
		function invalid () { res.status(400).send('Captcha Validation Failed'); }
	);
});

interface EmailMessage {
	name: string;
	email: string;
	message: string;
}

/**
 * Check if the body sent is a valid email message.
 * @param  body the body sent.
 * @return      a boolean indicating if the message is valid.
 */
function isValidBody (body: any): body is EmailMessage {
	if (!(body.name)) {
		logger.info('No name in submission.');
		return false;
	}

	if (!body.email) {
		logger.info('No email in submission.');
		return false;
	}

	if (!body.message) {
		logger.info('No message in submission.');
		return false;
	}

	/* Name */
	if (!validator.isLength(body.name, { min: 1, max: maxNameLength })) {
		logger.info(`Name length of ${body.name.length} invalid.`);
		return false;
	}

	/* Email */
	if (!validator.isLength(body.email, { min: 1, max: maxEmailLength })) {
		logger.info(`Email length of ${body.email.length} invalid.`);
		return false;
	}

	if (!validator.isEmail(body.email)) {
		logger.info(`Email ${body.email} not valid email.`);
		return false;
	}

	/* Message */
	// fix crlf counted as 2 characters TODO: do this better
	body.message = body.message.replace(/\r\n/g, '\n');

	if (!validator.isLength(body.message, { min: 1, max: maxMessageLength })) {
		logger.info(`Message length of ${body.message.length} invalid.`);
		return false;
	}

	return true;
}

/**
 * Send an email to the admin.
 * @param  email the email to send.
 * @return       a promise containing a response from the mailgun server.
 */
async function sendEmail (email: EmailMessage) {
	const emailData = {
		from: `${email.name}<${email.email}>`,
		to: process.env.MAILGUN_TO_ADDRESS as string,
		subject: process.env.MAILGUN_SUBJECT as string,
		text: email.message
	};

	logger.info('Sending email.');
	return mg.messages().send(emailData);
}

export default router;
