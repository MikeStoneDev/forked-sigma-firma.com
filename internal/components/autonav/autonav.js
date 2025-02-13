let np = document.getElementById("nav-portrait");
function showNavPortrait() {
        let wp = document.getElementById("wrapper");
        np.classList.add("expando");
        wp.classList.add("wrapper-blur");
        setTimeout(function () {
                document.addEventListener('click', tf, false);
        }, 50);
}
function tf() {
        let wp = document.getElementById("wrapper");
        np.classList.remove("expando");
        wp.classList.remove("wrapper-blur");
        document.removeEventListener('click', tf);
}
