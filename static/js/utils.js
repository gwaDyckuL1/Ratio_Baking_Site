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