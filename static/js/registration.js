const passwordInput = document.getElementById("password");
const strengthText = document.getElementById("strength-text");

passwordInput.addEventListener("input", () => {
    const password = passwordInput.value;
    let score = 0;

    if (password.length >= 8) score++;
    if (/[A-Z]/.test(password)) score++;
    if (/[a-z]/.test(password)) score++;
    if (/[0-9]/.test(password)) score++;
    if (/[^A-Za-z0-9]/.test(password)) score++;

    let strength = "";
    let width = "";
    let color = "";

    switch (score) {
        case 0:
        case 1:
            strength = "Very Weak";
            width = "20%";
            color = "red";
            break
        case 2:
            strength = "Weak";
            width = "40%";
            color = "orange";
            break
        case 3:
            strength = "Moderate";
            width = "60%";
            color = "yellow";
            break
        case 4:
            strength = "Strong"
            width = "80%";
            color = "lightgreen";
            break
        case 5:
            strength = "Very Strong";
            width = "100%";
            color = "green";
            break
    }

    strengthText.textContent = strength;
    strengthText.style.color = color;
    passwordInput.style.borderColor = color;
    passwordInput.style.boxShadow = `0 0 10px ${color}`;
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