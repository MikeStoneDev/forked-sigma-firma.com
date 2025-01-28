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
