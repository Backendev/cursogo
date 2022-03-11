package main

import "fmt"

func main() {
	//declaracion constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Pi 1-", pi, "\nPi 2-", pi2)

	//Declaracion variables
	//variables no declaradas
	base := 12
	//Con declaracion y asignacion
	var altura int = 14
	//Solo declaracion
	var area int = base * altura

	fmt.Println("Base: ", base, "\nAltura: ", altura, "\nArea: ", area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	//Operadore incremental y decremental
	x := 10
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

	//Paquete fmt
	hello := "Hello"
	world := "World"
	cant := 2
	ccc := 4
	fmt.Println(hello, world)
	fmt.Println(hello, world)
	fmt.Printf("%s numero %d\n", world, cant)
	cant++
	fmt.Printf("%s numero %d ", world, cant)
	cant++
	fmt.Printf("%s numero %d \n", world, cant)
	//format cualquier tipo de dato
	fmt.Printf("%v numero %v \n", world, cant)
	cant++
	message := fmt.Sprintf("%v numero %v", world, cant)
	fmt.Println(message)
	//tipo de dato
	fmt.Printf("Tipo: %T \n", hello)

	// General

	// %v	the value in a default format
	// when printing structs, the plus flag (%+v) adds field names
	// %#v	a Go-syntax representation of the value
	// %T	a Go-syntax representation of the type of the value
	// %%	a literal percent sign; consumes no value

	// Bool

	// %t	the word true or false

	// Integer:

	// %b	base 2
	// %c	the character represented by the corresponding Unicode code point
	// %d	base 10
	// %o	base 8
	// %O	base 8 with 0o prefix
	// %q	a single-quoted character literal safely escaped with Go syntax.
	// %x	base 16, with lower-case letters for a-f
	// %X	base 16, with upper-case letters for A-F
	// %U	Unicode format: U+1234; same as "U+%04X"

	// Floating-point and complex constituents:

	// %b	decimalless scientific notation with exponent a power of two,
	// 	in the manner of strconv.FormatFloat with the 'b' format,
	// 	e.g. -123456p-78
	// %e	scientific notation, e.g. -1.234456e+78
	// %E	scientific notation, e.g. -1.234456E+78
	// %f	decimal point but no exponent, e.g. 123.456
	// %F	synonym for %f
	// %g	%e for large exponents, %f otherwise. Precision is discussed below.
	// %G	%E for large exponents, %F otherwise
	// %x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
	// %X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20
	// String and slice of bytes (treated equivalently with these verbs):

	// %s	the uninterpreted bytes of the string or slice
	// %q	a double-quoted string safely escaped with Go syntax
	// %x	base 16, lower-case, two characters per byte
	// %X	base 16, upper-case, two characters per byte
	// Slice:

	// %p	address of 0th element in base 16 notation, with leading 0x
	// Pointer:

	// %p	base 16 notation, with leading 0x
	// The %b, %d, %o, %x and %X verbs also work with pointers,
	// formatting the value exactly as if it were an integer.
	// The default format for %v is:

	// bool:                    %t
	// int, int8 etc.:          %d
	// uint, uint8 etc.:        %d, %#x if printed with %#v
	// float32, complex64, etc: %g
	// string:                  %s
	// chan:                    %p
	// pointer:                 %p
	// For compound objects, the elements are printed using these rules, recursively, laid out like this:

	// struct:             {field0 field1 ...}
	// array, slice:       [elem0 elem1 ...]
	// maps:               map[key1:value1 key2:value2 ...]
	// pointer to above:   &{}, &[], &map[]
	// Width is specified by an optional decimal number immediately preceding the verb. If absent, the width is whatever is necessary to represent the value. Precision is specified after the (optional) width by a period followed by a decimal number. If no period is present, a default precision is used. A period with no following number specifies a precision of zero. Examples:

	// %f     default width, default precision
	// %9f    width 9, default precision
	// %.2f   default width, precision 2
	// %9.2f  width 9, precision 2
	// %9.f   width 9, precision 0
	// Other Flags
	// +	always print a sign for numeric values;
	// 	guarantee ASCII-only output for %q (%+q)
	// -	pad with spaces on the right rather than the left (left-justify the field)
	// #	alternate format: add leading 0b for binary (%#b), 0 for octal (%#o),
	// 	0x or 0X for hex (%#x or %#X); suppress 0x for %p (%#p);
	// 	for %q, print a raw (backquoted) string if strconv.CanBackquote
	// 	returns true;
	// 	always print a decimal point for %e, %E, %f, %F, %g and %G;
	// 	do not remove trailing zeros for %g and %G;
	// 	write e.g. U+0078 'x' if the character is printable for %U (%#U).
	// ' '	(space) leave a space for elided sign in numbers (% d);
	// 	put spaces between bytes printing strings or slices in hex (% x, % X)
	// 0	pad with leading zeros rather than spaces;
	// 	for numbers, this moves the padding after the sign
	miFuncion("Hola !!!", "Que tal")
	value := miFuncionReturn(3)
	fmt.Println(value)
	value1, value2 := miFuncionReturn2(3)
	fmt.Println(value1, value2)
	//descartar un valor de return
	value1, _ = miFuncionReturn2(5)
	fmt.Println(value1)
	ciclos()
	conditional()
	defer_keyword()
	arrays()
	isPalindromo("ala")

}

//Funciones
func miFuncion(message string, message2 string) {
	fmt.Println(message, " ", message2)
}
func miFuncionReturn(a int) int {
	return a * 2
}
func miFuncionReturn2(a int) (b, c int) {
	return a * 2, a / 2
}
func ciclos() {
	//For condicional
	for i := 0; i < 10; i++ {
		fmt.Printf("-%v ", i)
	}
	fmt.Println()

	//For while
	counter := 10
	for counter < 20 {
		fmt.Printf("-%v ", counter)
		counter++
	}
	fmt.Println()

	//For forever

	// counterf := 0
	// for {
	// 	fmt.Println(counterf)
	// 	counterf++
	// }

	// counterf := 0
	// for {
	// 	fmt.Println(counterf)
	// 	counterf++
	// if counterf == 10{
	// 	break
	// }
	// }

	//For range
	listan := []int{1, 2, 3, 4, 5, 6, 7}
	for i, num := range listan {
		fmt.Printf("-%v --%v |", num, i)
	}
	fmt.Println()
	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 4 {
			fmt.Println("Es 4")
			continue
		}

		// Break
		if i == 6 {
			fmt.Println("Break")
			break
		}
	}
}
func conditional() {
	//operadores logicos
	// And &&
	// Or ||

	//Not
	//myBool :=  true
	// fmt.Println(!myBool) // Esto retornará false
	if isPair(3) {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}
	if isValidUser("Alpha5", "MyPassword") {
		fmt.Println("Credentials are valid")
	} else {
		fmt.Println("Credentials aren't valid")
	}
	//switch
	var dia int8 = 4
	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	default:
		fmt.Println("otro dia")
	}
}
func isPair(num int) bool {
	return num%2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}

func defer_keyword() {
	defer fmt.Println("Ultimo")
	fmt.Println("Primero")
	fmt.Println("Segundo")
}
func arrays() {
	//array inmutables
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, "Len ", len(array), "Cap ", cap(array))
	//slices mutables
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, "Len ", len(slice), "Cap ", cap(slice))
	//metodos slice 1 param inclusivo 2 param exclusivo
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	slice = append(slice, 7)
	fmt.Println(slice, "Len ", len(slice), "Cap ", cap(slice))
	new_list := []int{8, 9, 10}
	// ... dese,paqueta lista
	slice = append(slice, new_list...)
	fmt.Println(slice, "Len ", len(slice), "Cap ", cap(slice))
	//Maps Llave valor items separados por espacios
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20
	fmt.Println(m)
	//Recorrer map por defecto reocorre de manera aleatoria
	for k, v := range m {
		fmt.Println(k, "-", v)
	}
	fmt.Println(m["Jose"])
	//Buscar un valor , se añade una variable que recibe un bool
	value, ok := m["Jose"]
	value2 := m["Joseh"]
	fmt.Println(value2, "-")
	value, ok = m["Josep"]
	fmt.Println(value, "-", ok)

}
func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if textReverse == text {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("No Es Palindromo")
	}
}
