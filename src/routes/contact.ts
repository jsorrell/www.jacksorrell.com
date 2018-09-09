import * as express from 'express';
const router = express.Router();
import * as bodyParser from 'body-parser';
import mailgun = require('mailgun-js');
import * as validator from 'validator';
import { logger } from '../logger/logger';
import { CONTACT } from '../constants';

if (!(process.env.MAILGUN_API_KEY && process.env.MAILGUN_DOMAIN && process.env.MAILGUN_TO_ADDRESS && process.env.MAILGUN_SUBJECT)) {
	console.error('Need MAILGUN_API_KEY, MAILGUN_DOMAIN, MAILGUN_TO_ADDRESS, and MAILGUN_SUBJECT set in environment.');
	process.exit(1);
}

const mg = mailgun({
	apiKey: process.env.MAILGUN_API_KEY as string,
	domain: process.env.MAILGUN_DOMAIN as string
});

const canonicalPath = '/contact/';

if (!process.env.GRECAPTCHA_SITE_KEY) {
	console.error('Need GRECAPTCHA_SITE_KEY set in environment.');
}

const pushHeader = [
	{ link: '/css/style.css', type: 'style' }
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
	res.setHeader('Link', `<${req.protocol}://${host + canonicalPath}>; rel="canonical", ` + pushHeader);
	res.render('contact', { contactData: CONTACT });
});

const parser = bodyParser.urlencoded({
	extended: false
});

/*
 * Handle message submissions.
 */
router.post('/', parser, function (req, res) {
	// Ensure honeypot not used
	if (req.body.website) {
		res.status(400).send('Go away bot!');
		return;
	}
	// Ensure body exists
	if (!isEmailMessage(req.body)) {
		res.status(400).send('Invalid submission.');
		return;
	}

	// Ensure message valid
	let messageStatus = emailMessageInvalid(req.body);
	if (messageStatus) {
		res.status(400).send(messageStatus);
		return;
	}

	// In development, do not send actual email.
	if (process.env.NODE_ENV === 'development') {
		logger.info('Email sending withheld in dev.');
	 	res.status(200).send('Message Received');
		return;
	}

	sendEmail(req.body).then((_resp) => {
		logger.info('Email sent.');
		res.status(200).send('Message Received');
	}).catch((err) => {
		logger.warn('Message send failed:', err);
		res.status(500).send('');
	});
});

interface EmailMessage {
	name: string;
	email: string;
	message: string;
}

/**
 * Check if the body sent is an email message.
 * @param  body the body sent.
 * @return      a boolean whether the body is a message.
 */
function isEmailMessage (body: any): body is EmailMessage {
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

	return true;
}

/**
 * Check if the message sent is invalid.
 * @param  message the EmailMessage sent.
 * @return         the reason for invalidity or false if the message is valid.
 */
function emailMessageInvalid (message: EmailMessage): string | false {
	/* Name */
	if (!validator.isLength(message.name, { min: 1, max: CONTACT.maxNameLength })) {
		let resp = `Name length of ${message.name.length} is invalid.`;
		logger.info(resp);
		return resp;
	}

	/* Email */
	if (!validator.isLength(message.email, { min: 1, max: CONTACT.maxEmailLength })) {
		let resp = `Email length of ${message.email.length} is invalid.`;
		logger.info(resp);
		return resp;
	}

	if (!validator.isEmail(message.email)) {
		let resp = `Email "${message.email}" is not a valid email.`;
		logger.info(resp);
		return resp;
	}

	/* Message */
	// fix crlf counted as 2 characters TODO: do this better
	message.message = message.message.replace(/\r\n/g, '\n');

	if (!validator.isLength(message.message, { min: 1, max: CONTACT.maxMessageLength })) {
		let resp = `Message length of ${message.message.length} is invalid.`;
		logger.info(resp);
		return resp;
	}

	return false;
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
