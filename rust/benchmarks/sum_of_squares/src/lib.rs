extern crate rayon;
use rayon::prelude::*;

pub fn sos(input: &[u64]) -> u64 {
    input.iter()
         .map(|&i| i * 2)
         .sum()
}

pub fn sos_parallel(input: &[u64]) -> u64 {
    input.par_iter()
         .map(|&i| i * 2)
         .sum()
}

pub fn sum_of_squares(input: &[u64], n_threads: usize) {
    let pool = rayon::ThreadPoolBuilder::new().num_threads(n_threads).build().unwrap();
    
    pool.install(|| sos(&input));
}

pub fn sum_of_squares_parallel(input: &[u64], n_threads: usize) {
    let pool = rayon::ThreadPoolBuilder::new().num_threads(n_threads).build().unwrap();
    
    pool.install(|| sos_parallel(&input));
}