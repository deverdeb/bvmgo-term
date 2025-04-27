package term

type Color [3]uint8

var (
	White     Color = [3]uint8{255, 255, 255}
	GrayLight Color = [3]uint8{170, 170, 170}
	Gray      Color = [3]uint8{140, 140, 140}
	GrayDark  Color = [3]uint8{100, 100, 100}
	Black     Color = [3]uint8{0, 0, 0}
	Red       Color = [3]uint8{220, 80, 80}
	Green     Color = [3]uint8{120, 220, 70}
	Blue      Color = [3]uint8{60, 110, 220}
	Yellow    Color = [3]uint8{250, 220, 60}

	ColorOk    = Green
	ColorInfo  = Blue
	ColorWarn  = Yellow
	ColorError = Red
)

func ColorRGB(red, green, blue uint8) Color {
	return [3]uint8{red, green, blue}
}

func (color *Color) Red() uint8 {
	return color[0]
}

func (color *Color) SetRed(red uint8) {
	color[0] = red
}

func (color *Color) Green() uint8 {
	return color[1]
}

func (color *Color) SetGreen(green uint8) {
	color[1] = green
}

func (color *Color) Blue() uint8 {
	return color[2]
}

func (color *Color) SetBlue(blue uint8) {
	color[2] = blue
}

func (color *Color) Average(otherColor Color) Color {
	return ColorRGB(
		uint8((uint16(color.Red())+uint16(otherColor.Red()))/2),
		uint8((uint16(color.Green())+uint16(otherColor.Green()))/2),
		uint8((uint16(color.Blue())+uint16(otherColor.Blue()))/2),
	)
}

func ColorAdd(color1 *Color, color2 *Color) *Color {
	if color1 == nil {
		return color2
	}
	if color2 == nil {
		return color1
	}
	newColor := color1.Add(*color2)
	return &newColor
}

func (color *Color) Add(otherColor Color) Color {
	return ColorRGB(
		uint8(max(uint16(color.Red())+uint16(otherColor.Red()), 255)),
		uint8(max(uint16(color.Green())+uint16(otherColor.Green()), 255)),
		uint8(max(uint16(color.Blue())+uint16(otherColor.Blue()), 255)),
	)
}

func ColorAverage(color1 *Color, color2 *Color) *Color {
	if color1 == nil {
		return color2
	}
	if color2 == nil {
		return color1
	}
	newColor := color1.Average(*color2)
	return &newColor
}
