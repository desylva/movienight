$( "#movieSearchResult" ).replaceWith('<%= partial("movies/omdb.html") %>');
$( "#movieSearchResult" ).css("display", "block");

function fillFormFields(id) {
    $ ( "#movie-ImdbID" ).val(id);
    name = $( "#movie-"+id ).children('h2').text();
    $ ( "#movie-Name" ).val(name);
}