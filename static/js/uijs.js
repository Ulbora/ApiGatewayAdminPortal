
// $(document)
// .ready(function() {

$(function () {
    $('.ui.menu .ui.dropdown').dropdown({
        on: 'hover'
    });
    $('.ui.menu a.item')
        .on('click', function () {
            $(this)
                .addClass('active')
                .siblings()
                .removeClass('active')
                ;
        })
        ;
});
// })
// ;

$(function () {
    $('a.item').click(function () {
        $('.item').removeClass('active');
        $(this).addClass('active');
    })
});


$(function () {
    $('.activating.element')
    .popup()
  ;
});

