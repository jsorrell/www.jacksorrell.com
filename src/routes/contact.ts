import * as express from 'express';
const router = express.Router();
import * as bodyParser from 'body-parser';
import * as mailgun from 'mailgun-js';
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

router.get('/', function (req, res) {
	const host = req.get('Host');
	res.setHeader('Link', '<' + req.protocol + '://' + host + canonicalPath + '>; rel="canonical"');
	res.render('contact', renderVars);
});

const parser = bodyParser.urlencoded({
	extended: false
});

router.post('/', parser, function (req, res) {
	const gRecaptchaResponse = req.body['g-recaptcha-response'];
	verifyRecaptcha(gRecaptchaResponse,
		function valid () { sendMessage(req.body, res); },
		function invalid () { res.status(400).send('Captcha Validation Failed'); }
	);
});

function sendMessage (body: any, res: express.Response) {
	// fix crlf counted as 2 characters TODO: do this better
	body.message = body.message.replace(/\r\n/g, '\n');

	if (!validator.isLength(body.email, {
		min: 1,
		max: maxEmailLength
	}) || !validator.isEmail(body.email)) {
		res.status(400).send('Invalid Email');
		return;
	} else if (!validator.isLength(body.name, {
		min: 1,
		max: maxNameLength
	})) {
		res.status(400).send('Name Too Long');
		return;
	} else if (!validator.isLength(body.message, {
		min: 1,
		max: maxMessageLength
	})) {
		res.status(400).send('Message Too Long');
		return;
	}

	const emailData = {
		from: `${body.name}<${body.email}>`,
		to: process.env.MAILGUN_TO_ADDRESS as string,
		subject: process.env.MAILGUN_SUBJECT as string,
		text: body.message as string
	};

	mg.messages().send(emailData)
	.then((_resp) => {
		res.status(200).send('Message Received');
	}).catch((err) => {
		logger.warn(err);
		res.status(500).send('Failed to Send Message');
	});
}

export default router;
