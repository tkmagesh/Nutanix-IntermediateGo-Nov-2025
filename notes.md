# Intermediate Go

## Magesh
- tkmagesh77@gmail.com

## Schedule
| What | Duration |
| ---- | -------- |
| Session - 01 | 1 hr 30 mins |
| Tea Break | 15 mins |
| Session - 02 | 1 hr 30 mins |
| Lunch Break | 45 mins |
| Session - 03 | 1 hr 30 mins |
| Tea Break | 15 mins |
| Session - 04 | 1 hr 30 mins |

## Methodology
- No powerpoints
- Discussion & Code

## Software Prerequisites
- go tools (https://go.dev/dl)
```shell
go version
```

## Go Prerequisites
- Data Types
- Variables , Contants, Iota
- Program Structure
- How to compile?
- Modules & Packages
- Programming constructs
    - if else, for, switch case
- Functions
    - Named Results
    - Higher Order Functions
    - Anonymous Functions
    - Variadic Functions
    - Deferred Functions
- Error Handling
- Panic & Recovery
- Structs
    - Struct composition
- Methods
- Type Assertion
- Interfaces
    - Interface Composition

## Repository
- https://github.com/tkmagesh/nutanix-intermediatego-nov-2025

## Concurrency 

### What is concurrency?
`Concurrency is NOT parallelism`

`Concurrent Design` is designing the application in such a way that it has more than one execution path.

### Managed Concurrency
- Builtin Scheduler in the application binary

### Language Support
- `go` keyword
- `chan` datatype
- `<-` operator 
- `range` and `select-case` constructs

### Standard Library Support
- `sync` package
- `sync/atomic` package

### Synchronization
#### sync.WaitGroup
- Semaphore based counter
- Has the ability to block the execution of a function until the counter becomes 0

### Data Race
```shell
go run --race <...>
```
```shell
go build --race <...>
```
```shell
go test --race <...>
```







