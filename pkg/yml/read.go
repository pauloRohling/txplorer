package yml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func Read(filename string, v any) error {
	file, err := os.Open(filename)

	if err != nil {
		return fmt.Errorf("error reading %s", filename)
	}

	defer func(file *os.File) { _ = file.Close() }(file)

	if err = yaml.NewDecoder(file).Decode(v); err != nil {
		return fmt.Errorf("error decoding %s", filename)
	}

	return nil
}
