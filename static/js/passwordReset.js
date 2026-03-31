let passwordInput = document.getElementById("password")
passwordInput.addEventListener("input", () => {
  const password = passwordInput.value;
  passwordStrength(password);
})

document.getElementById("passwordReset").addEventListener("submit", async function(e) {
  e.preventDefault();

 const toHide = ["password-error", "messageContainer"];
 toHide.forEach(id => {
   hide(id);
 })  

  const firstPassword = document.getElementById("password").value;
  const checkPass = document.getElementById("check_password").value;

  if (firstPassword !== checkPass) {
      show("password-error");
      changeText("password-error", "Passwords do not match");
      return
  }
  if (firstPassword.length < 8 ) {
      show("password-error");
      changeText("password-error", "Password should be at least 8 characters long.");
      return
  }

  const formData = new FormData(e.target);
  const results = await fetch("/passwordResetSubmit", {
    method: "POST",
    headers: {
      "Accept": "application/json"
    },
    body: formData,
  });

  message = document.getElementById("message");
  const data = await results.json(); 
  message.innerText = data.message;
  console.log("This is the message: ", data.message)
  if(!data.ok) {
    message.classList.add("warning");
    message.classList.remove("success");
  } else {
    message.classList.add("success");
    message.classList.remove("warning");
  }
  show("messageContainer");
  setTimeout(() => {
    window.location.href = "/login";
  }, 5000);
})
