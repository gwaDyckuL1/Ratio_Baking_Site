const graphContainer = document.getElementById("dashBarChart");
const table_body = document.getElementById("table-body")
let recipeData = null;
let counts = {};
let graphStage = 0;
let current_category = "";
let current_subcategory = "";
let current_name = "";

function getRecipeList() {
  let list_to_show = "";
  let filtered = [];
  let recipe_title = document.getElementById("recipe-title");
  let title = "";
  switch (graphStage) {
    case 0:
      filtered = recipeData;
      title = "All Recipes";
      break;
    case 1:
      filtered = recipeData.filter(item => item.category == current_category);
      title = `All ${current_category} Recipes`;
      break;
    case 2:
      filtered = recipeData.filter(item => item.category == current_category && item.subcategory == current_subcategory);
      title = `All ${current_category} Recipes Made By ${current_subcategory}`;
      break;
    case 3:
      filtered = recipeData.filter(item => item.category == current_category &&
        item.subcategory == current_subcategory &&
        item.recipe_name == current_name);
      title = `${current_name} ${current_category} Recipes Made By ${current_subcategory}.`;
  }
  filtered.forEach(item => {
    recipe = JSON.parse(item.recipe_data);
    recipe_line = createRecipeLine(item, recipe);
    list_to_show += recipe_line;
  })
  table_body.innerHTML = list_to_show;
  recipe_title.textContent = title;
}

function createRecipeLine(item, recipe) {
  return recipe_line = `
          <tr>
            <td>${item.recipe_name}</td>
            <td>${recipe.FlourOut}</td>
            <td>${recipe.HydrationOut}</td>
            <td>${recipe.FatOut}</td>
            <td>${recipe.SugarOut}</td>
            <td>${recipe.EggGrams}g / ${recipe.EggWhole} whole eggs</td>
            <td>${recipe.SaltOut}</td>
            <td>${recipe.Leavener}</td>
            <td>${recipe.LeavenerOut}</td>
            <td>${recipe.TangzhongFlour}</td>
            <td>${recipe.TangzhongHydration}</td>
            <td>${item.notes}</td>
          </tr>
        `
}

function backup() {
  graphStage--;
  let result;
  if (graphStage === 0) {
    result = Object.entries(counts).map(([category, value]) => [category, value.total]);
  }
  if (graphStage === 1) {
    result = Object.entries(counts[current_category].subcategories).map(([category, value]) => [category, value.total]);
  }
  if (graphStage === 2) {
    result = Object.entries(counts[current_category].subcategories[current_subcategory].names)
      .map(([category, value]) => [category, value]);
  }
  renderBarChart(result);
}

function renderBarChart(data) {
  getRecipeList();
  graphContainer.innerHTML = "";

  const padding = 25;
  const bottomPadding = 75;
  const fontColor = "white";
  const fontSize = "12";
  const viewWidth = 500;
  const viewHieght = 300;
  const base = viewHieght - bottomPadding;
  const graphWidth = viewWidth - padding;
  const ytick_gap = base / 10;

  const svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
  svg.setAttribute("width", "100%");
  svg.setAttribute("height", "500");
  svg.setAttribute("viewBox", `0 0 ${viewWidth} ${viewHieght}`);

  const tempSQ = document.createElementNS("http://www.w3.org/2000/svg", "rect");
  tempSQ.setAttribute("x", 0);
  tempSQ.setAttribute("y", 0);
  tempSQ.setAttribute("width", "100%");
  tempSQ.setAttribute("height", "500");
  tempSQ.setAttribute("stroke", "blue");
  tempSQ.setAttribute("fill", "blue");

  const yAxis = document.createElementNS("http://www.w3.org/2000/svg", "line");
  yAxis.setAttribute("x1", padding);
  yAxis.setAttribute("y1", padding);
  yAxis.setAttribute("x2", padding);
  yAxis.setAttribute("y2", base);
  yAxis.setAttribute("stroke", "white");

  const yLabel = document.createElementNS("http://www.w3.org/2000/svg", "text");
  yLabel.setAttribute("x", padding / 2);
  yLabel.setAttribute("y", (padding + base) / 2);
  yLabel.setAttribute("text-anchor", "middle");
  yLabel.setAttribute("dominant-baseline", "middle");
  yLabel.setAttribute("fill", fontColor);
  yLabel.textContent = "Percentage";
  yLabel.setAttribute("font-size", fontSize);
  yLabel.setAttribute("transform", `rotate(-90, ${padding / 2}, ${(padding + base) / 2} )`);

  const xAxis = document.createElementNS("http://www.w3.org/2000/svg", "line");
  xAxis.setAttribute("x1", padding);
  xAxis.setAttribute("y1", base);
  xAxis.setAttribute("x2", graphWidth);
  xAxis.setAttribute("y2", base);
  xAxis.setAttribute("stroke", "white");

  const xLabel = document.createElementNS("http://www.w3.org/2000/svg", "text");
  xLabel.setAttribute("x", (graphWidth + padding) / 2);
  xLabel.setAttribute("y", viewHieght - padding / 2);
  xLabel.setAttribute("text-anchor", "middle");
  xLabel.setAttribute("dominant-baseline", "middle");
  xLabel.setAttribute("fill", fontColor);
  xLabel.setAttribute("font-size", fontSize);
  xLabel.textContent = "Recipe Catagories";

  const back_button = document.createElementNS('http://www.w3.org/2000/svg', "rect");
  back_button.setAttribute("x", viewWidth - 50);
  back_button.setAttribute("y", "0");
  back_button.setAttribute('width', "50");
  back_button.setAttribute('height', padding * 0.75);
  back_button.setAttribute("fill", "black");
  back_button.style.cursor = "pointer";
  back_button.addEventListener("click", () => {
    backup();
  });

  const back_button_text = document.createElementNS('http://www.w3.org/2000/svg', "text");
  back_button_text.setAttribute("x", viewWidth - 40);
  back_button_text.setAttribute("y", "10");
  back_button_text.textContent = "BACK";
  back_button_text.setAttribute("fill", "white");
  back_button_text.setAttribute("font-size", fontSize);
  back_button_text.style.cursor = "pointer";
  back_button_text.setAttribute("text-anchor", "start");
  back_button_text.setAttribute("dominant-baseline", "middle");
  back_button_text.addEventListener("click", () => {
    backup();
  })

  svg.appendChild(tempSQ);
  svg.appendChild(yAxis);
  svg.appendChild(yLabel);
  svg.appendChild(xAxis);
  svg.appendChild(xLabel);

  let ytick = base - ytick_gap;
  while (ytick > 0) {
    const tick = document.createElementNS("http://www.w3.org/2000/svg", "line");
    tick.setAttribute("x1", padding - 0.25 * padding);
    tick.setAttribute("y1", ytick);
    tick.setAttribute("x2", padding + 0.25 * padding);
    tick.setAttribute("y2", ytick);
    tick.setAttribute("stroke", "white");
    svg.appendChild(tick);
    ytick -= ytick_gap;
  }

  let categoryCount = 0;
  let totalRecipeCounts = 0;

  data.forEach(item => {
    categoryCount++;
    totalRecipeCounts += item[1];
  })

  const gapBetweenBars = 10;
  const barWidth = ((graphWidth - padding) / categoryCount) - (2 * gapBetweenBars);
  let startX = padding + gapBetweenBars;

  data.forEach(item => { //loop for placing bars
    let name = item[0];
    let value = item[1];
    let barHeight = (viewHieght - padding - bottomPadding) * (value / totalRecipeCounts);

    let bar = document.createElementNS("http://www.w3.org/2000/svg", "rect");
    bar.setAttribute("x", startX);
    bar.setAttribute("y", base - barHeight);
    bar.setAttribute("width", barWidth);
    bar.setAttribute("height", barHeight);
    bar.setAttribute("fill", "brown");
    bar.setAttribute("id", "name");
    bar.setAttribute("class", "bar");
    bar.style.cursor = "pointer";
    bar.addEventListener("click", () => {
      if (graphStage === 0) {
        current_category = name;
      }

      if (graphStage === 1) {
        current_subcategory = name;
      }
      if (graphStage === 2) {
        current_name = name;
      }
      if (graphStage != 3) {
        graphStage++;
        handleBarClicks();
      }
    });

    let barName = document.createElementNS("http://www.w3.org/2000/svg", "text");
    barName.textContent = name;
    barName.setAttribute("x", startX + barWidth / 2);
    barName.setAttribute("y", base + 10);
    barName.setAttribute("text-anchor", "end");
    barName.setAttribute("dominant-baseline", "middle");
    barName.setAttribute("fill", fontColor);
    barName.setAttribute("font-size", fontSize);
    barName.setAttribute("transform", `rotate(-45, ${startX + barWidth / 2}, ${base + 2} )`);

    svg.appendChild(bar);
    svg.appendChild(barName);
    startX += barWidth + gapBetweenBars;
  })
  graphContainer.appendChild(svg);
  if (graphStage > 0) {
    svg.appendChild(back_button);
    svg.appendChild(back_button_text);
  }
}

function handleBarClicks() {
  let result;
  if (graphStage === 1) {
    result = Object.entries(counts[current_category].subcategories)
      .map(([category, value]) => [category, value.total]);
  }

  if (graphStage === 2) {
    result = Object.entries(counts[current_category].subcategories[current_subcategory].names)
      .map(([category, value]) => [category, value]);
  }

  if (graphStage === 3) {
    let value = counts[current_category].subcategories[current_subcategory].names[current_name];
    result = [[current_name, value]];
  }
  renderBarChart(result);
}

document.addEventListener("DOMContentLoaded", async () => {
  const response = await fetch("/users/dashboard/data");
  recipeData = await response.json();
  recipeData.forEach(item => {
    const category = item.category;
    const subCategory = item.subcategory;
    const name = item.recipe_name

    if (!counts[category]) {
      counts[category] = {
        total: 0,
        subcategories: {}
      };
    }

    if (!counts[category].subcategories[subCategory]) {
      counts[category].subcategories[subCategory] = {
        total: 0,
        names: {}
      };
    }

    if (!counts[category].subcategories[subCategory].names[name]) {
      counts[category].subcategories[subCategory].names[name] = 0;
    }

    counts[category].total++;
    counts[category].subcategories[subCategory].total++;
    counts[category].subcategories[subCategory].names[name]++;
  })

  const result = Object.entries(counts).map(([category, value]) => [category, value.total]);
  renderBarChart(result);
})
