let dropdownBtn = document.getElementsByClassName("btn");
let dropdownMenu = document.getElementsByClassName("dropdown");
let toggleDropdown = function () {
  dropdownMenu.classList.toggle("show");
};
//open/close when dropdown button is clicked
dropdownBtn.addEventListener("click", function () {
  toggleDropdown();
});
// Close dropdown when dom element is clicked
document.documentElement.addEventListener("click", function () {
  if (dropdownMenu.classList.contains("show")) {
    toggleDropdown();
  }
});