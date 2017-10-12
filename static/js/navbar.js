$(function() {

    function seeOrDelNavbar(name) {
        name.on("click", function(e) {
            e.preventDefault()
            href = $(this).attr('href')
                // console.log(href, $('this').parent())
            parent = $(this).parent()
            $.post(href)
                .done(function() {
                    console.log("ok")
                    parent.slideUp()
                })
                .fail(function() {
                    console.log("erreur")
                })
        })
    }

    seeOrDelNavbar($(".navbar-vu"))
    seeOrDelNavbar($(".navbar-del"))

    $('.close-modal').on('click', function() {
        setTimeout(function() {
            location.reload(true)
        }, 500)
    })

});