$('#contact-form').on('submit', function (e) {
	e.preventDefault();
	$('#contact-submit .loading-spinner').addClass('spinning');
	$('#contact-submit').prop('disabled', true);
	$('.alert').hide(); // hide alerts
	grecaptcha.reset();
	grecaptcha.execute();
});

$('.close-alert').on('click', function (_e) {
	$(this).closest('.alert').hide();
});

export function contactOnSubmit (_token: string) {
	$.ajax({
		data: $('#contact-form').serialize(),
		type: 'POST',
		complete: function () {
			$('#contact-submit .loading-spinner').removeClass('spinning');
			$('#contact-submit').prop('disabled', false);
		},
		success: function (_data, _textStatus, _jqXHR) {
			$('#message-success-alert').show();
		},
		error: function (_jqXHR, _textStatus, _errorThrown) {
			$('#message-error-alert').show();
		}
	});
}
