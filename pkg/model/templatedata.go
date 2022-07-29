package model

//Template data holds data sents from handlers to templates
type TemplateData struct {
	StringMap    map[string]string      //for strings
	IntMap       map[string]int         //for integers
	FloatMap     map[string]float32     //for float values
	Data         map[string]interface{} // for variadic data type
	CSRFToken    string                 //SECURITY TOKEN: cross sight request forgery token
	FlashMessage string                 //messages to send to the user
	Warning      string
	Error        string
}
