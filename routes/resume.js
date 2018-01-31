var express = require('express');
var router = express.Router();
var path = require('path');
var resume_data = require('../resume/resume_data.json')

/* GET users listing. */
router.get('/', function(req, res, next) {
	//not the cleanest
	resume_data.moment = require('moment');
	res.render('resume', resume_data);
});	

module.exports = router;
