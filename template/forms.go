package templ

func AdminLoginForm(message string) string {
	return frontend(`
	<div class='app--floating-modal'>
		<form action='/admin/login' method='POST'>
			<div class='modal--header'>
				LOGIN
			</div>
			<div class='modal--body'>
			` + Input(InputStruct{Type: "email", Id: "str_email", Name: "str_email", Label: "Email", Value: "", Atts: nil}) +
		Input(InputStruct{Type: "password", Id: "str_passw", Name: "str_passw", Label: "Password", Value: "", Atts: nil}) + `
			</div>
			<div class='modal--footer'>
				<div class='button--wrapper button--group'>
				` + FRButton(ButtonStruct{Type: "submit", Varient: "primary", Name: "btn_login", Id: "btn_login", Label: "Log In", Active: true, Atts: nil}) +
		LinkButton(LinkStruct{Varient: "secondary", Id: "btn_home", Name: "btn_home", Label: "Home", Link: "/", Atts: nil}) + `
				</div>
			</div>
		</form>
	</div>
	`)
}

func ClientLoginForm() string {
	return frontend(``)
}
