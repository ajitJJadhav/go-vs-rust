extern crate rayon;
use rayon::prelude::*;

pub fn sos(input: &[u32]) -> u32 {
    input.iter()
         .map(|&i| i * i)
         .sum()
}

pub fn sos_parallel(input: &[u32]) -> u32 {
    input.par_iter()
         .map(|&i| i * i)
         .sum()
}

pub fn sum_of_squares(input: &[u32], n_threads: usize) {
    let pool = rayon::ThreadPoolBuilder::new().num_threads(n_threads).build().unwrap();
    
    pool.install(|| sos(&input));
}

pub fn sum_of_squares_parallel(input: &[u32], n_threads: usize) {
    let pool = rayon::ThreadPoolBuilder::new().num_threads(n_threads).build().unwrap();
    
    pool.install(|| sos_parallel(&input));
}