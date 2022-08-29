package tokenize

type Position struct {
	LineNo int // the line number
	Lat    int // starts AT (Line)
	Wat    int // starts AT (Whole)
}

func NewPosition(ln, lat, wat int) *Position {
	return &Position{
		LineNo: ln,
		Lat:    lat,
		Wat:    wat,
	}
}
