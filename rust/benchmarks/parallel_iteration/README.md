# Parallel iteration using Rayon

[Rayon](https://github.com/rayon-rs/rayon) is a data parallelism library for Rust, which makes it easier to convert sequential computation into parallel one.

We tested a simple iteration loop with Rayon's library and found runtime improvement when input size is greater than 200K.
