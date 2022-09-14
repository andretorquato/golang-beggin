$("#register-form").on(createUser);

function createUser(event){
	event.preventDefault();
	console.log("registe")
	if($("#password").val() != $("#password-confirm").val()){
		alert("fill the same password");
		return
	}

	$.ajax({
		url: "/users",
		method: "POST",
		data: {
			name: $("#name").val(),
			email: $("#email").val(),
			nick: $("#nick").val(),
			password: $("#password").val()
		}
	});
}