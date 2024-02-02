package Helper

func AddNumber(num1, num2 int) int {

	return num1 + num2
}

// Adds numbers from range
func AddNumber2(nums ...int) int {

	var total int = 0

	for _, value := range nums {
		total += value
	}

	return total
}

// working with methods
type CustomInt int

func (variable CustomInt) IsEven() bool {
	return variable%2 == 0
}

// working with interface and type assertions
type Printer interface { //Defining our interface
	Print() string
}

// defining structs of different type
type Vehicle struct {
	TypeOfVehicle   string
	CountryOfOrigin string
}

type Book struct {
	AuthorName string
	Lang       string
}

// Implementing interface for each of the struct types
func (book Book) Print() string {
	return "Book details: author name " + book.AuthorName + " lang " + book.Lang
}

func (vehicle Vehicle) Print() string {
	return "Vehicle details: type: " + vehicle.TypeOfVehicle + " origin " + vehicle.CountryOfOrigin
}
