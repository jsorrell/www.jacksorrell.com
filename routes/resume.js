var express = require('express');
var router = express.Router();
var path = require('path');
var resumeData = require('../resume/resume_data.json')

const canonicalPath = '/resume/';

/* GET users listing. */
router.get('/', function(req, res, next) {
	var host = req.get('Host');
	res.setHeader('Link', '<' + req.protocol + '//' + host + canonicalPath + '>; rel="canonical"');

	res.locals.moment = require('moment');
	res.render('resume', resumeData);
});

module.exports = router;
