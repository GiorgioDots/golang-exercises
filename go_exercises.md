
Go Practice Exercises
=====================

BASIC (5 Exercises)
-------------------
1. FizzBuzz
   - Print numbers from 1 to 100.
   - For multiples of 3, print "Fizz"; for multiples of 5, print "Buzz"; for multiples of both, print "FizzBuzz".

2. Reverse a Slice
   - Write a function that reverses a slice of integers in place.

3. Word Frequency Counter
   - Input: a string.
   - Output: a map showing how many times each word appears.

4. Prime Number Checker
   - Function: IsPrime(n int) bool
   - Return true if n is a prime number; false otherwise.

5. Simple CLI Calculator
   - Accept command-line flags: --op (add/sub/mul/div), --a, --b.
   - Perform the operation and print the result.


INTERMEDIATE (5 Exercises)
--------------------------
1. Bank Account Struct
   - Define a BankAccount struct with balance, deposit, and withdraw methods.
   - Prevent overdrawing (balance < 0).

2. File Line Counter
   - Read a file and count how many lines it contains.
   - Return an error if the file can’t be read.

3. Custom Logger Interface
   - Create a Logger interface with Log(msg string).
   - Implement ConsoleLogger and FileLogger that satisfy this interface.

4. Goroutine Worker Pool
   - Create a fixed number of workers.
   - Each worker processes jobs from a channel (e.g., squares integers).

5. HTTP Client Wrapper
   - Wrap Go’s http.Client and log request method, URL, and response status.


ADVANCED (5 Exercises)
----------------------
1. REST API for a Todo List
   - Build an API with endpoints for creating, listing, updating, and deleting todos.
   - Use Gorilla Mux or Chi.
   - Store data in memory or SQLite.

2. Concurrency-Safe Cache
   - Implement a Cache struct with Get, Set, and Delete methods.
   - Use sync.RWMutex to make it safe for concurrent access.

3. Generics: Map, Filter, Reduce
   - Write generic functions for Map, Filter, and Reduce.
   - Example: Map([]int{1,2,3}, func(x int) int { return x * 2 }) -> []int{2,4,6}

4. gRPC Math Server
   - Define a .proto file for basic math operations.
   - Implement a gRPC server and client in Go.

5. Testable CLI App
   - Build a CLI todo manager.
   - Organize the code for testability (e.g., using interfaces).
   - Write unit tests for the core logic.
