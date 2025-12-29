//for choices made
let breadSelect;
let tangzhongChoice;
let leavenerChoice;
let panShape;

//for tangzhong checking
let hydration;
let tangzhongPercentage;
let tangzhongRatio;

document.addEventListener("DOMContentLoaded", setup)

function setup() {
    breadSelect = document.getElementById("calculator-bread");
    tangzhongChoice = document.getElementById("tangzhong-choice");
    leavenerChoice = document.getElementById('leavener-choice');
    panShape = document.getElementById("shape");

    breadSelect.addEventListener("change", breadSelectChange);
    tangzhongChoice.addEventListener("change", tangzhongChoiceChange);
    leavenerChoice.addEventListener("change", leavenerChoiceChange);
    panShape.addEventListener("change", panShapeChange);
}

function breadSelectChange() {
    resetOptions();
    switch (breadSelect.value) {
        case "flour-weight":
            show("flour-weight");
            show("flour-view");
            makeRequired("flour");
            break;
        case "total-weight":
            show("total-weight");
            show("dough-weight-view");
            makeRequired("dough-weight");
            break;
        case "pan-dimension":
            show("pan-dimension");
            break;
    }
    if (breadSelect.value != "") {
        document.querySelectorAll(".allRecipe").forEach(ingredient => {
            ingredient.classList.remove("hidden");
        })
        show("bread-submission");
    }
}

function leavenerChoiceChange() {
    hide("sourdough-view");
    hide("yeast-view");
    switch (leavenerChoice.value) {
        case "Sourdough":
            show("sourdough-view");
            break;
        case "Yeast":
            show("yeast-view");
    }
}

function panShapeChange() {
    hide("square");
    hide("round");
    switch (panShape.value) {
        case "square":
            show("square");
            show("depth-view");
            break;
        case "round":
            show("round");
            show("depth-view");
            break;
    }
}

function tangzhongChoiceChange() {
    hydration = document.getElementById("hydration");
    tangzhongPercentage = document.getElementById("tangzhong-percentage");
    tangzhongRatio = document.getElementById("tangzhong-ratio");

    if (tangzhongChoice.value === "No") {
        hide("tangzhong-select");
        hydration.removeEventListener("input", onChanges)
        tangzhongPercentage.removeEventListener("input", onChanges)
        tangzhongRatio.removeEventListener("input", onChanges)
    } else {
        show("tangzhong-select");
        hydration.addEventListener("input", onChanges)
        tangzhongPercentage.addEventListener("input", onChanges)
        tangzhongRatio.addEventListener("input", onChanges)
    }
}

function onChanges() {
    let mainHydration = Number(hydration.value)
    let tangPercentage = Number(tangzhongPercentage.value)
    let tangRatio = Number(tangzhongRatio.value)
    checkHydrationLevels(mainHydration, tangPercentage, tangRatio);
}

function checkHydrationLevels(mainHydration, tangPercentage, tangRatio) {
    hide("all-tangzhong");
    hide("hydration-needed");
    const tangzhongHydration = (tangPercentage / (tangRatio + 1)) * tangRatio;

    if (mainHydration === tangzhongHydration)  {
        show("all-tangzhong");
    }
    if (mainHydration < tangzhongHydration) {
        show("hydration-needed");
    }
}

function resetOptions() {
    document.querySelectorAll(".bread-options").forEach(option => {
        option.classList.add("hidden");
    })

    document.querySelectorAll(".allRecipe").forEach(ingredient => {
        ingredient.classList.add("hidden");
    })
    document.querySelectorAll("[required]").forEach(element => {
        element.removeAttribute("required");
    })
    zeroField("flour");
    zeroField("dough-weight");
    hide("bread-submission");
    hide("depth-view");
    tangzhongChoice.value = "No";
    tangzhongChoiceChange();
}
