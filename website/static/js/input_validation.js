function validateConfirm() {
  password = document.querySelector("input[name=password]");
  confirmPassword = document.querySelector("input[name=confirm-password]");
  validationMessage = document.getElementById("confirm-validation-message");

  if (password.value === confirmPassword.value) {
    confirmPassword.setCustomValidity("");
    validationMessage.className = "text-black";
    validationMessage.innerHTML = "match";
  } else {
    confirmPassword.setCustomValidity("Passwords should match");
    validationMessage.className = "text-red";
    validationMessage.innerHTML = "don't match";
  }
}
