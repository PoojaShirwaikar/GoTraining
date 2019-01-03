package DB

import "fmt"
import "DataAccess/File"

func ReadSQLDB() {
	fmt.Println("DB package")
	File.ReadCSVFile()
}
