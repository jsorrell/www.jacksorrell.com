$(function () {
  $('[data-toggle="popover"]').popover()
})

$('.popover-dismiss').popover({
  trigger: 'focus'
})

$('a.section-row-toggle').click(function(){
	$(this).find('.dropdown-arrow').toggleClass('dropdown-arrow-flipped');
});
