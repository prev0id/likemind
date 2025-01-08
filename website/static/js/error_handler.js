document.addEventListener("htmx:responseError", function (event) {
  const status = event.detail.xhr.status;

  if (status === 401) {
    window.location.href = "/signin";
  } else {
    const responseMessage =
      event.detail.xhr.responseText ||
      `Error ${status}: An unexpected error occurred`;

    showErrorMessage(responseMessage);
  }
});

function showErrorMessage(message) {
  let container = document.getElementById("notification-container");
  if (!container) {
    container = document.createElement("div");
    container.id = "notification-container";
    document.body.appendChild(container);
  }

  const errorDiv = document.createElement("div");
  errorDiv.setAttribute("role", "alert");
  errorDiv.setAttribute("tabindex", "-1");
  errorDiv.className =
    "max-w-xs bg-white border border-black rounded-xl shadow-lg";

  errorDiv.innerHTML = `
    <div class="flex p-4">
      <div class="shrink-0">
        <svg class="shrink-0 size-4 text-orange mt-0.5" xmlns="http://www.w3.org/2000/svg" fill="currentColor" width="16" height="16" viewBox="0 0 16 16">
          <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zM8 4a.905.905 0 0 0-.9.995l.35 3.507a.552.552 0 0 0 1.1 0l.35-3.507A.905.905 0 0 0 8 4zm.002 6a1 1 0 1 0 0 2 1 1 0 0 0 0-2z"></path>
        </svg>
      </div>
      <div class="ms-3">
        <p class="text-sm text-black">
          ${message}
        </p>
      </div>
    </div>
  `;

  container.appendChild(errorDiv);

  setTimeout(() => {
    errorDiv.remove();
  }, 10000);
}
