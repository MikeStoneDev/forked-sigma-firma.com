var cbutt = document.getElementById("contact-butt");
async function postContact() {
    if (!sent) {
        var fd = document.getElementById("formData");
        var formData = new FormData(fd);

        // cbutt.style.backgroundColor = "#ffdece";
        if (validateFormData(formData)) {
            var fd_s =     JSON.stringify(Object.fromEntries(formData));

            cbutt.innerHTML = "sending...";
            fd.style.filter = "blur(6px)";
            const response = await fetch("/contact", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: fd_s,
            });

            let res = await response.json();
            if (res.success == "true") {
                var splash = document.getElementById("splash-inner");
                splash.onclick = function() {
                    window.location = "https://sigma-firma.com/";
                }
                document.getElementById("term-command").innerHTML = "welcome --to=destiny;";

                splash.style.backgroundImage = "url(public/media/hubble.jpg)";
                splash.style.backgroundSize = "contain";
                splash.style.zIndex = "-2";
                splash.style.filter = "unset";
                splash.style.opacity = "1";
                splash.classList.add("splash-zoom");
                var about = document.getElementById("about-outer");
                about.innerHTML = "";
                var contact = document.getElementById("contact-outer");
                contact.innerHTML = "";
                var foot = document.getElementById("footer-outer");
                foot.innerHTML = "";
                sent = true;
            } else {
                fd.style.filter = "unset";
            }
        } else {
            cbutt.classList.add("invalid-butt");
        }
    }
}
var form = document.getElementById("formData");
var formData_ = new FormData(form);
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

function validateFormData(formData) {
    // Check if formData is an instance of FormData
    if (!(formData instanceof FormData)) {
        throw new Error('Invalid argument: formData must be an instance of FormData.');
    }
    let isValid = true;

    // Iterate over the form data entries
    for (let [key, value] of formData.entries()) {

        // Perform validation based on field name
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
                console.log(key, value);
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
    // Use a regular expression for email validation
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
}
