let form = document.querySelecter('form');

form.addEventListener('submit', (e) => {
  e.preventDefault();
  return false;
});

function login() {
  document.getElementById("login").click();
}

function register() {
  document.getElementById("register").click();
}