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
        var a = document.getElementById("splash-inner");
        var w = document.getElementById("spw");
        var y = window.scrollY;
        var f = ((window.innerHeight/50)/y)
        if (f <= 0.09) {
            f = 0;
        }
        if (f >= 0.9) {
            f = 1;
        }
        a.style.opacity = f;
        w.style.opacity = f;
    } 
}
