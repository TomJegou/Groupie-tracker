
function toggleDropdown() {
  document.getElementsByClassName("dropdown").classList.toggle("show");
}
//open/close when dropdown button is clicked
document.getElementById("btn").addEventListener("click",function(){ toggleDropdown()});
// Close dropdown when dom element is clicked
document.documentElement.addEventListener("click", function () {
  if (document.getElementsByClassName("dropdown").classList.contains("show")) {
    toggleDropdown();
  }
});

