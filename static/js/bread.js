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
            document.getElementById('tangzhong-view').style.display = 'block';
            document.getElementById('leavener-view').style.display = 'block';

        }
        if (selection === 'flour-weight') {
            document.getElementById('flour-weight').style.display = 'flex';
            document.getElementById('flour-view').style.display = 'flex';
            document.getElementById('flour').required = true;
            document.getElementById('hydration-view').style.display = 'flex';
            document.getElementById('hydration').required = true;
            document.getElementById('fat-view').style.display = 'flex';
            document.getElementById('sugar-view').style.display = 'flex';

        }
        if (selection === 'total-weight') {
            document.getElementById('total-weight').style.display = 'flex';
            document.getElementById('dough-weight-view').style.display = 'flex';
            document.getElementById('dough-weight').required = true;
            document.getElementById('hydration-view').style.display = 'flex';
            document.getElementById('hydration').required = true;
            document.getElementById('fat-view').style.display = 'flex';
            document.getElementById('sugar-view').style.display = 'flex';
            
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
            document.getElementById('hydration-view').style.display = 'flex';
            document.getElementById('hydration').required = true;
            document.getElementById('fat-view').style.display = 'flex';
            document.getElementById('sugar-view').style.display = 'flex';
        }

        leavenerChoice.addEventListener('change', function() {
            const leavenerSelected = this.value;
            if (leavenerSelected === 'Choose Leavener') {
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
            }
            if (tangzhonSelect === 'No') {
                document.getElementById('tangzhong-select').style.display = 'none';
            }
        })
    })
})