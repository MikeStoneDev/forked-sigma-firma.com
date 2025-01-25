var sent = false;
async function postContact() {
    if (!sent) {
        var cbutt = document.getElementById("contact-butt");
        var fd = document.getElementById("formData");
        var formData = new FormData(fd);

        if (validateFormData(formData)) {
            var fd_s =     JSON.stringify(Object.fromEntries(formData));

            cbutt.innerHTML = "sending...";
            fd.style.filter = "blur(6px)"
            cbutt.style.backgroundColor = "#c05c3f";
            const response = await fetch("/contact", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: fd_s,
            });

            let res = await response.json();
            if (res.success == "true") {
                cbutt.innerHTML = "Thanks!"
                fd.innerHTML = ""
                fd.style.padding = "0"
                sent = true;
            } else {
                cbutt.innerHTML = "error"
                fd.style.filter = "unset"
            }
        }

    } else {

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
            case 'last_name':
            case 'phone':
                if (value.trim() === '') {
                    console.error(key+' field is required.');
                    isValid = false;
                }
                break;

            case 'email':
                if (!validateEmail(value)) {
                    console.error('Invalid email format.');
                    isValid = false;
                }
                break;

                // Add more validation rules for other fields as needed
        }
    }

    return isValid;
}

function validateEmail(email) {
    // Use a regular expression for email validation
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
}
