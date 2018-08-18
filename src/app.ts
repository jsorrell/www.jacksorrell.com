/**
 * Module dependencies.
 */

import * as path from 'path';
require('dotenv').config({ path: path.join(__dirname, '../.env') });
import { app } from './server';
import * as http from 'http';
import { logger } from './logger/logger';

/**
 * Get port from environment and store in Express.
 */

const port = process.env.PORT || 3000;
app.set('port', port);

/**
 * Create HTTP server.
 */

const server = http.createServer(app);

/**
 * Listen on provided port, on all network interfaces.
 */

server.listen(port);
server.on('error', onError);
server.on('listening', onListening);

/**
 * Event listener for HTTP server "error" event.
 */

function onError (error: Error & {syscall: string, code: string }) {
	if (error.syscall !== 'listen') {
		throw error;
	}

	// handle specific listen errors with friendly messages
	switch (error.code) {
		case 'EACCES':
			console.error(`Port ${port} requires elevated privileges`);
			process.exit(1);
			break;
		case 'EADDRINUSE':
			console.error(`Port ${port} is already in use`);
			process.exit(1);
			break;
		default:
			throw error;
	}
}

/**
 * Event listener for HTTP server "listening" event.
 */

function onListening () {
	logger.info(`Listening on Port ${port}`);
}
