package maps

const (
	HOME_PAGE = "HOME"
	CONTACTS = "CONTACTS"
)

var url map[string]string = map[string]string{
	HOME_PAGE: "/",
	CONTACTS: "/contacts",
}

func GetUrl(page string) string {
	return url[page]
}