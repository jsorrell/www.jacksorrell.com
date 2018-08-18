import { Router } from 'express';
import { default as resumeData } from '../resume/resume_data';

const router = Router();
const canonicalPath = '/resume/';

/* GET users listing. */
export default router.get('/', function (req, res, _next) {
	const host = req.get('Host');
	res.setHeader('Link', '<' + req.protocol + '//' + host + canonicalPath + '>; rel="canonical"');

	res.locals.moment = require('moment');
	res.render('resume', resumeData);
});
