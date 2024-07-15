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
### Scenario 1

Insert command after till, for safety reasons you can use `&&` instead. In most cases `;` will do fine.

   ```sh
	till 13:30; kill -9 `pidof firefox`
   ```

### Scenario 2

 Run the whole in subshell (backgrounded).

   ```sh
	(till 13:31; beep) &
   ```
### Scenario 3

 Add time first, then decide what to run later. Not every terminal will allow you to do this, it should work as well in most cases.

   ```sh
	till 13:30 <enter>
	command <enter>
   ```


## Timezone misdetection bug

On some systems, namely `Termux` and some virtual machines there can be an issue with setting correct timezone. Any suggested fix didn't resolve this among all systems tested, it has definitely something to do with TZ and system locale, output incoherrent.

Edit the source code to pass your preferred TZ string. 
