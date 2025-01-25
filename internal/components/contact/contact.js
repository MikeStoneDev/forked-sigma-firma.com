async function postContact() {
    if (!sent) {
        var cbutt = document.getElementById("contact-butt");
        var fd = document.getElementById("formData");
        var formData = new FormData(fd);

        cbutt.style.backgroundColor = "#c05c3f";
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
                splash.style.backgroundImage= "url(public/media/hubble.jpg)";
                splash.style.zIndex = "-2";
                splash.style.filter = "unset";
                splash.style.opacity = "1";
                var asplash = document.getElementById("autosplash-firma");
                var bl = document.getElementById("bl");
                asplash.style.color = "white";
                bl.style.color = "white";
                bl.style.fontStyle = "italic";
                asplash.style.opacity = "0.8";
                bl.style.opacity = "0.8";
                var about = document.getElementById("about-outer");
                about.innerHTML = "";
                var contact = document.getElementById("contact-outer");
                contact.innerHTML = "";
                sent = true;
            } else {
                cbutt.innerHTML = "error";
                fd.style.filter = "unset";
            }
        }
    }
}
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
                isValid = inputinvalid(key, value)
                break;

            case 'last_name':
                isValid = inputinvalid(key, value)
                break;

            case 'phone':
                isValid = inputinvalid(key, value)
                break;

            case 'email':
                if (!validateEmail(value)) {
                    var cbutt = document.getElementById("contact-butt");
                    console.error('Invalid email format.');
                    cbutt.innerHTML = "every field is required.";
                    isValid = inputinvalid(key, value)
                }
                break;

                // Add more validation rules for other fields as needed
        }
    }

    return isValid;
}

function inputinvalid(key, value) {
    var cbutt = document.getElementById("contact-butt");
    if (value.trim() === '') {
        console.error(key+' field is required.');
        cbutt.innerHTML = "every field is required.";
        var inp = document.getElementById("contact_" +key);
        inp.style.background = "red";
        return false
    }
    return true
}

function validateEmail(email) {
    // Use a regular expression for email validation
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
}
