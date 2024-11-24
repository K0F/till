# Timer with Command Executor: till

![Build Status](https://github.com/K0F/till/actions/workflows/go.yml/badge.svg)


This program functions as a simple countdown timer that executes a specified command once the target time is reached. The program supports options for setting the target time and an optional command to run after the countdown ends.

## Installation

If you don't have Go installed, download it from the [official Go website](https://golang.org/dl/).

1. Clone or download this repository.
2. In the terminal, navigate to the folder containing the `till.go` file and run the following command to compile the program:

```bash
   go build till.go
```
This will generate an executable file named till. Copy it to your path to install.

## Usage

The program is run with the following options:

    -t to set the target time in the HH:MM format.
    -c to specify the command to be executed after the timer finishes.

## Example 1: Simple countdown

To set only the target time and start the countdown:

```
./till -t 14:30
```

This command will wait until 14:30 and then exit.

## Example 2: Set a command to be executed

To run a command after the countdown reaches the target time:

	./till -t 14:30 -c "echo 'Hello, world!'"

This command will wait until 14:30, then execute echo 'Hello, world!' and display the output.
Countdown Output

During the countdown, the remaining time and the command to be executed will be displayed on the screen:


	Time remaining: 00:15:34.123
	Command to execute: echo 'Hello, world!'


Once the target time is reached, the program will execute the command and show its output.
Example output after reaching the target time

If the target time is reached and the command is executed:

	Target time reached 14:30!
	Executing command: echo 'Hello, world!'
	Hello, world!
	Unblocking...

## Dependencies

    Go 1.18+ (for building)

## License

This project is licensed under the GPL 3.0 Licence


## Timezone misdetection bug

On some systems, namely `Termux` and some virtual machines there can be an issue with setting correct timezone. Any suggested fix didn't resolve this among all systems tested, it has definitely something to do with TZ and system locale, output incoherrent.

Edit the source code to pass your preferred TZ string. 
