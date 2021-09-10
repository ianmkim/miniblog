
RGBA(`
float channel(vec2 p, float t) {
  float a = atan(p.x, -p.y);
  float w = sin(a*8.0 + t*2.0)*sin(t+a);
  float d = length(p) - w*0.05* smoothstep(0.85, 1.4, abs(a*0.5));
  d = abs(d - 0.25);
  return smoothstep(0.005, 0.0005, d);
}

void main() {
  vec2 p = gl_FragCoord.xy/resolution-0.5;
  p.x *= 0.6;
  p.y *= 0.6;
  gl_FragColor = vec4(
    channel(p, time*0.7),
    channel(p, time*0.9+1.0),
    channel(p, time*1.1+2.0),
    1.0);
}`, {record:false, target: document.getElementById("graphic")});


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
