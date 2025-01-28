var sent = false;
window.onscroll = function() {
    if (!sent) {
        var a = document.getElementById("autosplash");
        var y = window.scrollY;
        if (y <= 1) {
            y = 1;
        }
        var f = (window.innerHeight/(y*40))
        if (f <= 0.031) {
            f = 0;
        }
        a.style.opacity = f;
    } 
}
