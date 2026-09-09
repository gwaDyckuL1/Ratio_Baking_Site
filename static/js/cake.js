let cakeSelect;
let panShape;

document.addEventListener("DOMContentLoaded", setup)

function setup() {
  cakeSelect = document.getElementById("calculator-cake");
  panShape = document.getElementById("shape");

  cakeSelect.addEventListener("change", cakeSelectChange);
  panShape.addEventListener("change", panShapeChange);
}

function resetOptions() {
  document.querySelectorAll(".cake-options").forEach(option => {
      option.classList.add("hidden");
  })
  document.querySelectorAll(".allRecipe").forEach(ingredient => {
      ingredient.classList.add("hidden");
  })
  document.querySelectorAll("[required]").forEach(element => {
      element.removeAttribute("required");
  })
  document.querySelectorAll(".pan-shape").forEach(element => {
    hide(element.id);
  })
}

function panShapeChange() {
  document.querySelectorAll(".pan-shape").forEach(element => {
    hide(element.id);
  })
  switch (panShape.value) {
    case "square":
      show("measure");
      show("square");
      show("depth-view");
      break;
    case "round":
      show("measure");
      show("round");
      show("depth-view");
      break;
    case "weird":
      show("weird");
      hide("measure");
      break;
  }  
  show("cake-submission");
}

function cakeSelectChange() {
  resetOptions();
  panShape.value = "";
  hide("cake-submission");
  show("pan-dimension");
  switch (cakeSelect.value) {
    case "free-form":
      show("free-form");
      document.querySelectorAll(".allRecipe").forEach(element => {
        show(element.id);
        console.log(element.id);
      })
      break;
    case "high-ratio-cake":
      show("high-ratio-cake");
      break;
    case "pound-cake":
      show("pound-cake");
      break;
    case "sponge-cake":
      show("sponge-cake");
      break;
  }
}
