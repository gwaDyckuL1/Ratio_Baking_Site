document.getElementById("saveRecipe").addEventListener("submit", async function(e) {
    e.preventDefault();

    const formData = new FormData(e.target);
    const result = await fetch("/saveRecipe", {
        method: "POST",
        headers: {
            "Accepted": "application/json"
        },
        body: formData
    })

    const data = await result.json();
    if (data.ok) {
        show("success")
    } else {
        show("save-error")
    }
})
