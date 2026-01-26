document.getElementById("calc-form").addEventListener("focusin", (e) => {
    inputSelected = e.target.id;

    let insertHtml = ""; 
    let bakersPercentage = `
        <p>Provide the baker's percentage for this ingredient. Use whole numbers. For example if you want 50% just type 50.</p>
        <p>
            <b>Baker's Percentage</b>: Expresses a ratio in percentage of each ingredient's weight to the total flour weight.
            For example: If you have 1000 grams of flour and 500 grams of water, then the bakers percentage for the flour is 100% and for the water it is 50%.
            The baker's percentage for flour is always 100%.
        </p>
    `;

    switch(inputSelected) {
        case "calculator-bread":
            insertHtml = `
                <p>Choose how you want to build this recipe</p> 
                <p><b>Flour Weight</b>: Recipe is calculated based on the amount of flour you want to use.</p>
                <p><b>Total Weight</b>: Recipe is calculated on the final dough weight. This is great for things like hamburger buns. You know the total amount of dough that goes into each bun. So multiply that by how many buns you want.</p>
                <p><b>Pan Dimension</b>: Recipe is calculated so the bread fits that pan.</p>
            `;
            break;
        case "measurement":
            insertHtml = `
                <p>Are you measuring the container in inches or centimeters?</p>
            `;
            break;
        case "shape":
            insertHtml = `
                <p>Does the container have 4 sides or is it more cylinder?</p>
                <p>Crazier shapes not implemented yet!</p> 
            `;
            break;
        case "height":
        case "width":
        case "depth":
            insertHtml = `
                <p>What is the length of this side?</p>
            `;
            break;
        case "diameter":
            insertHtml = `
                <p>Measurement across the entire circle.</p> 
            `;
            break;
        case "flour":
            insertHtml = `
                <p>What is the total weight, in grams, of flour that you want to use?</p>
            `;
            break;
        case "dough-weight":
            insertHtml = `
                <p>What is the desired final weight that you want for the dough?.</p>
            `;
            break;
        case "hydration":
        case "fat":
        case "sugar":
        case "egg":
            let flour = document.getElementById("flour").value;
            let wholeEgg = 56 / flour * 100;
            insertHtml = bakersPercentage + `
                    <b>Note:</b> The average large egg weighs 56 grams. To use one large egg this percenate would need to be ${Math.ceil(wholeEgg)}.
                `;
            break;
        case "salt":
            insertHtml = bakersPercentage + `
                 <p><b>Note:</b> I generally do 2%.</p> 
            `;
            break;
        case "leavener-choice":
            insertHtml = `
                <p>How are you planning on leavening this bread?</p>  
            `;
            break;
        case "sourdough":
            insertHtml = bakersPercentage + `
                <p>How much sourdough starter are you using. The general recommendations I have seen are from 10 to 30 percent.  I generally go with the happy medium
                of 20%</p>
            `; 
            break;
        case "yeast":
            insertHtml = bakersPercentage + `
                <p>How much yeast do you want to use? General recommendations is about 1%.</p>
            `;
            break;
        case "tangzhong-choice":
            insertHtml = `
                <p>Tangzhong or Yudane are techniques to create softer bread by gelantinizing some of the flour.
                While the two methods have some differences, this calculator can help figure out the ratio and then you can decide the method you want to follow.</p>
            `;
            break;
        case "tangzhong-percentage":
            insertHtml = bakersPercentage + `
                 <p><b>Note:</b> This will not add any additional flour or hydration to your recipe. It will take from the flour and hydration calculated above.
                 If too much is stolen, it will let you know.</p>
            `;
            break;
        case "tangzhong-ratio":
            insertHtml = `
                <p>Here you provide the water to flour ratio. What multiple should be applied.</p>
                <p><b>Example:</b> Inserting a 1 says you want a 1:1 ratio. Use the same amount of flour and water.
                Inserting a 4 says you want a 4:1 ratio.  Use 4 times the amount of water as flour.</p> 
            `;
    }
    document.getElementById("info-box").innerHTML = insertHtml;

})
