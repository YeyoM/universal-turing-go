package main

/*
The syntax for the input turing machine is the following...

The program will run a Turing machine that is read from a file passed by the user,
and will simulate the machine on the input string passed by the user.

There are going to be some specific flags that the user can pass to the program to
change the behavior of the program.

Possible Flags:

  --help, -h: Display the help message.
  --verbose, -v: Display the tape and the current state of the machine at each step.
  --step, -s: Display the tape and the current state of the machine at each step, and wait for the user to press enter to continue.
  --input, -i: Specify the input string for the machine. If not specified, the user will be prompted to enter the input string.
  --file, -f: Specify the file containing the Turing machine. If not specified, the user will be prompted to enter the file name.

The input file should contain a turing machine with the correct syntax. The syntax is the following...

Syntax:

  Each line should contain one tuple of the form '<current state> <current symbol> <new symbol> <direction> <new state>'.
  You can use any number or word for <current state> and <new state>, eg. 10, a, state1. State labels are case-sensitive.
  You can use almost any character for <current symbol> and <new symbol>, or '_' to represent blank (space). Symbols are case-sensitive.
  You can't use ';', '*', '_' or whitespace as symbols.
  <direction> should be 'l', 'r' or '*', denoting 'move left', 'move right' or 'do not move', respectively.
  Anything after a ';' is a comment and is ignored.
  The machine halts when it reaches any state starting with 'halt', eg. halt, halt-accept.

Also:

  '*' can be used as a wildcard in <current symbol> or <current state> to match any character or state.
  '*' can be used in <new symbol> or <new state> to mean 'no change'.


Ideas on how we can make this work

1. Read the file (and check if the syntax is correct, if not, throw an error) and store the tuples in a list.
2. Create a dictionary with the tuples as keys and the values as the new symbol, direction and new state.
3. Create a doubly linked list to represent the tape. 
4. Initialize the tape with the input string. (in case the syntax is wrong, throw an error)

Functions needed:

1. read_file(file_name): Reads the file and returns a list of tuples.
2. check_syntax(tuples): Checks if the syntax of the tuples is correct.
3. create_dict(tuples): Creates a dictionary with the tuples as keys and the values as the new symbol, direction and new state.

4. initialize_tape(input_string): Initializes the tape with the input string.
5. run_machine(tape, dictionary): Runs the machine on the tape using the dictionary.

6. display_tape(tape): Displays the tape.
7. display_state(state): Displays the current state of the machine.

8. main(): Main function that calls all the other functions.
*/

import (
  "fmt"
  "os"
  "flag"
  "bufio"
)

// TAPE 
// [ BOX ] <-> [ BOX ] <-> [ BOX ]

type Box struct {
  symbol string
  prev *Box
  next *Box
}

type Tape struct {
  start *Box
  end *Box
}

func read_file(file_name string, tuples []string, allowed_symbols map[string]string, allowed_states map[string]string, allowed_transitions map[string]string) {
  // Read the file and return a list of tuples.
  file, err := os.Open(file_name)
  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  fmt.Println("Reading, and checking the syntax of the file...")
  for scanner.Scan() {
    line := scanner.Text()
    check_line_syntax(line, tuples, allowed_symbols, allowed_states, allowed_transitions)
  }
  fmt.Println("File read successfully.\n\n")

  if err := scanner.Err(); err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }
}

func check_line_syntax(line string, tuples []string, allowed_symbols map[string]string, allowed_states map[string]string, allowed_transitions map[string]string) {
  if len(line) == 0 {
    return 
  }

  // Check if the line is a comment.
  if line[0] == ';' {
    return
  }

  // Divide the line into tokens. Each token is found between two spaces.
  tokens := []string{}
  token := ""
  for i := 0; i < len(line); i++ {
    if line[i] == ' ' {
      tokens = append(tokens, token)
      token = ""
    } else if line[i] == ';' {
      // break the for loop
      break
    } else {
      token += string(line[i])
    }

    if i == len(line) - 1 {
      tokens = append(tokens, token)
    }
  }

  if len(tokens) != 5 {
    fmt.Println("Error: Invalid syntax.")
    os.Exit(1)
  }

  _, ok := allowed_symbols[tokens[1]]
  if ok != true {
    allowed_symbols[tokens[1]] = tokens[1]
  }

  _, ok1 := allowed_symbols[tokens[2]]
  if ok1 != true {
    allowed_symbols[tokens[2]] = tokens[2]
  }

  var completeTransition string
  var partialTransition string

  completeTransition = tokens[0] + " "  + tokens[1] + " " + tokens[2] + " " + tokens[3] + " " + tokens[4]
  partialTransition = tokens[0] + " " + tokens[1] 

  //fmt.Println("Complete Transition:", completeTransition)
  //fmt.Println("Partial Transition:", partialTransition)

  _, ok2 := allowed_transitions[partialTransition]
  if ok2 == true {
    fmt.Println("Fatal error, there can not be 2 transtitions that are the same")
    os.Exit(0)
  } else {
    allowed_transitions[partialTransition] = completeTransition
  }

  _, ok3 := allowed_states[tokens[0]]
  if ok3 != true {
    allowed_states[tokens[0]] = tokens[0]
  }

  _, ok4 := allowed_states[tokens[4]]
  if ok4 != true {
    allowed_states[tokens[4]] = tokens[4]
  }
}

func newTape() *Tape {
  return &Tape{}
}

func display_tape(t *Tape) {
  fmt.Print("Tape: ")
  box := t.start
  for box != nil {
    fmt.Print(box.symbol)
    box = box.next
  }
  fmt.Println()
}

func (t *Tape) addBox(symbol string) {
  newBox := &Box{symbol: symbol}
  if t.start == nil {
    t.start = newBox
    t.end = newBox
  } else {
    t.end.next = newBox
    newBox.prev = t.end
    t.end = newBox
  }
}

// given the tape and a pointer to the current box, move the tape to the left or right and return the new box 
func (t *Tape) moveBox(current_box *Box, direction string) *Box {
  if direction == "l" {
    if current_box.prev == nil {
      newBox := &Box{symbol: "_"}
      current_box.prev = newBox
      newBox.next = current_box
      t.start = newBox
      return newBox
    } else {
      return current_box.prev
    }
  }
    
  if direction == "r" {
      
    if current_box.next == nil {
      newBox := &Box{symbol: "_"}
      current_box.next = newBox
      newBox.prev = current_box
      t.end = newBox
      return newBox
    } else {
      return current_box.next
    }
  }

  return current_box
}

func (t *Tape) moveLeft() {
  if t.start.prev == nil {
    newBox := &Box{symbol: "_"}
    t.start.prev = newBox
    newBox.next = t.start
    t.start = newBox
  } else {
    t.start = t.start.prev
  }
}

func (t *Tape) moveRight() {
  if t.end.next == nil {
    newBox := &Box{symbol: "_"}
    t.end.next = newBox
    newBox.prev = t.end
    t.end = newBox
  } else {
    t.end = t.end.next
  }
}

func (t *Tape) getCurrentSymbol() string {
  return t.start.symbol
}

func (t *Tape) setCurrentSymbol(symbol string) {
  t.start.symbol = symbol
}

func initialize_tape(input_string string, allowed_symbols map[string]string, tape *Tape) {
  // Initialize the tape with the input string. 
  for i := 0; i < len(input_string); i++ {
    _, ok := allowed_symbols[string(input_string[i])]
    if ok != true {
      fmt.Println("Error: Invalid symbol in the input string.")
      os.Exit(1)
    }
    tape.addBox(string(input_string[i]))
  }
  fmt.Println("Tape initialized successfully.")
}

func run_machine(tape *Tape, allowed_symbols map[string]string, allowed_states map[string]string, allowed_transitions map[string]string, verbose bool) {
  // Run the machine on the tape using the dictionary.
  fmt.Println("Running the machine on the tape...")
  fmt.Println()

  // initialize the current state and current symbol
  current_state := "0"
  current_symbol := tape.getCurrentSymbol()
  current_wildcard := false

  current_box := tape.start

  // loop until the machine halts
  for {
    current_symbol = current_box.symbol

    if verbose {
      display_tape(tape)
      fmt.Println("Current state:", current_state)
      fmt.Println("Current symbol:", current_symbol)
    }

    // check if the current state is in the allowed states 
    _, ok := allowed_states[current_state]
    if ok != true {
      fmt.Println("Error: Invalid state.")
      os.Exit(1)
    }
    
    // check if the current symbol is in the allowed symbols
    _, ok1 := allowed_symbols[current_symbol]
    if ok1 != true {
      fmt.Println("Error: Invalid symbol.")
      os.Exit(1)
    }

    // check if the current transition is in the allowed transitions 
    _, ok2 := allowed_transitions[current_state + " " + current_symbol]
    if ok2 != true {
      // checking for transitions that have a wildcard in the current symbol 
      _, ok3 := allowed_transitions[current_state + " *"]
      if ok3 == true {
        current_wildcard = true
      } else {
        fmt.Println("Halting the machine.")
        fmt.Println("")
        display_tape(tape)
        fmt.Println("Last state:", current_state)
        fmt.Println("Last symbol:", current_symbol)
        os.Exit(0)
      }
    }

    var transition string

    // get the transition 
    if current_wildcard == true {
      transition = allowed_transitions[current_state + " *"]
    } else {
      transition = allowed_transitions[current_state + " " + current_symbol]
    }

    if verbose {
      fmt.Println("Transition:", transition)
    } 

    // get the new symbol, direction and new state 
    // the transition looks like this: "0 1 _ r 1i" -> "curr_state curr_symbol new_symbol direction new_state"

    // divide the string into tokens separated by a space 
    tokens := []string{}
    token := ""
    for i := 0; i < len(transition); i++ {
      if transition[i] == ' ' {
        tokens = append(tokens, string(token))
        token = ""
      } else {
        token += string(transition[i])
      }

      if i == len(transition) - 1 {
        tokens = append(tokens, string(token))
      }
    }

    if current_wildcard == true {
      current_wildcard = false
      if tokens[2] == "*" {
        tokens[2] = current_symbol
      }
    } 

    new_symbol := tokens[2]
    direction := tokens[3]
    new_state := tokens[4]


    if verbose {
      fmt.Println("New symbol:", new_symbol)
      fmt.Println("New Direction:", direction)
      fmt.Println("New state:", new_state)
      fmt.Println()
    }

    // set the current symbol to the new symbol 
    current_box.symbol = new_symbol 

    // move the tape to the left or right 
    if direction == "l" { 
      current_box = tape.moveBox(current_box, direction)
    }

    if direction == "r" {
      current_box = tape.moveBox(current_box, direction)
    }

    // set the current state to the new state 
    current_state = new_state 
  }
}

func help() {
  fmt.Println("Usage: turing [options]")
  fmt.Println("Options:")
  fmt.Println("  --help, -h: Display the help message.")
  fmt.Println("  --verbose, -v: Display the tape and the current state of the machine at each step.")
  fmt.Println("  --step, -s: Display the tape and the current state of the machine at each step, and wait for the user to press enter to continue.")
  fmt.Println("  --input, -i: Specify the input string for the machine. If not specified, the user will be prompted to enter the input string.")
  fmt.Println("  --file, -f: Specify the file containing the Turing machine. If not specified, the user will be prompted to enter the file name.")
}

func main() {

  tuples := []string{}

  // here is a dictionary with the allowed symbols based on the input file 
  allowed_symbols := make(map[string]string)
  allowed_states := make(map[string]string)
  allowed_transitions := make (map[string]string)

  help_flag := flag.Bool("help", false, "Display the help message.")
  verbose_flag := flag.Bool("verbose", false, "Display the tape and the current state of the machine at each step.")
  input_flag := flag.String("input", "", "Specify the input string for the machine.")
  file_flag := flag.String("file", "", "Specify the file containing the Turing machine.")
  flag.Parse()

  if *help_flag {
    help()
    os.Exit(0)
  }

  // If the file name is not passed as a flag, prompt the user to enter the file name.
  if *file_flag == "" {
    fmt.Print("Enter the file name: ")
    var file_name string
    fmt.Scan(&file_name)
    *file_flag = file_name
  }

  // Read the file and return a list of tuples.
  if file_flag != nil && *file_flag != "" {
    read_file(*file_flag, tuples, allowed_symbols, allowed_states, allowed_transitions)
  } else {
    fmt.Println("Error: No file name provided.")
    os.Exit(1)
  }

  // If no input string is passed as a flag, prompt the user to enter the input string.
  if *input_flag == "" {
    fmt.Print("Enter the input string (content of the tape): ")
    var input_string string 
    fmt.Scan(&input_string)
    *input_flag = input_string
  }

  // If the input string is passed as a flag, display the input string.
  fmt.Println("Input string:", *input_flag)

  if *verbose_flag {
    fmt.Println("Verbose mode enabled.")
  }

  // Create a doubly linked list to represent the tape.
  tape := newTape()

  // Initialize the tape with the input string.
  initialize_tape(*input_flag, allowed_symbols, tape)

  // Run the machine on the tape using the dictionary.
  run_machine(tape, allowed_symbols, allowed_states, allowed_transitions, *verbose_flag)
}
