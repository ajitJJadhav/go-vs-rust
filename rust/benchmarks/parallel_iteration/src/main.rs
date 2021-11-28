extern crate parallel_iteration;
//extern crate rayon;
//use rayon::prelude::*;

fn main() {
    /*let pool = rayon::ThreadPoolBuilder::new().num_threads(4).build().unwrap();
    let mut arr: [i64; 1000000] = [0; 1000000];
    for i in 0..1000000 {
        arr[i] = 2;//rand::random();
    }
    let _nss = pool.install(|| sum_of_squares_parallell(&arr));*/
}


/*fn sum_of_squares_parallell(input: &[i64]) -> i64 {
    input.par_iter()
         .map(|&i| i * i)
         .sum()
}*/
