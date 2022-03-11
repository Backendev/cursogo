package mypackage

//crear struct
type carPrivate struct {
	brand string
	year  int
}

//Propiedades publicas con primera letra en mayuscula

//CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

func PrintMessage(text string) {
	println(text)
}

//Print message private
func printMessagePrivate(text string) {
	println(text)
}
