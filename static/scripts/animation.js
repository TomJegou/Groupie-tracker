
const dropdowns = document.querySelectorAll('.dropdown-toggle');
dropdowns.forEach(dropdown => {
  if (dropdown.classList.contains('click-dropdown')) {
    dropdown.addEventListener('click', (event) => {
      event.preventDefault();
      const parent = dropdown.parentElement;
      const dropdownMenu = parent.querySelector('.dropdown-menu');
      const isOpen = dropdownMenu.classList.contains('dropdown-active');

      closeDropdown();

      if (!isOpen) {
        parent.classList.add('dropdown-open');
        dropdownMenu.classList.add('dropdown-active');
      }
    });
  }
});

dropdowns.forEach(dropdown => {
  console.log('start')
  if (dropdown.classList.contains('dropdown')) {
    dropdown.addEventListener('click', (event) => {
      const parent = dropdown.parentElement;
      const dropdownMenu = parent.querySelector('.dropdown-menu');
      parent.classList.add('dropdown-open');
      dropdownMenu.classList.add('dropdown-active');
    });
    dropdown.addEventListener('click', (event) => {
      const parent = dropdown.parentElement;
      const dropdownMenu = parent.querySelector('.dropdown-menu');
      dropdownMenu.addEventListener('click', closeDropdown);
    });
  }
});

function closeDropdown() { 
  document.querySelectorAll('.dropdown-open').forEach((openDropdown) => {
    openDropdown.classList.remove('dropdown-open');
  });

  document.querySelectorAll('.dropdown-active').forEach((activeDropdown) => {
    activeDropdown.classList.remove('dropdown-active');
  });
}
window.addEventListener('click', (event) => {
  console.log('close')
  if (!event.target.closest('.dropdown-container')) {
    closeDropdown();
  }
});