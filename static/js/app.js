var bcryptIt = function() {
    var textBlock = $("#textBlock").val();
    console.log(textBlock)
    $.ajax({
        url: "/bcrypt-this?textBlock="+textBlock,
        method: "POST",
    }).fail(function(data, status){
        console.log("error: " + status);
        console.log("data: " + data["message"]);
    }).done(function(data, status){
        console.log("success: " + status);
        console.log("data: " + data["bcrypt"]);
    });
}

$("#bcryptBtn").click(bcryptIt);
