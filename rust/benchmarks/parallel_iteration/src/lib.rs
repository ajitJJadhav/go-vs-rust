extern crate rayon;
use rayon::prelude::*;

pub fn sum_of_squares(input: &[i32]) -> i32 {
    input.iter()
         .map(|&i| i * i)
         .sum()
}

pub fn sum_of_squares_parallel(input: &[i32]) -> i32 {
    input.par_iter()
         .map(|&i| i * i)
         .sum()
}

pub fn sum_of_squares_inp(input: &[i32]) {
    let pool = rayon::ThreadPoolBuilder::new().num_threads(1).build().unwrap();
    
    pool.install(|| sum_of_squares_parallel(&input));
}

pub fn sum_of_squares_parallel_1(input: &i32) -> i32 {
    let a = (0..*input).into_par_iter()
         .map(|i| i * 2)
         .sum();
    println!("{}", a);
    return a;
}

pub fn sum_of_squares_inp1(input: &i32) {
    let pool = rayon::ThreadPoolBuilder::new().num_threads(6).build().unwrap();
    
    pool.install(|| sum_of_squares_parallel_1(input));
}