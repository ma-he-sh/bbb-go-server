package templ

type ErrorType struct {
	Message string
	show    bool
}

func BKButton() string {
	return ""
}

func FRButton() string {
	return ""
}

type InputStruct struct {
	Type, Id, Name, Lable, Value string
	Error                        ErrorType
}

func Input() string {
	return ""
}
