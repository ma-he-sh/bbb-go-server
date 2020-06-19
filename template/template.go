package templ

import (
    "html/template"
    "net/http"
)

var Templates *template.Template

type PageStruct struct {
    pageTitle, bodyCont, bkScript, frScript, bkStyle, frStyle string
    isAdmin bool
}

func init() {
    Templates = template.Must( template.ParseGlob("include/*.html") )
}

func Render(w http.ResponseWriter, templName string, data map[string]interface{}) {
    templFile := templName + ".html"
    if err := Templates.ExecuteTemplate(w, templFile, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// PageObject
func PageObj( title string ) PageStruct {
    e := PageStruct{
        pageTitle: title,
    }

    e.bkScript = "<script src='/public/script/backend.js'></script>"
    e.frScript = "<script src='/public/script/frontend.js'></script>"

    e.bkStyle = "<link rel='stylesheet' type='text/css' href='/public/styles/backend.css'></link>"
    e.frStyle = "<link rel='stylesheet' type='text/css' href='/public/styles/frontend.css'></link>"

    return e
}

// SetBody
func (e *PageStruct) SetBody( content string  ) {
    e.bodyCont = content
}

// IsAdmin
func (e *PageStruct) IsAdmin( admin bool ) {
    e.isAdmin = admin
}

func (e *PageStruct) GetBKScripts() string {
    return e.bkScript
}

func (e *PageStruct) GetFRScripts() string {
    return e.frScript
}

func (e *PageStruct) GetBKStyles() string {
    return e.bkStyle
}

func (e *PageStruct) GetFRStyles() string {
    return e.frStyle
}

// GetTemplPayload
func (e *PageStruct) GetTemplPayload() map[string]interface{}  {

    scripts := e.GetFRScripts()
    styles := e.GetFRStyles()
    if e.isAdmin {
        scripts = e.GetBKScripts()
        styles = e.GetBKStyles()
    }

    data := map[string]interface{}{
        "page_title" : e.pageTitle,
        "body_cont"  : e.StrToHTML( e.bodyCont ),
        "header_script": e.StrToHTML( styles ),
        "footer_script": e.StrToHTML(scripts),
    }
    return data
}

// String to html
func (e *PageStruct) StrToHTML( input string ) template.HTML {
    return template.HTML(input)
}
