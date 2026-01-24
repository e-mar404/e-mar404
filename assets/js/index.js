const details = document.querySelector("details");

document.addEventListener("DOMContentLoaded", () => {
  const smallScreen = window.matchMedia("(max-width: 768px)").matches;
  if (details && smallScreen) {
    details.open = false;
  }
});

document.addEventListener("resize", () => {
  const isSmallScreen = window.innerWidth <= 768
  details.open = !isSmallScreen;
  console.log("resize window")
});
