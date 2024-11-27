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
  let container = document.getElementById("notification-container");
  if (!container) {
    container = document.createElement("div");
    container.id = "notification-container";
    document.body.appendChild(container);
  }

  const errorDiv = document.createElement("div");
  errorDiv.className = "error-popup";
  errorDiv.innerHTML = htmlContent;

  container.appendChild(errorDiv);

  setTimeout(() => {
    errorDiv.remove();
  }, 10000);
}
