package defaults

import "fmt"

//nolint:unused
type example struct {
	Name string `default:"Tom"`
}

//nolint:unused,deadcode
func useDefault() {
	var e example
	err := Apply(&e)
	if err != nil {
		panic(err)
	}
	fmt.Println(e)
}
