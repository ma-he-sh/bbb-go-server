package templ

type ButtonStruct struct {
	Type, Varient, Name, Id, Label string
	Active                         bool
	Atts                           map[string]string
}

func FRButton(params ButtonStruct) string {
	varient := "normal"
	if params.Varient != "" {
		varient = params.Varient
	}

	active := ""
	if !params.Active {
		active = "disabled"
	}

	atts := ""
	if params.Atts == nil {
		atts = ""
	} else {
		for k, v := range params.Atts {
			atts += "data-" + k + "='" + v + "' "
		}
	}

	return `<button type='` + params.Type + `' id='` + params.Id + `' name='` + params.Name + `' ` + atts + ` class='button--type button--` + varient + `' ` + active + `>` + params.Label + `</button>`
}

type LinkStruct struct {
	Varient, Name, Id, Label, Link string
	Atts                           map[string]string
}

func LinkButton(params LinkStruct) string {
	varient := "normal"
	if params.Varient != "" {
		varient = params.Varient
	}

	atts := ""
	if params.Atts == nil {
		atts = ""
	} else {
		for k, v := range params.Atts {
			atts += "data-" + k + "='" + v + "' "
		}
	}

	return `<a href='` + params.Link + `' id='` + params.Id + `' name='` + params.Name + `' ` + atts + ` class='button--type button--` + varient + `'>` + params.Label + `</a>`
}

type InputStruct struct {
	Type, Id, Name, Label, Value string
	Atts                         map[string]string
}

func Input(params InputStruct) string {
	atts := ""
	if params.Atts == nil {
		atts = ""
	} else {
		for k, v := range params.Atts {
			atts += "data-" + k + "='" + v + "' "
		}
	}

	return `<div class='input--wrapper'>
		<label class='input--label'>` + params.Label + `<label>
		<input class='input--type-text' ` + atts + ` id='` + params.Id + `' type='` + params.Type + `' value='` + params.Value + `' name='` + params.Name + `' />
	</div>`
}
