extern crate rayon;
use rayon::prelude::*;

fn sos(input: &[i32]) -> i32 {
    input.iter()
         .map(|&i| i * i)
         .sum()
}

fn sos_parallel(input: &[i32]) -> i32 {
    input.par_iter()
         .map(|&i| i * i)
         .sum()
}

pub fn sum_of_squares(input: &[i32], n_threads: usize) {
    let pool = rayon::ThreadPoolBuilder::new().num_threads(n_threads).build().unwrap();
    
    pool.install(|| sos(&input));
}

pub fn sum_of_squares_parallel(input: &[i32], n_threads: usize) {
    let pool = rayon::ThreadPoolBuilder::new().num_threads(n_threads).build().unwrap();
    
    pool.install(|| sos_parallel(&input));
}