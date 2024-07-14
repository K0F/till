# Sleep Until Time Program

![Build Status](https://github.com/K0F/till/actions/workflows/go.yml/badge.svg)

This Go program allows you to sleep until a specified time. It takes a time in the format `HH:MM` (24-hour format) as a command-line argument and blocks execution until that exact time is reached.

## Usage

   ```sh
   git clone https://github.com/K0F/till.git
   cd till
   go mod tidy
   go build
   (copy or link program to path)
   ```

   ```sh
	till 13:30; ./someScript.sh
   ```
