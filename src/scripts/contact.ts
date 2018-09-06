let sendingAnimation = false;

let contact = document.getElementById('contact') as HTMLDivElement;
let contactBox = document.getElementById('contact-box') as HTMLDivElement;
let contactForm = document.getElementById('contact-form') as HTMLFormElement;
let submitButton = document.getElementById('contact-submit-button') as HTMLButtonElement;
let background = document.getElementById('contact-background') as HTMLDivElement;
let spinners = document.getElementsByClassName('loading-spinner') as HTMLCollectionOf<HTMLObjectElement>;
let openers = document.getElementsByClassName('contact-opener');

/* On Submit */
contactForm.addEventListener('submit', function (e) {
	e.preventDefault();

	submitButton.disabled = true;

	// Show spinner
	for (let i = 0; i < spinners.length; ++i) {
		spinners.item(i).classList.add('spinning');
	}

	let request = new XMLHttpRequest();
	request.open('POST', '/contact/?ajax=1', true);
	request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');

	request.onload = function () {
		if (200 <= this.status && this.status < 400) {
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
	let elementsArr = [].slice.call(contactForm.elements);
	let data = elementsArr.filter((el: Element) => { return el.hasAttribute('name') && !el.hasAttribute('disabled'); })
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
		spinners.item(i).classList.remove('spinning');
	}

	if (clear) {
		contactForm.reset();
	}
}

for (let i = 0; i < openers.length; ++i) {
	openers.item(i).addEventListener('click', function (e) {
		e.preventDefault();
		contact.classList.add('shown');
	});
}

function hideContact () {
	if (!sendingAnimation) contact.classList.remove('sent', 'shown');
}

/* Hide contact on click outside of box or escape press */
background.addEventListener('click', hideContact);
document.onkeydown = function (e) {
	e = e || window.event;
	if (e.key === 'Escape' && contact.classList.contains('shown')) {
		hideContact();
		e.preventDefault();
	}
};
