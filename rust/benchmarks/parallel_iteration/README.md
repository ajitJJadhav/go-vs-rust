# Parallel iteration using Rayon

[Rayon](https://github.com/rayon-rs/rayon) is a data parallelism library for Rust, which makes it easier to convert sequential computation into parallel one.

We tested a simple iteration loop with Rayon's library and found runtime improvement when input size is greater than 200K.

Below is the running time captured using Criterion benchmark tool. 

| Input Size  | Sequential (us) | Parallel (us) |
| ----------- | --------------- | ------------- |
|   50000     |    13.06        |     34.22     |
|  100000     |    25.69        |     40.02     |
|  200000     |    51.92        |     49.49     |
|  500000     |   131.14        |     73.78     |
| 1000000     |   267.29        |    112.06     |
| 2000000     |   588.75        |    195.36     |

Run time is higher for parallel execution when input size is smaller, this is due to the overhead from thread switching.
