function refreshOnModalClose(popoverID) {
  const popover = document.getElementById(popoverID);
  popover.addEventListener("beforetoggle", (event) => {
    if (event.newState === "closed") {
      window.location.reload();
    }
  });
}
