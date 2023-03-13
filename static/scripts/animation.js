
// const dropdowns = document.querySelectorAll('.dropdown-toggle');

// dropdowns.forEach(dropdown => {
//   if (dropdown.classList.contains('click-dropdown')) {
//     dropdown.addEventListener('click', (event) => {
//       event.preventDefault();
//       console.log('click')
//       const parent = dropdown.parentElement;
//       const dropdownMenu = parent.querySelector('.dropdown-menu');
//       const isOpen = dropdownMenu.classList.contains('dropdown-active');

//       closeDropdown();

//       if (!isOpen) {
//         parent.classList.add('dropdown-open');
//         dropdownMenu.classList.add('dropdown-active');
//       }
//     });
//   }
// });

// function closeDropdown() { 
//   document.querySelectorAll('.dropdown-open').forEach((openDropdown) => {
//     openDropdown.classList.remove('dropdown-open');
//   });

//   document.querySelectorAll('.dropdown-active').forEach((activeDropdown) => {
//     activeDropdown.classList.remove('dropdown-active');
//   });
// }




// window.addEventListener('click', (event) => {
//   if (!event.target.closest('.dropdown-container')) {
//     closeDropdown();
//     console.log('close')
//   }
// });