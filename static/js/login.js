document.getElementById("login").addEventListener("submit", async function(e) {
    e.preventDefault();

    hide("login-error");

    const formData = new FormData(e.target);
    const result = await fetch("/loginSubmit", {
        method: "POST",
        headers: {
            "Accept": "application/json"
        },
        body: formData
    });
    const data = await result.json();
    console.log(data);
    if (data.ok) {
        window.location.href = "/";
    } else {
        show("login-error");
    }
})