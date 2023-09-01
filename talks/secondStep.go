package talks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SecondStep() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`Let's talk to me! (Yes/No)`)
loop:
	for {
		startTalk, _ := reader.ReadString('\n')
		startTalk = strings.TrimSpace(startTalk)
		switch startTalk {
		case "Yes", "YES", "yes":
			fmt.Println("Nice!")
			break loop
		case "No", "NO", "no":
			fmt.Println("Bye👋")
			os.Exit(0)
		default:
			fmt.Println("Sorry? Please try again😥")
		}
	}
}
