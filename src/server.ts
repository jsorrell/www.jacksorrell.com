import * as path from 'path';
import * as express from 'express';
import * as favicon from 'serve-favicon';
import { default as resume } from './routes/resume';
import { default as contact } from './routes/contact';
import * as error from './middleware/error';

export const app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'pug');
if (process.env.NODE_ENV === 'production') {
	app.set('trust proxy', 'loopback');
}

app.use(favicon(path.join(__dirname, 'public/favicon.ico')));

app.get('/', function (_req, res, _next) {
	res.redirect(307, '/resume/');
});

app.use(express.static(path.join(__dirname, 'public/')));

app.use('/resume/', resume);
app.use('/contact/', contact);

// catch 404 and forward to error handler
app.use(function (_req, _res, next) {
	next(new error.ServerError(error.Status.NOT_FOUND));
});

app.use(error.handler);
