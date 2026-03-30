package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/RobertWesner/ClamClient/core/client"
)

func main() {
	c := client.Connect("localhost:25565", "RobertWesner2")

	go func() {
		select {
		case <-c.Done():
			return
		case <-c.Ready():
		}

		scanner := bufio.NewScanner(os.Stdin)

		for {
			select {
			case <-c.Done():
				return
			default:
			}

			fmt.Print("> ")
			if !scanner.Scan() {
				if err := scanner.Err(); err != nil {
					log.Print(err)
				}
				c.Close()

				return
			}

			line := scanner.Text()
			if line == "/exit" {
				c.Close()

				return
			}

			c.Chat(line)
		}
	}()

	go func() {
		for {
			select {
			case message := <-c.Events().Chat():
				fmt.Println(message)
			case <-c.Events().SetSlot():
				break
			case <-c.Events().WindowItems():
				break
			case <-c.Events().Transaction():
				break
			case reason := <-c.Events().Disconnect():
				fmt.Println(reason)
			}
		}
	}()

	<-c.Done()
	if err := c.Err(); err != nil {
		slog.Error("fatal client error", "error", err)
	}
}
