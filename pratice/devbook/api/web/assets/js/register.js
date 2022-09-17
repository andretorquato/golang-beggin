$("#register-form").on('submit', createUser);

function createUser(event){
	event.preventDefault();
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
	}). done(() => alert('User Created')).fail(() => alert('error on create new user'));
}