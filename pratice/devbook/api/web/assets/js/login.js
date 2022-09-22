$("#login").on("submit", login);

function login(event) {
  event.preventDefault();
  $.ajax({
    url: "/login",
    method: "POST",
    data: {
      email: $("#email").val(),
      password: $("#password").val(),
    },
  }).done(() => window.location.href = "/home")
    .fail((err) => {
      console.log(err);
      alert("error on login");
    });
}
