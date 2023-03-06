
// function toggleDropdown() {
//   document.getElementById("dropdown").classList.toggle("show");
// }
// //open/close when dropdown button is clicked
// document.getElementById("btn").addEventListener("click",function(){ toggleDropdown()});
// Close dropdown when dom element is clicked
// document.documentElement.addEventListener("click", function () {
//   if (document.getElementsByClassName("dropdown").classList.classList("show")) {
//     toggleDropdown();
//   }

// 
//   // Get the button, and when the user clicks on it, execute myFunction
//   document.getElementById("hamburger-icon").onclick = function() {myFunction()};

//   /* myFunction toggles between adding and removing the show class, which is used to hide and show the dropdown content */
//   function myFunction() {
//     document.getElementById("dropdown-content").classList.toggle("show");
//   }
// }
window.onload = () => {
  document.getElementById("btn").addEventListener("click", toggleDropdown)
  // document.getElementById("btn").onclick = function() {toggleDropdown()};
  function toggleDropdown() {
    document.getElementById("dropdown").classList.toggle("show");
  }
  // document.documentElement.addEventListener("click", function () {
  //   if (document.getElementsByClassName("dropdown").classList.classList("show")) {
  // toggleDropdown();
      }
// })}