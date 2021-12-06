# RUST

### Key Concepts

- Rust's primary focus is on performance and safety
- Offers safe concurrency and memory safety without garbage collection
- Memory management is achived through *Ownership* concept
  - Each value in Rust has a variable that’s called its *owner*
  - There can only be one owner at a time
  - When the owner goes out of scope, the value will be dropped
- Borrowing allows variables to be borrowed without taking the ownership. Borrow checker statically ensures that references point to valid objects and ownership rules are not violated.
- **RAII** - Resource acquisition is initialization. Rust enforces RAII so that when a value is initialized the variable owns the resources associated and its destructor is called when the variable goes out of scope freeing the resources. This ensures that we will never have to manually free memory or worry about memory leaks.
- **Fearless concurrency** Rust has support for both concurrent programming and parallel programming.
- Rust’s ownership and type checking makes sure any concurrent errors are identified during compile time rather than runtime
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

### Benchmarking

Benchmarking was done using [Criterion.rs](https://github.com/bheisler/criterion.rs).
