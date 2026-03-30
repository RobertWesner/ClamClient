package main

import (
	"bufio"
	"clamclient/client"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	c := client.Connect("localhost:25565", "RobertWesner")

	go func() {
		for {
			select {
			case <-c.Done():
				return
			case <-c.Ready():
				fmt.Print("$ ")
				scanner := bufio.NewScanner(os.Stdin)

				for scanner.Scan() {
					line := scanner.Text()

					if line == "/exit" {
						c.Close()

						return
					}

					c.Chat(line)
				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	go func() {
		for {
			fmt.Println("---")
			message := <-c.Events().Chat()
			fmt.Println(message)
		}
	}()

	<-c.Done()
	if err := c.Err(); err != nil {
		slog.Error("fatal client error", "error", err)
	}
}
