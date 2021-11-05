use criterion::{criterion_group, criterion_main, Criterion};
use parallel_iteration::sum_of_squares;

pub fn criterion_benchmark(c: &mut Criterion) {
    let mut arr: [i32; 2000000] = [0; 2000000];
    for i in 0..2000000 {
        arr[i] = rand::random();
    }
    c.bench_function("sum_of_squares", |b| b.iter(|| sum_of_squares(&arr)));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
