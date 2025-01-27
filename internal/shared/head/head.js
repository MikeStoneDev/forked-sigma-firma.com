function jumpTo(eid) {
    var jump = document.getElementById(eid);
    jump.scrollIntoView({
        behavior: 'smooth',
        block: 'start',
        inline: 'center'
    });
}

function toggleDisplay(elem) {
    let divs = document.getElementById("hiddenTop").children;
    let formDisplay = document.getElementById(elem);
    for (let i=0;i<divs.length;i++) {
        if (divs[i].id != formDisplay.id) {
            divs[i].style.display = "none";
        }
    }
    if (formDisplay.style.display == "none" || formDisplay.style.display == "") {
        formDisplay.style.display = "unset";
    } else {
        formDisplay.style.display = "none";
    }
}


var sent = false;
window.onscroll = function() {
    if (!sent) {
        var a = document.getElementById("autosplash");
        // var w = document.getElementById("spw");
        // var u = document.getElementById("autosplash-firma");
        var y = window.scrollY;
        if (y <= 1) {
            y = 1;
        }
        var f = (window.innerHeight/(y*40))
        if (f <= 0.01) {
            f = 0;
        }
        a.style.opacity = f;
        // w.style.opacity = f;
        // u.style.opacity = f;
    } 
}
