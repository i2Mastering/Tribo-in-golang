# Tribonacci Sequence Calculator in Go

## Overview
This project implements a **Tribonacci Sequence Calculator** in Go. The program calculates the first `n` numbers in the Tribonacci sequence using the `math/big` package to handle very large numbers. The Tribonacci sequence starts with three ones, and each subsequent number is the sum of the previous three numbers.

## Features
- Uses Go's `math/big` package to handle extremely large numbers.
- Provides robust input type handling similar to Python's error messages.
- Demonstrates interface type assertions in Go.
- Computes and outputs the Tribonacci sequence as a slice of big integers.

## How It Works
1. **Tribonacci Function**: 
   - Generates a Tribonacci sequence based on the user-defined `n` length.
   - Uses big integers (`big.Int`) to prevent overflow and ensure correctness for large values.
   - Returns a slice of `*big.Int` representing the sequence.

2. **Main Function**:
   - Accepts an `interface{}` input to simulate Python-like type checking.
   - Provides clear error messages for string and float inputs.
   - Computes and prints the Tribonacci sequence when provided a valid integer input.

## Usage
### Prerequisites
- Go installed (version 1.13 or later recommended).

### Run the Program
1. Save the code in a file named `main.go`.
2. Open your terminal or command prompt.
3. Navigate to the directory containing `main.go`.
4. Run the program using:
   ```bash
   go run main.go
   ```

### Example Output
```
[1 1 1 3 5]
```

For non-integer inputs:
```
TypeError: 'float' object cannot be interpreted as an integer
```
Or:
```
TypeError: '<=' not supported between instances of 'str' and 'int'
```

## Notes
- The program initializes the input as `interface{}` and currently defaults to an integer (`5`).
- Modify the `input` variable in the `main` function to test different cases (string, float, int).
- The Tribonacci sequence here starts from three initial ones: `[1, 1, 1]`.

## License
This project is provided for educational purposes and is open-source.

---

### Author
**Ike Iloegbu**  
Developed in Visual Studio Code
