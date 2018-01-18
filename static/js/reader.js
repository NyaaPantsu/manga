var simplemde = new SimpleMDE({
    element: $("#description")[0]
});
document.addEventListener("keyup", function(e) {
    var key = e.which || e.keyCode;
    switch (key) {
        //left arrow
        case 37:
            document.getElementById("prevLink").click();
            break;
            //right arrow
        case 39:
            document.getElementById("nextLink").click();
            break;
    }
});