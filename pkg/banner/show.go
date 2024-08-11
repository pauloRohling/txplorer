package banner

import (
	"fmt"
	"os"
)

const FileName = "banner.txt"

func Show() {
	bannerPath := fmt.Sprintf("./%s", FileName)
	file, err := os.ReadFile(bannerPath)
	if err != nil {
		return
	}
	fmt.Println(string(file))
}
