$( "#movieSearchResult" ).css("display", "block");
$( "#searchResultTitle" ).text("<%= movie.Title %>");
$( "#searchResultImdbId" ).text("<%= movie.ImdbID %>");
$( "#searchResultPlot" ).html("<%= movie.Plot %>").text();
