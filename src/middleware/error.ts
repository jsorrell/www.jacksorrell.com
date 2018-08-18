import * as express from 'express';
import { logger } from '../logger/logger';

export enum Status {
	FORBIDDEN = 403,
	NOT_FOUND = 404,
	INTERNAL_SERVER_ERROR = 500
}

function getStatusMessage (status: Status) {
	switch (status) {
		case Status.FORBIDDEN:
			return 'Hey! Stop! You\'re not allowed here!';
		case Status.NOT_FOUND:
			return 'This is not the page you are looking for.';
		case Status.INTERNAL_SERVER_ERROR:
		 	return 'My bad.';
		default:
			logger.warn(`server status ${status} has no status message.`);
			return `Error: ${status}`;
	}
}

export class ServerError extends Error {
	status: Status;

	constructor (status: Status) {
		super(getStatusMessage(status));
		this.status = status;
	}
}

export function handler (err: ServerError, _req: express.Request, res: express.Response, _next: express.NextFunction) {
	res.locals.status = err.status || Status.INTERNAL_SERVER_ERROR;
	res.locals.message = getStatusMessage(res.locals.status);

	// set locals, only providing error in development
	res.locals.error = process.env.NODE_ENV === 'development' ? err : null;

	// render the error page
	res.status(res.locals.status);
	res.render('error');
}
