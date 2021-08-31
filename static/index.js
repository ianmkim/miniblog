let create_item = (id, title, author, createdAt) => {
    return `
    <a href="/d/` + id + `">
    <div class="card">
        <div class="card-body">
            <h3>` + title + `</h3>
            <h5>` + author + `</h5>
            <h6>` + createdAt + `</h6>
        </div>
    </div>
    </a>
    `
}

$.ajax({
    url: "/api/posts",
    type: "get",
    beforeSend: function() {
        $("#spinner").show();
    },
    success: function(resp){
        if(resp["success"]){
            let posts = resp["data"]["posts"]
            posts.forEach(post => {
                $("#results").append(create_item(post.id, post.title, post.author, post.createdAt))
            })
            if(posts.length == 0){
                $("#results").append("No posts at the moment");
            }
        }
        $("#spinner").hide()
    }, 
    complete: function(data){
        $("#spinner").hide()
    }
});
