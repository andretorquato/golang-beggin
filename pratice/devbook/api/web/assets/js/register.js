$("#register-form").on(createUser);

function createUser(event){
	event.preventDefault();

	if($("#password").val() != $("#password-confirm").val()){
		alert("fill the same password");
		return
	}

	$.ajax({
		url: "users",
		method: "POST",
		data: {
			name: $("#name").val()
		}
	});
}