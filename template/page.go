package templ

// backend template
func backend(content string) string {
	return `<div class='app--backend-wrapper'>
	` + content + `
	</div>`
}

// frontend template
func frontend(content string) string {
	return `<div class='app--frontend-wrapper'>
	` + content + `
	</div>`
}
