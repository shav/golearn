package temperature

type Farenheite float32

func (f *Farenheite) ToCelcius() Celcius {
	return Celcius(((*f - 32) * 5.0) / 9)
}
