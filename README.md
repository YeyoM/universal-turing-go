# Universal Turing Machine

The Universal Turing Machine is a Go program that simulates a Turing machine based on a description provided in a file and an input string. The machine processes the input string according to the rules defined in the file, simulating the computation step by step.

### Features

- Customizable Machine Description: Specify the Turing machine description via a file.
- Input String Handling: Pass the input string via a flag or prompt.
- Verbose Mode: Display the tape and current state at each step.


### Installation

1. Clone the repository:

```sh
Copy code
git clone https://github.com/yourusername/universal-turing-machine.git
cd universal-turing-machine
```

2. Build the project:

```sh
go build -o turing
```

#### Usage

The program accepts several command-line flags to control its behavior:

```sh
./turing [options]
```

####Options

- --help, -h: Display the help message.
- --verbose, -v: Display the tape and the current state of the machine at each step.
- --step, -s: Display the tape and the current state of the machine at each step, and wait for the user to press enter to continue.
- --input, -i <input_string>: Specify the input string for the machine. If not specified, the user will be prompted to enter the input string.
- --file, -f <file_path>: Specify the file containing the Turing machine description. If not specified, the user will be prompted to enter the file name.

#### Example

```sh
./turing --file example.tm --input 0101 --verbose
```

#### Turing Machine Description File Syntax

The Turing machine description file should follow this syntax:

- Each line contains a tuple in the form: <current state> <current symbol> <new symbol> <direction> <new state>
- <current state>: The current state of the machine.
- <current symbol>: The symbol currently being read.
- <new symbol>: The symbol to write.
- <direction>: The direction to move the tape ('l' for left, 'r' for right, '*' for no move).
- <new state>: The next state of the machine.
- Use _ to represent a blank (space) symbol.
- Use ; to add comments.
- States and symbols are case-sensitive.

#### Special Symbols

- * can be used as a wildcard in <current state> or <current symbol> to match any state or symbol.
- * in <new state> or <new symbol> means no change.

#### Example File

```
0 0 1 r 1
1 1 0 l 0
0 * * * halt
; This is a comment
```

#### Implementation Details

The project consists of several key functions:

- read_file(file_name string): Reads the file and returns a list of tuples.
- check_line_syntax(line string): Checks if the syntax of a line is correct.
- initialize_tape(input_string string): Initializes the tape with the input string.
- run_machine(tape *Tape): Runs the Turing machine on the tape using the specified transitions.
- display_tape(tape *Tape): Displays the current tape.
- display_state(state string): Displays the current state of the machine.
- main(): The main function that orchestrates the reading of inputs and running the machine.

#### Contributing

Feel free to open issues or submit pull requests if you find bugs or want to add features.

#### License

This project is licensed under the GNU GENERAL PUBLIC License. See the LICENSE file for details.
