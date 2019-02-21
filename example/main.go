package main

import (
	"bufio"
	"fmt"
	"os"
	"systray"
)

const appIconResID = 7

func main() {

	tray, err := systray.New()
	if err != nil {
		panic(err)
	}

	err = tray.Show(appIconResID, "Test systray")
	if err != nil {
		panic(err)
	}

	// Append more menu items and use tray.AppendSeparator() to separate them.
	tray.AppendMenu("Close", func() {
		fmt.Println("Close")
		os.Exit(0)
	})

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Println("Input text, press <enter>, wait for notification to close:")
			fmt.Print(">> ")
			data, _, _ := reader.ReadLine()
			line := string(data)
			if len(line) == 0 {
				break
			}
			err := tray.ShowMessage("Got Text", line)
			if err != nil {
				fmt.Println(err)
			}
		}

		err = tray.Stop()
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	}()

	err = tray.Run()
	if err != nil {
		fmt.Println(err)
	}
}
