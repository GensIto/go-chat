package talks

import (
	"bufio"
	"fmt"
	"go-chat/utils"
	"os"
	"strings"
)

func ThirdStep() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`do you like food? (Yes/No)`)

	for {
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		switch answer {
		case "Yes", "YES", "yes":
			fmt.Println("Nice!")
			for {
				fmt.Println("what kind of food do you like?")
				answer, _ := reader.ReadString('\n')
				answer = strings.TrimSpace(answer)

				exists, err := utils.IsVariableExists("variable/food.go", answer)
				if err != nil {
					fmt.Println("Error checking variable existence:", err)
					continue
				}
				if exists {
					fmt.Println("That's already I know! Please other food!")
					continue
				}

				err = utils.IsCheckCharacter(answer)
				if err != nil {
					fmt.Println(err)
					continue
				}

				f, err := os.OpenFile("variable/food.go", os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					fmt.Println(err)
					continue
				}
				defer f.Close()
				fmt.Fprintf(f, "var %s = \"%s\"\n", strings.Title(answer), strings.Title(answer))
				fmt.Println("That's good! thank you for your answer!")
				return
			}

		case "No", "NO", "no":
			fmt.Println("Tell us what you like?")
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)
			fmt.Printf("Do you like %s", answer)
			fmt.Println("That's good! thank you for your answer!")
			fmt.Println("Bye👋")
			os.Exit(0)
		default:
			fmt.Println("Sorry? Please try again😥")
		}
	}
}
