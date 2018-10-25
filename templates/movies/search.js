$( "#movie-Name" ).val("<%= movie.Title %>");
$( "#movie-Imdb" ).val("<%= movie.ImdbID %>");

$( "#movieSearchResult" ).css("display", "block");
$( "#searchResultTitle" ).text("<%= movie.Title %>");
$( "#searchResultYear" ).text("<%= movie.Year %>");
$( "#searchResultPlot" ).html("<%= movie.Plot %>").text();
