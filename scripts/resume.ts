/// <reference types="jquery" />
/// <reference types="popper.js" />
/// <reference types="bootstrap" />

$(function() {
	$('[data-toggle="popover"]').popover()

	$(document).on("click", function(e: JQuery.Event<Element, null>) {
		$('[data-toggle="popover"],[data-original-title]').each(function() {
			if (!$(this).is(e.target) && $(this).has(e.target).length === 0 && $('.popover').has(e.target).length === 0) {
				(($(this).popover('hide').data('bs.popover') || {}).inState || {}).click = false
			}
		});
	});

	$('a.section-row-toggle').on("click", function() {
		$(this).find('.dropdown-arrow').toggleClass('dropdown-arrow-flipped');
	});
})
