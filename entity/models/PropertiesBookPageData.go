package models

type PropertiesBookPageData struct {
	AuthorID      uint
	AuthorName    string
	AuthorAccount string
	IgnoreAuthor  bool
	PageText      string
}

func (prop PropertiesBookPageData) Clone() PropertiesBookPageData {
	return PropertiesBookPageData{
		AuthorID:      prop.AuthorID,
		AuthorName:    prop.AuthorName,
		AuthorAccount: prop.AuthorAccount,
		IgnoreAuthor:  prop.IgnoreAuthor,
		PageText:      prop.PageText,
	}
}
