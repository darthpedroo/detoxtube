package types

type VideoSort string

const (
	Alphabetically VideoSort = "Alphabetically"
	Date VideoSort = "Date"
)

func (s VideoSort) String() string {
	switch s {
	case Alphabetically:
		return "Alphabetically"
	case Date:
		return "Date"
	}

	return "unknown"
}

type Order string

const (
	Ascendant Order = "Ascendant"
	Descending Order = "Descending"
)

func (s Order) String() string {
	switch s {
	case Ascendant:
		return "Ascendant"
	case Descending:
		return "Descending"
	}
	return "unknown"
}