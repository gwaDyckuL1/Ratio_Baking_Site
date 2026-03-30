let passwordInput = document.getElementById("password")

passwordInput.addEventListener("input", () => {
    const password = passwordInput.value;

     passwordStrength(password);

})

document.getElementById("registration").addEventListener("submit", async function(e) {
    e.preventDefault();

    const toHide = ["registration-error", "username-error", "email-error", "password-error"];
    toHide.forEach(id => {
        hide(id);
    })

    const firstPassword = document.getElementById("password").value;
    const checkPass = document.getElementById("check_password").value;

    if (firstPassword !== checkPass) {
        show("password-error");
        changeText("password-error", "Passwords do not match");
        return
    }
    if (firstPassword.length < 8 ) {
        show("password-error");
        changeText("password-error", "Password should be at least 8 characters long.");
        return
    }

    const form = e.target;
    const formData = new FormData(form);
    const result = await fetch("/registrationSubmit", {
        method: "POST",
        headers: {
            "Accept": "application/json"
        },
        body: formData,
    });

    const data = await result.json();

    if (data.ok) {
        console.log("Registration successful");
        show("registration-success");
        setTimeout(() => {
            window.location.href = "/login";
        }, 2000);
    } else {
        const errorMessage = data.field + "-error";
        show(errorMessage);
    }
})
