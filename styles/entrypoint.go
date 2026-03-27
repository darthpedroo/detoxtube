package styles

type EntryPoint struct {
	ListItemStyle ListItemStyle
	TitleStyle    TitleStyle
	Footer        FooterStyle
	Terminal      TerminalStyle
}

func NewEntryPoint() EntryPoint {

	return EntryPoint{
		ListItemStyle: *NewListItemStyle(),
		TitleStyle:    *NewTitleStyle(),
		Footer:        *NewFooterStyle(),
		Terminal:      *NewTerminalStyle(),
	}
}
