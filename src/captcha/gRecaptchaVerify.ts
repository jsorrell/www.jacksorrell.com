import * as request from 'request';

export default function (grecaptchaResponse: string, validCallback: () => void, invalidCallback: () => void) {
	const requestData = {
		secret: process.env.GRECAPTCHA_SECRET_KEY,
		response: grecaptchaResponse
	};

	request.post('https://www.google.com/recaptcha/api/siteverify', {
		form: requestData
	},
		function (error, response, body) {
			if (error || response.statusCode !== 200 || !(JSON.parse(body).success)) {
				invalidCallback();
			} else {
				validCallback();
			}
		});
}
