# go-vs-rust
Comparison of Go and Rust language features for efficient systems programming.


## Installation

The programs were run on Amazon EC2 C5 instances. The instructions for installion on running the programs on the machine are given below.

### Install Rust
```
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```
```
source $HOME/.cargo/env && sudo yum update && sudo yum groupinstall "Development Tools"
```
### Install Go
```
sudo amazon-linux-extras install epel && sudo yum update && sudo yum install golang
```

### Clone the GitHub repository
```
git clone https://github.com/ajitJJadhav/go-vs-rust.git
```


### Running the codes
We used sum of squares implementation (both parallel and non-parallel codes) for comparison of both languages. The parallel code was run of AWS EC2 C5 instances with multiple configurations (2 vCPU, 4 vCPU, 8 vCPU, 16 vCPU).

#### Rust:

To run the Rust code, go to `/rust/programs/sum_of_squares` and run the code using the following commands:

```
cargo build
```
```
cargo bench
```

This will run the non parallel and parallel codes and show the benchmarks for the same.

#### Go:

To run the Go code, go to `/Go/programs` folder in this repository and run the `parallel.go` code using the command:

```
go run parallel.go
```

This will run the non parallel version followed by the parallel version of the code with the run times of both.
(There are a few other codes in the same folder which can be run, but some of them are incomplete)
