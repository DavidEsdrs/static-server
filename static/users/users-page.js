const addUserButton = document.querySelector(".add-user")

addUserButton.addEventListener("click", ev => {
  redirect("/usuarios/criacao")
})