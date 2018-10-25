require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

function fillSearchMovieTitleField() {
    title = $("#movie-Name" ).val()
    $("#searchTitleLink").attr('href','/search/?title='+title)
}
$("#movie-Name" ).change( fillSearchMovieTitleField );


$(() => {

});
