document.addEventListener("htmx:responseError", function (event) {
  const status = event.detail.xhr.status;

  if (status === 401) {
    window.location.href = "/login";
  } else {
    const responseHTML = event.detail.xhr.responseText;
    showErrorMessage(
      responseHTML || `Error ${status}: An unexpected error occurred`,
    );
  }
});

function showErrorMessage(htmlContent) {
  let popupContainer = document.getElementById("notification-container");
  if (!popupContainer) {
    popupContainer = document.createElement("div");
    popupContainer.id = "popup-container";
    document.body.appendChild(popupContainer);
  }

  const errorDiv = document.createElement("div");
  errorDiv.className = "error-popup";
  errorDiv.innerHTML = htmlContent;

  popupContainer.appendChild(errorDiv);

  setTimeout(() => {
    errorDiv.remove();
  }, 15000);
}
