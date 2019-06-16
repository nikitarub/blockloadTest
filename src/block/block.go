package block

// Position – позиция блока на странице
type Position struct {
	Line   int
	Column int
}

// Data - данные блока
type Data string

// BlockOld - единица хранения
type BlockOld struct {
	ID       int
	Data     Data
	Position Position
}

// Block - кдиница хранения с flat - структурой
type Block struct {
	ID         uint
	DocumentID uint
	Data       string
	Line       uint
	Column     uint
}
