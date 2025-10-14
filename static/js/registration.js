document.getElementById("registration").addEventListener("submit", async function(e) {
    e.preventDefault();

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
        console.log("Registration successful")
        document.getElementById("registration-success").innerText = data.message;
        document.getElementById("registration-success").style.display = "block";
        form.reset();
    } else {
        if (data.field) {
            const errorMessage = document.getElementById(data.field + "-error");
            console.log("errorMessage is: ")
            errorMessage.innerText = data.message;
            errorMessage.style.display = "block";
        } else {
            console.log("Calling alert, skipped unique message")
            alert(data.message);
        }
    }
})