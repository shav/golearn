package mc // = my collections

type BitSet uint8
type Flag uint8

func (b *BitSet) Set(flag Flag)      { *b = *b | BitSet(flag) }
func (b *BitSet) Clear(flag Flag)    { *b = *b &^ BitSet(flag) }
func (b *BitSet) Toggle(flag Flag)   { *b = *b ^ BitSet(flag) }
func (b *BitSet) Has(flag Flag) bool { return *b&BitSet(flag) != 0 }
