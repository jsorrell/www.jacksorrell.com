require('app-module-path').addPath(__dirname + '/lib');
var express = require('express');
var path = require('path');
var favicon = require('serve-favicon');
var sassMiddleware = require('node-sass-middleware');
var resume = require('./routes/resume');
var contact = require('./routes/contact');
var redirect = require('express-simple-redirect');

var app = express();


// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'pug');

app.use(favicon(path.join(__dirname, 'favicon/public', 'favicon.ico')));

app.use(redirect({
	'/': '/resume'
}));

app.use(
	sassMiddleware({
		src: path.join(__dirname, 'sass'), 
		dest: path.join(__dirname, 'public/stylesheets'),
		debug: process.env.NODE_ENV === 'development',
		outputStyle: 'compressed',
		prefix: '/stylesheets'
	})
);
app.use(express.static(path.join(__dirname, 'public')));
app.use(express.static(path.join(__dirname, 'favicon/public')));

app.use('/resume', resume);
app.use('/contact', contact);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
	var err = new Error('Not Found');
	err.status = 404;
	next(err);
});

// error handler
app.use(function(err, req, res, next) {
	res.locals.status = err.status || 500;
	switch (res.locals.status) {
		case 403:
		res.locals.message = "Hey! Stop! You're not allowed here!";
		break;
		case 404:
		res.locals.message = "This is not the page you are looking for.";
		break;
		case 405:
		res.locals.message = "You can't do it that way!";
		break;
		case 408:
		res.locals.message = "Come on. Hurry up.";
		break;
		case 500:
		res.locals.message = "My bad.";
		break;
		case 502:
		res.locals.message = "Something's broken.";
		break;
		case 504:
		res.locals.message = "Someone in the middle is asleep at work.";
		break;
		default:
		res.locals.message = "Well, that's a weird error.";
	}
	
	// set locals, only providing error in development
	res.locals.error = req.app.get('env') === 'development' ? err : null;

	// render the error page
	res.status(res.locals.status);
	res.render('error');
});

module.exports = app;
