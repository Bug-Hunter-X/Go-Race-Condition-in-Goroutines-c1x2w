# Go Race Condition in Goroutines

This repository demonstrates a common race condition bug in Go programs that use goroutines and shared mutable state.

## Description

The `bug.go` file contains a program that uses goroutines to increment a shared counter.  However, it lacks proper synchronization, leading to a race condition.  Multiple goroutines may try to update the counter simultaneously resulting in an incorrect final count.

The `bugSolution.go` file provides a solution that uses a mutex to protect access to the counter, preventing the race condition. 

## How to reproduce

1. Clone this repository.
2. Navigate to the directory in your terminal.
3. Run `go run bug.go` to see the incorrect count. (The count will likely be less than 1000)
4. Run `go run bugSolution.go` to see the corrected count (1000). 

## Solution

The solution utilizes a `sync.Mutex` to ensure only one goroutine can access the counter at any time, resolving the race condition.