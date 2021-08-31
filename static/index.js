$.ajax({
    url: "/api/posts",
    type: "get",
    beforeSend: function() {
        $("#spinner").show();
    },
    success: function(resp){
        $("#spinner").hide()
    }, 
    complete: function(data){
        $("#spinner").hide()
    }
});
