package types

type VideoSort string

const (
	Alphabetically VideoSort = "Alphabetically"
)

func (s VideoSort) String() string {
	switch s {
	case Alphabetically:
		return "Alphabetically"
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