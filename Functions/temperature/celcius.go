package temperature

type Celcius float32

func (c *Celcius) ToFarenheite() Farenheite {
	return Farenheite((*c * 9.0 / 5) + 32)
}
