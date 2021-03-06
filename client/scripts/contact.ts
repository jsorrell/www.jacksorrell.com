let sendingAnimation = false;

const contact = document.getElementById('contact') as HTMLDivElement;
const contactBox = document.getElementById('contact-box') as HTMLDivElement;
const contactForm = document.getElementById('contact-form') as HTMLFormElement;
const submitButton = document.getElementById('contact-submit-button') as HTMLButtonElement;
const background = document.getElementById('contact-background') as HTMLDivElement;
const spinners = document.getElementsByClassName('loading-spinner') as HTMLCollectionOf<HTMLObjectElement>;
const openers = document.getElementsByClassName('contact-opener');

/* On Submit */
contactForm.addEventListener('submit', function (e) {
	e.preventDefault();

	submitButton.disabled = true;

	// Show spinner
	for (let i = 0; i < spinners.length; ++i) {
		(spinners.item(i) as HTMLObjectElement).classList.add('spinning');
	}

	const request = new XMLHttpRequest();
	request.open('POST', '/contact/?ajax=1', true);
	request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');

	request.onload = function () {
		if (this.status <= 200 && this.status < 400) {
			// Success (Message sent)
			contact.classList.add('sent');
			sendingAnimation = true;
			contactBox.addEventListener('animationend', function (_e) {
				sendingAnimation = false;
				hideContact();
				resetForm(true);
			});
		} else {
			// Error
			displayError(this.responseText);
			resetForm(false);
		}
	};

	request.onerror = function () {
		// Connection Error
		displayError('Connection Error');
		resetForm(false);
	};

	// Generate data string
	const elementsArr = [].slice.call(contactForm.elements);
	const data = elementsArr.filter((el: Element) => { return el.hasAttribute('name') && !el.hasAttribute('disabled'); })
		.map((el: HTMLInputElement) => { return encodeURIComponent(el.getAttribute('name') as string) + '=' + encodeURIComponent(el.value); })
		.join('&');

	request.send(data);
});

/* Display submission errors */
function displayError (message: string) {
	// TODO make this better
	alert(message);
}

function resetForm (clear: boolean) {
	submitButton.disabled = false;
	for (let i = 0; i < spinners.length; ++i) {
		(spinners.item(i) as HTMLObjectElement).classList.remove('spinning');
	}

	if (clear) {
		contactForm.reset();
	}
}

for (let i = 0; i < openers.length; ++i) {
	(openers.item(i) as HTMLObjectElement).addEventListener('click', function (e) {
		e.preventDefault();
		contact.classList.add('shown');
	});
}

function hideContact () {
	if (!sendingAnimation) contact.classList.remove('sent', 'shown');
}

/* Hide contact on click outside of box or escape press */
background.addEventListener('click', hideContact);
document.addEventListener('keydown', e => {
	if (e.key === 'Escape' && contact.classList.contains('shown')) {
		hideContact();
		e.preventDefault();
	}
});
