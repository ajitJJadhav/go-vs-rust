use criterion::{criterion_group, criterion_main, Criterion};
use parallel_iteration::{sum_of_squares, sum_of_squares_parallel, sum_of_squares_inp, sum_of_squares_inp1};

pub fn criterion_benchmark(c: &mut Criterion) {
    let mut group = c.benchmark_group("sum_of_squares");
    group.sample_size(10);
    let mut arr: [i32; 2000000] = [0; 2000000];
    for i in 0..2000000 {
        arr[i] = rand::random();
    }
    let inp = 2000000;
 /* group.bench_function("ss 50000", |b| b.iter(|| sum_of_squares(&arr[0..50000])));
    group.bench_function("ss 100000", |b| b.iter(|| sum_of_squares(&arr[0..100000])));
    group.bench_function("ss 200000", |b| b.iter(|| sum_of_squares(&arr[0..200000])));
    group.bench_function("ss 500000", |b| b.iter(|| sum_of_squares(&arr[0..500000])));
    group.bench_function("ss 1000000", |b| b.iter(|| sum_of_squares(&arr[0..1000000])));
    group.bench_function("ss 2000000", |b| b.iter(|| sum_of_squares(&arr)));

    group.bench_function("ssp 50000", |b| b.iter(|| sum_of_squares_parallel(&arr[0..50000])));
    group.bench_function("ssp 100000", |b| b.iter(|| sum_of_squares_parallel(&arr[0..100000])));
    group.bench_function("ssp 200000", |b| b.iter(|| sum_of_squares_parallel(&arr[0..200000])));
    group.bench_function("ssp 500000", |b| b.iter(|| sum_of_squares_parallel(&arr[0..500000])));
    group.bench_function("ssp 1000000", |b| b.iter(|| sum_of_squares_parallel(&arr[0..1000000])));
    group.bench_function("ssp 2000000", |b| b.iter(|| sum_of_squares_parallel(&arr)));*/

    //group.bench_function("ss 201", |b| b.iter(|| sum_of_squares_inp(&arr)));
    group.bench_function("ss 202", |b| b.iter(|| sum_of_squares_inp1(&inp)));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
