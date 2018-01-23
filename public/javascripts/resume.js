$(function () {
	$('[data-toggle="popover"]').popover()
	
	$(document).on('click', function (e) {
		$('[data-toggle="popover"],[data-original-title]').each(function () {
			if (!$(this).is(e.target) && $('.popover').has(e.target).length === 0) {                
				(($(this).popover('hide').data('bs.popover')||{}).inState||{}).click = false
			}
		});
	});

	$('a.section-row-toggle').click(function(){
		$(this).find('.dropdown-arrow').toggleClass('dropdown-arrow-flipped');
	});
})

