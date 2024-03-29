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

func ClientLoginForm(event map[string]interface{}) string {

	var showEmail string = ""
	if event["toggle_email"].(bool) {
		showEmail = Input(InputStruct{Type: "email", Id: "str_email", Name: "str_email", Label: "Email", Value: "", Atts: nil})
	}

	return frontend(`
	<div class='app--floating-modal'>
		<form action='' method='POST'>
			<div class='modal--header'>
				JOIN ` + event["name"].(string) + `
			</div>
			<div class='modal--body'>
			<input type='hidden' name='str_eventid' id='str_eventid' value='` + event["eventid"].(string) + `'>
			` + showEmail + Input(InputStruct{Type: "text", Id: "str_name", Name: "str_name", Label: "Name", Value: "", Atts: nil}) +
		Input(InputStruct{Type: "password", Id: "str_token", Name: "str_token", Label: "Access Token", Value: "", Atts: nil}) + `
			</div>
			<div class='modal--footer'>
				<div class='button--wrapper button--group'>
				` + FRButton(ButtonStruct{Type: "submit", Varient: "primary", Name: "btn_join", Id: "btn_join", Label: "Join In", Active: true, Atts: nil}) +
		LinkButton(LinkStruct{Varient: "secondary", Id: "btn_home", Name: "btn_home", Label: "Home", Link: "/", Atts: nil}) + `
				</div>
			</div>
		</form>
	</div>
	`)
}
