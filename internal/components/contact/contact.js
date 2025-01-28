var cbutt = document.getElementById("contact-butt");
var form = document.getElementById("formData");
var formData = new FormData(form);
form.addEventListener("input", function () {
    cbutt.innerHTML = "contact";
    cbutt.classList.remove("invalid-butt");
    for (let [key, value] of formData_.entries()) {
        switch (key) {
            case 'first_name':
                var inp = document.getElementById("contact_" +key);
                inp.classList.remove("invalid");
            case 'last_name':
                var inp = document.getElementById("contact_" +key);
                inp.classList.remove("invalid");
            case 'phone':
                var inp = document.getElementById("contact_" +key);
                inp.classList.remove("invalid");
            case 'email':
                var inp = document.getElementById("contact_" +key);
                inp.classList.remove("invalid");
        }
    }
});
async function postContact() {
    if (!sent) {
        if (validateFormData(formData)) {
            var fd_s =     JSON.stringify(Object.fromEntries(formData));
            cbutt.innerHTML = "sending...";
            form.style.filter = "blur(6px)";

            const response = await fetch("/contact", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: fd_s,
            });

            let res = await response.json();
            if (res.success == "true") {
                document.getElementById("splash-inner").classList.add("splash-zoom");
                document.getElementById("about-outer").innerHTML = "";
                document.getElementById("contact-outer").innerHTML = "";
                document.getElementById("footer-outer").innerHTML = "";
                document.getElementById("autosplash").style.opacity = "1";
                document.getElementById("term-command").innerHTML = "welcome --to=destiny;";
                document.getElementById("term-outer").style.position = "fixed";
                sent = true;
                setTimeout(function() {location.reload()}, 5000);
            } else {
                form.style.filter = "unset";
            }
        } else {
            cbutt.classList.add("invalid-butt");
        }
    }
}
function validateFormData(formData) {
    if (!(formData instanceof FormData)) {
        throw new Error('Invalid argument: formData must be an instance of FormData.');
    }
    let isValid = true;

    // Iterate over the form data entries
    for (let [key, value] of formData.entries()) {
        switch (key) {
            case 'first_name':
                if (!inputinvalid(key, value)) {
                    isValid = false;
                    break;
                }
                break;
            case 'last_name':
                if (!inputinvalid(key, value)) {
                    isValid = false;
                    break;
                }
                break;
            case 'phone':
                if (!inputinvalid(key, value)) {
                    isValid = false;
                    break;
                }
                break;
            case 'email':
                if (validateEmail(value)) {
                    isValid = true
                    break;
                }
                var inp = document.getElementById("contact_" +key);
                inp.classList.add("invalid");
                cbutt.innerHTML = "invalid email";
                isValid = false
                break;
        }
    }
    return isValid;
}
function inputinvalid(key, value) {
    var inp = document.getElementById("contact_" +key);
    if (value.trim() === '') {
        cbutt.innerHTML = "every field is required";
        inp.classList.add("invalid");
        return false
    }
    inp.classList.remove("invalid");
    return true
}
function validateEmail(email) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
}
