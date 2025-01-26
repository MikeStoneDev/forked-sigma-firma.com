let np = document.getElementById("nav-portrait");
function showNavPortrait() {
    np.classList.add("expando");
    setTimeout(function () {
        document.addEventListener('click', tf, false);
    }, 50);
}

function tf() {
    np.classList.remove("expando");
    document.removeEventListener('click', tf);
}
