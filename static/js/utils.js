function hide(id) {
    document.getElementById(id).classList.add("hidden");
}

function show(id) {
    document.getElementById(id).classList.remove("hidden");
}

function changeText(id, text) {
    document.getElementById(id).innerText = text;
}

function makeRequired(id) {
    document.getElementById(id).setAttribute("required", "true");
}

function clearRequired(id) {
    document.getElementById(id).removeAttribute("required");
}

function zeroField(id) {
    document.getElementById(id).value = "";
}

function passwordStrength(password) {
    const passwordBox = document.getElementById("password");
    const strengthText = document.getElementById("strength-text");
    let score = 0;

    if (password.length >= 8) score++;
    if (/[A-Z]/.test(password)) score++;
    if (/[a-z]/.test(password)) score++;
    if (/[0-9]/.test(password)) score++;
    if (/[^A-Za-z0-u]/.test(password)) score++;

    let strength = "";
    let color = "";

    switch (score) {
        case 0:
        case 1:
            strength = "Very Weak";
            color = "red";
            break
        case 2:
            strength = "Weak";
            color = "orange";
            break
        case 3:
            strength = "Moderate";
            color = "yellow";
            break
        case 4:
            strength = "Strong"
            color = "lightgreen";
            break
        case 5:
            strength = "Very Strong";
            color = "green";
            break
    }
    strengthText.textContent = strength;
    strengthText.style.color = color;
    passwordBox.style.borderColor = color;
    passwordBox.style.boxShadow = `0 0 10px ${color}`;
}
