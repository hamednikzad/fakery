package color

type Color string

const (
	None Color = "\033[0m"
	Red  Color = "\033[0;31m"
	RED  Color = "\x1B[31m"
	NRM  Color = "\x1B[0m"
	GRN  Color = "\x1B[32m"
	YEL  Color = "\x1B[33m"
	BLU  Color = "\x1B[34m"
	MAG  Color = "\x1B[35m"
	CYN  Color = "\x1B[36m"
	WHT  Color = "\x1B[37m"
)
