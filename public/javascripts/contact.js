function contactOnSubmit(token) {
	$.ajax({
		data: $('#contact-form').serialize(),
		type: 'POST',
		complete: function() {
			$('#contact-submit .loading-spinner').removeClass('spinning');
			$('#contact-submit').prop('disabled', false);
		},
		success: function(data, textStatus, jqXHR) {
			$('#message-success-alert').show();
		},
		error: function(jqXHR, textStatus, errorThrown) {
			$('#message-error-alert').show();
		}
	});
}

$('#contact-form').submit(function(e) {
	e.preventDefault();
	$('#contact-submit .loading-spinner').addClass('spinning');
	$('#contact-submit').prop('disabled', true);
	$('.alert').hide(); //hide alerts
	grecaptcha.reset();
	grecaptcha.execute();
});

$('.close-alert').click(function(e) {
	$(this).closest('.alert').hide();
});
