package talks

import (
	"bufio"
	"fmt"
	"go-chat/utils"
	"os"
	"strings"
)

func FirstStep() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello😃 I'm Go! What's your name?")

	for {
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		exists, err := utils.IsVariableExists("variable/name.go", strings.Title(name))

		if err != nil {
			fmt.Println("Error checking variable existence:", err)
			continue
		}
		if exists {
			fmt.Println("Oh!We meet again😄 How are you?(good/bad)")
			greeting, _ := reader.ReadString('\n')
			greeting = strings.TrimSpace(greeting)
			switch greeting {
			case "Good", "good":
				fmt.Println("That's good!")
				break
			case "Bad", "bad":
				fmt.Println("I'm sorry to hear that😥")
				break
			default:
				break
			}
		} else {
			f, err := os.OpenFile("variable/name.go", os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				fmt.Println(err)
				continue
			}
			defer f.Close()
			fmt.Fprintf(f, "var %s = \"%s\"\n", strings.Title(name), strings.Title(name))
		}

		if name != "" {
			fmt.Printf("Hello! %s\n", name)
			break
		}
		fmt.Println("Sorry? Please try again😥")
	}
}
