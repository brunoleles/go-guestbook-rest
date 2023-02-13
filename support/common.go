package support

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"

	"github.com/joho/godotenv"
)

func FindDotenvFile(filename string, filepath string) string {
	var filepathname = path.Clean(path.Join(filepath, filename))

	if _, err := os.Stat(filepathname); err == nil {
		return filepathname
	}

	if filepath == "" || filepath == "/" {
		log.Println(".env file not found")
		return ""
	}

	return FindDotenvFile(filename, path.Join(filepath, ".."))
}

func InitEnv() {
	p, _ := BasePath("")
	dotenvPathname := FindDotenvFile(".env", p)

	if dotenvPathname != "" {
		if err := godotenv.Load(dotenvPathname); err != nil {
			log.Print(err)
			log.Printf("Error acourred while loading .env file in \"%s\"\n", dotenvPathname)
		}
	}
}

func BasePath(value string) (string, error) {
	var base string
	var err error

	if base, err = os.Executable(); err != nil {
		return "", err
	}

	base = filepath.Dir(base)

	if value != "" {
		base = path.Join(base, value)
	}

	return base, nil
}

func CopyInto(a, b interface{}) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b).Elem()

	at := av.Type()
	for i := 0; i < at.NumField(); i++ {
		name := at.Field(i).Name

		bf := bv.FieldByName(name)
		if bf.IsValid() {
			bf.Set(av.Field(i))
		}
	}
}

func PanicOnError(e error) {
	if e == nil {
		return
	}
	log.Fatalf("Fatal error occurred: \"%s\"", e.Error())
}
