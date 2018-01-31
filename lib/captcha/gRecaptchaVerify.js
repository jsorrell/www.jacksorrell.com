var request = require('request');
const recaptchaKeys = require('config/captcha');

module.exports = function (grecaptchaResponse, validCallback, invalidCallback) {

	var requestData = {
		secret: recaptchaKeys.secretKey,
		response: grecaptchaResponse
	}

	request.post('https://www.google.com/recaptcha/api/siteverify',
		{
			form: requestData
		},
		function (error, response, body) {
			if (error || response.statusCode != 200 || !(JSON.parse(body).success) ) {
				invalidCallback(false);
			} else {
				validCallback(true);
			}
		});
}
