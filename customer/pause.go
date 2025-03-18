package customer

import (
	"bufio"
	"fmt"
	"os"
)

func Pause() {
	fmt.Println("按任意键继续...")
	bufio.NewReader(os.Stdin).ReadByte()
}
