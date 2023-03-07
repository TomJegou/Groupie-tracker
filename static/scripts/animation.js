function dropdown() {
  var container = document.getElementById("dropdown-container");
  if (container.style.display == "block") {
    container.style.display = "none";
  } else {
    container.style.display = "block";
  }
window.onclick = function (event) {
  if (!event.target.matches('.dropdown-btn')) {
      document.getElementById('dropdown-container').style.display = "none";
  }
}  
} 