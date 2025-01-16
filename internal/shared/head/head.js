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

async function getExample(view) {
    const response = await fetch("/api/getExample", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({special:view}),
    });

    let res = await response.json();
    if (res.success == "true") {
        // do stuff
    } else {
        console.log("error");
    }
}

window.onscroll = function() {
    var a = document.getElementById("splash-inner");
    var y = window.scrollY;
    var f = ((window.innerHeight/20)/y)
    // a.style.backdropFilter = "invert(" + f + ")";
    // a.style.filter = "invert(" + f + ")" + " sepia(" + f + ")" + "hue-rotate(" + (f+f)*300 + "deg)";
    if (f <= 0.04) {
        f = 0;
    }
    a.style.opacity = f;
}
