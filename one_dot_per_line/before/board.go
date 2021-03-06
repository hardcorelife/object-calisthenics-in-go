package before

import "bytes"

type board struct {
    locations []*location
}

func NewBoard() *board {
    locations := []*location{
        NewLocation(NewPiece("London")),
        NewLocation(NewPiece("New York")),
        NewLocation(NewPiece("Dubai")),
    }
    return &board{
        locations: locations,
    }
}

func (b *board) squares() []*location {
    return b.locations
}

func (b *board) BoardRepresentation() string {
    var buffer = &bytes.Buffer{}
    for _, l := range b.squares() {
        buffer.WriteString(l.current.representation[0:1])
    }
    return buffer.String()
}