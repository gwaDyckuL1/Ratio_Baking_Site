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
        document.getElementById("registration-success").innerText = data.Message;
        document.getElementById("registration-success").display = "block";
        form.reset();
    } else {
        if (data.Field) {
            const errorMessage = document.getElementById(data.Field + "-error");
            errorMessage.innerText = data.Message;
            errorMessage.style.display = "block";
        } else {
            alert(data.Message);
        }
    }
})