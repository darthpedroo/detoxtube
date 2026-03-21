package styles

type EntryPoint struct{
	ListItemStyle ListItemStyle
	TitleStyle  TitleStyle
}

func NewEntryPoint() EntryPoint{
	
	return EntryPoint{
		ListItemStyle: *NewListItemStyle(),
		TitleStyle: *NewTitleStyle(),
	}
}