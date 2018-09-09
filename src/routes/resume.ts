import { Router } from 'express';
import { default as resumeData } from '../resume/resume_data';
import { CONTACT } from '../constants';

const router = Router();
const canonicalPath = '/resume/';
const pushHeader = [
	{ link: '/css/style.css', type: 'style' },
	{ link: '/js/contact.js', type: 'script' },
	{ link: '/images/person.svg', type: 'image' },
	{ link: '/images/beaker.svg', type: 'image' },
	{ link: '/images/paperclip.svg', type: 'image' },
	{ link: '/images/briefcase.svg', type: 'image' },
	{ link: '/images/fork.svg', type: 'image' },
	{ link: '/images/terminal.svg', type: 'image' },
	// { link: '/images/myface-nobg.png', type: 'image' },
	{ link: '/images/octocat.svg', type: 'image' },
	{ link: '/images/Twitter_Social_Icon_Circle_Color.svg', type: 'image' },
	{ link: '/images/keybase_logo_official.svg', type: 'image' }
].map((asset) => {
	return `<${asset.link}>; as=${asset.type}; rel=preload`;
}).reduce((acc, val) => {
	return acc + ', ' + val;
});

/* GET users listing. */
export default router.get('/', function (req, res, _next) {
	const host = req.get('Host');
	res.setHeader('Link', `<${req.protocol}://${host + canonicalPath}>; rel="canonical", ` + pushHeader);

	res.locals.moment = require('moment');
	res.render('resume', { resumeData: resumeData, contactData: CONTACT });
});
