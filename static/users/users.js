// url
const url = "http://127.0.0.1:3030/usuarios/criar"

// queries
const form = document.querySelector("form")
const input = document.querySelector("input")

form.addEventListener("submit", ev => {
  ev.preventDefault()

  fetch(url, {
    method: "post",
    body: JSON.stringify({
      username: input.value
    })
  })
  .then(res => {
    cleanInput()
    redirect("/usuarios")
  })
})

function cleanInput() {
  input.value = ""
}