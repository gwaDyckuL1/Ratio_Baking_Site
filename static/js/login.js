document.getElementById("login").addEventListener("submit", async function(e) {
    e.preventDefault();

    hide("login-error");

    const formData = new FormData(e.target);
    const action = e.submitter.value;

    if (action === "login") {
        const result = await fetch("/loginSubmit", {
            method: "POST",
            headers: {
                "Accept": "application/json"
            },
            body: formData
        });
        const data = await result.json();

        if (data.ok) {
            window.location.href = "/";
        } else {
            show("login-error");
        }
    }

    if (action === "forgot") {
        // Need to add await and do backend email sending still.
        const result = await fetch("/forgotLoginSubmit", {
            method: "Post",
            headers: {
                "Accept": "application/json"
            },
            body: formData
        });
        const data = await result.json();

        
        show("message");
        setTimeout(() => {
            window.location.href = "/";        
            
        },2000)
    }
})
