extern crate rayon;
use rayon::prelude::*;

pub fn sum_of_squares(input: &[i32]) -> i32 {
    input.par_iter()
         .map(|&i| i * i)
         .sum()
}