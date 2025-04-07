document.addEventListener("DOMContentLoaded", function () {
  let currentYear = new Date().getFullYear();
  document.getElementById(
    "copyright-footer"
  ).innerHTML = `&#169 ${currentYear}`;
});
