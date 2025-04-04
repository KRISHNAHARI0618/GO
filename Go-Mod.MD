## Go is Compiled or Interpreted  Language `Why Go`
- Go is A compiled and statically typed
- Go Tool Can run a file without precompiling
- Compiled executables are os specific
- It is invented By Google
- Concurrency
- Efficient compilation
- efficient execution
- ease of programming

### Use Cases:
- Go Is Statically Typed Language

## Some areas I think I should focus on:

- Concurrency and Goroutines
- Interfaces and Structs
- Error Handling Best Practices
- Memory Management in Go
- Performance Optimization

## Key Components
- Functions --> `main`
- Organizing Code with packages
- Single package or multiple packages organize the code
- Standard Library
- why we need to package name is main
- main is special package 
- Go Mod --> module can have multiple package

### Go Concurrency:
- Running multiple tasks at a time unlike parallelism 
- It is managaed by go run time environment
- Parallalism is running multiple task at a timeon multiple cpus

### Concurrency Features
- Go Routines: 
  - Lightweight threads managed by go runtime environment
  - Created using the go keyword
- Go Channels:
  - Used to communicate beteewn two go routines
  - Safely and avoids race condition
- Wait Groups & Sync Mechanisms
  - sync.WaitGroup:
  - sync.mutex
- Select statement
  - mutiple channel operations

### Go Sync feature
- Provides the Sunc package to synchronize go routines
- `sync.WaitGroup` → Ensures all goroutines complete execution before continuing.
- `sync.Mutex` → Prevents multiple goroutines from modifying shared data at the same time.
- `sync.RMWMutex` -> reads by multiple routines and write only one
- `sync.Once -> executes only once
- `sync.Cond` -> Synchronization based on conditions
- `sync.Pool` Resuse objects to reduce memory allocations

### Go Context:
- The context package in Go is used for managing deadlines, cancellations, and request-scoped values across goroutines. 
- It is widely used in network requests, database operations, and concurrent programs to prevent resource leaks.
- withCancel,with Deadline, with timeOut, with value, context.background

### Memory Management
- Automatic Garbage Collection
- Stack and Heap Memory Allocation
- Escape Analysis
- Pointer Safety
- Zero-value Initialization
- Reuses memory called sync.Pool

### Functions in Go 
- Basic Functions
- Veriadic Functions
- Closures -> take varibale from outer / another function
- Function as input
- Methods: When a function is assigned to a Type it is called method 
- Anonymous Functions: 
- Function as a Argument.
