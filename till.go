package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// Definování přepínačů
	timeFlag := flag.String("t", "", "Target time in HH:MM format")
	commandFlag := flag.String("c", "", "Command to execute after the timer ends")

	// Parsování přepínačů
	flag.Parse()

	// Kontrola, zda je čas nastaven
	if *timeFlag == "" {
		fmt.Println("Error: Target time (-t) is required.")
		flag.Usage()
		os.Exit(1)
	}

	// Nastavení časové zóny
	loc, err := time.LoadLocation("Europe/Prague")
	if err != nil {
		fmt.Println("Error loading location:", err)
		os.Exit(1)
	}

	// Získání aktuálního času
	currentTime := time.Now().In(loc)

	// Parsování cílového času
	targetTime, err := time.Parse("15:04", *timeFlag)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		os.Exit(1)
	}

	localTime := currentTime.In(loc)

	// Spojení cílového času s aktuálním dnem
	targetDateTime := time.Date(localTime.Year(), localTime.Month(), localTime.Day(), targetTime.Hour(), targetTime.Minute(), 0, 0, localTime.Location())
	duration := targetDateTime.Sub(currentTime)

	// Pokud je cílový čas dříve než aktuální, přidej 24 hodin
	if duration < 0 {
		duration += 24 * time.Hour
	}

	fmt.Printf("Sleeping until %s, that is in %v.\n", targetDateTime.Format("3:04 PM"), duration)

	// Počet snímků za sekundu
	const fps = 50
	const frameDuration = time.Second / fps

	// Odpočet
	for {
		currentTime = time.Now().In(loc)
		remaining := targetDateTime.Sub(currentTime)

		// Pokud je zbývající čas <= 0, ukonči smyčku
		if remaining <= 0 {
			break
		}

		// Vymazání obrazovky (ANSI escape sekvence)
		fmt.Print("\033[H\033[2J")

		// Zobrazení zbývajícího času a příkazu
		fmt.Printf("Time remaining: %02d:%02d:%02d.%03d\n",
			int(remaining.Hours()), int(remaining.Minutes())%60, int(remaining.Seconds())%60, int(remaining.Milliseconds())%1000)
		if *commandFlag != "" {
			fmt.Printf("Command to execute: %s\n", *commandFlag)
		}

		// Vyprázdnění bufferu, aby se výstup ihned zobrazil
		os.Stdout.Sync()

		// Pauza na dobu trvání jednoho snímku
		time.Sleep(frameDuration)
	}

	// Po dokončení odpočtu
	fmt.Printf("Target time reached %s!\n", targetDateTime.Format("3:04 PM"))

	// Spuštění příkazu, pokud byl zadán
	if *commandFlag != "" {
		fmt.Printf("Executing command: %s\n", *commandFlag)
		// Spuštění příkazu
		cmdParts := strings.Split(*commandFlag, " ")
		if len(cmdParts) > 0 {
			err := executeCommand(cmdParts[0], cmdParts[1:]...)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		}
	}

	fmt.Println("Unblocking...")
}

// Funkce pro spuštění příkazu
func executeCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
