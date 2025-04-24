## Go is Compiled or Interpreted  Language `Why Go`
- Go is A compiled and statically typed
- Go Tool Can run a file without precompiling
- Compiled executables are os specific
- It is invented By Google
- Concurrency
- Efficient compilation
- efficient execution
- ease of programming
- No runtime environment to worry about
- Object Oriented (sort of)
- cross platform
- excellent package management
- 

### Use Cases:
- Go Is Statically Typed Language

## Types:
- Case Sensitive
- Variable should be known in advance
- Everything is Type here
- String,Bool,Integer,Floating,Complex,
- uint8,uint64

## Comma and ok
- User inputs using scanf
- Using Bufio.io & Bufio.io standard

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
- Veriadic Functions Funcname(nums ...int) --> anynumers you can give
- Closures -> take varibale from outer / another function
- Function as input
- Methods: When a function is assigned to a Type it is called method 
- Anonymous Functions: 
- Function as a Argument.

### Defer Statements:
- Will make the compiler to run at the end of the function

### Handling The web request
- 
### Interfaces --> Like a Contracts 
- It can implement the methods whihc i having over riden values

IBM Job Opening:

🚀 We’re Hiring: Full Stack Developer (Golang/Python & Terraform) – Kochi 🌍
Are you a passionate Full Stack Developer looking for an exciting opportunity in Kochi? Do you thrive in building scalable applications and working with cutting-edge technologies like Golang, Python, and Terraform (IaC)? If yes, we want to hear from you!

🔹 What You’ll Do:
✔️ Develop and maintain robust backend services using Golang/Python
 ✔️ Design and implement scalable Infrastructure as Code (IaC) with Terraform
 ✔️ Collaborate with cross-functional teams to create seamless applications
 ✔️ Optimize performance and ensure best coding practices

🔹 What We’re Looking For:
🔹 2+ years of experience in Golang or Python development
 🔹 Hands-on experience with Terraform for cloud infrastructure management
 🔹 Strong knowledge of microservices architecture, APIs, and cloud platforms (AWS/GCP/Azure)
 🔹 Passion for clean, efficient, and scalable code


 ### Building The api:
 - API Building 
 - JWT Token
 - Event Booking API
 - Get List Update Post Put Fetch

### Packages
- Http Packages: Handled Http requests
- Gin Library `Developed BY Go Community`
- go get -u github.com/gin-gonic/gin

### Master Go Using Concurrency Patterns
- 

### Postgress Details 

- docker pull postgres:17.4-alpine3.21
- docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=949383Hk -d postgres:17.4-alpine3.21

## Building Modules with go
- Using Go Workspaces
- 