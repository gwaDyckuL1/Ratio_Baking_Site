document.addEventListener("DOMContentLoaded", function() {
    const breadSelect = this.getElementById("calculator-bread");
    const tangzhongChoice = this.getElementById("tangzhong-choice");
    const leavenerChoice = this.getElementById('leavener-choice');

    breadSelect.addEventListener("change", function() {
        const selection = this.value;

        const breadOptions = document.querySelectorAll(".bread-options")
        breadOptions.forEach(option => {
            option.style.display = 'none';
            document.getElementById('bread-submission').style.display = 'none';
        })

        const ingredients = document.querySelectorAll(".ingredient")
        ingredients.forEach(ingredient => {
            ingredient.style.display = 'none';
            document.getElementById('tangzhong-choice').value = 'No';
            document.getElementById('leavener-choice').value = "Choose Leavener";
            ingredient.querySelector("input").required = false;
        })
        if (selection != '' ){
            document.getElementById('bread-submission').style.display = 'block';
            document.getElementById('egg-view').style.display = 'flex';
            document.getElementById('fat-view').style.display = 'flex';
            document.getElementById('hydration-view').style.display = 'flex';
            document.getElementById('hydration').required = true;
            document.getElementById('leavener-view').style.display = 'block';
            document.getElementById('salt-view').style.display = 'flex';
            document.getElementById('sugar-view').style.display = 'flex';
            document.getElementById('tangzhong-view').style.display = 'block';


        }
        if (selection === 'flour-weight') {
            document.getElementById('flour-weight').style.display = 'flex';
            document.getElementById('flour-view').style.display = 'flex';
            document.getElementById('flour').required = true;
        }
        if (selection === 'total-weight') {
            document.getElementById('total-weight').style.display = 'flex';
            document.getElementById('dough-weight-view').style.display = 'flex';
            document.getElementById('dough-weight').required = true;            
        }
        if (selection === 'pan-dimension') {
            document.getElementById('pan-dimension').style.display = 'block';
            document.getElementById('shape').addEventListener('change', function() {
                const shape = this.value
                document.getElementById('square').style.display = 'none'
                document.getElementById('round').style.display = 'none'

                if (shape === 'square') {
                    document.getElementById('square').style.display = 'block'
                }
                if (shape === 'round') {
                    document.getElementById('round').style.display = 'block'
                }
            })
        }

        leavenerChoice.addEventListener('change', function() {
            const leavenerSelected = this.value;
            if (leavenerSelected === '') {
                document.getElementById('sourdough-view').style.display='none';
                document.getElementById('yeast-view').style.display='none';
                
            }
            if (leavenerSelected === 'Sourdough') {
                document.getElementById('sourdough-view').style.display='flex';
                document.getElementById('yeast-view').style.display='none';
            }
            if (leavenerSelected == 'Yeast') {
                document.getElementById('yeast-view').style.display='flex';
                document.getElementById('sourdough-view').style.display='none';
            }
        })

        tangzhongChoice.addEventListener('change', function() {
            const tangzhonSelect = this.value;
            if (tangzhonSelect === 'Yes') {
                document.getElementById('tangzhong-select').style.display = 'block';
                const tangzhongPercentageInput = document.getElementById("tangzhong-percentage");
                const tangzhongRatioInput = document.getElementById("tangzhong-ratio")
                const hydrationInput = document.getElementById("hydration")
                tangzhongPercentageInput.addEventListener('input', function() {
                    if (tangzhongRatioInput.value != "") {
                        checkHydrationLevels(parseFloat(hydrationInput.value), parseFloat(tangzhongPercentageInput.value), parseFloat(tangzhongRatioInput.value));
                    }
                })
                tangzhongRatioInput.addEventListener('input', function () {
                    if (tangzhongPercentageInput.value != "") {
                        checkHydrationLevels(parseFloat(hydrationInput.value), parseFloat(tangzhongPercentageInput.value), parseFloat(tangzhongRatioInput.value));
                    }
                })
                hydrationInput.addEventListener("input", function() {
                    if (tangzhongRatioInput.value != "" && tangzhongPercentageInput.value != "") {
                        checkHydrationLevels(parseFloat(hydrationInput.value), parseFloat(tangzhongPercentageInput.value), parseFloat(tangzhongRatioInput.value));
                    }
                })
            }
            if (tangzhonSelect === 'No') {
                document.getElementById('tangzhong-select').style.display = 'none';
            }
        })
    })
})

function checkHydrationLevels(hydration, tangzhongPercentage, tangzhongRatio) {
    document.getElementById("all-tangzhong").style.display='none';
    document.getElementById("hydration-needed").style.display='none';
    const tangzhongHydration = (tangzhongPercentage / (tangzhongRatio + 1)) * tangzhongRatio;
    if (hydration === tangzhongHydration)  {
        document.getElementById("all-tangzhong").style.display='flex';
    }
    if (hydration < tangzhongHydration) {
        document.getElementById("hydration-needed").style.display='flex';
    }
}