import * as winston from 'winston';

export const logger = winston.createLogger({
	level: process.env.LOG_LEVEL || 'warn',
	transports: [
		new winston.transports.Console({
			format: winston.format.combine(winston.format.colorize(), winston.format.simple())
		})
	]
});
