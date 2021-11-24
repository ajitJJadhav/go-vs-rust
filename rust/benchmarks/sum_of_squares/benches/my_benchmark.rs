use criterion::{criterion_group, criterion_main, Criterion};
use sum_of_squares::{sum_of_squares, sum_of_squares_parallel};

pub fn criterion_benchmark(c: &mut Criterion) {
    let mut group = c.benchmark_group("sum_of_squares");
    group.sample_size(10);
    let mut arr: [i32; 2000000] = [0; 2000000];
    for i in 0..2000000 {
        arr[i] = rand::random();
    }

    //group.bench_function("ss 1M_1", |b| b.iter(|| sum_of_squares(&arr[0..1000000], 1)));
    group.bench_function("ss 2M_1", |b| b.iter(|| sum_of_squares(&arr, 1)));
    group.bench_function("ss 2M_2", |b| b.iter(|| sum_of_squares(&arr, 2)));
    group.bench_function("ss 2M_4", |b| b.iter(|| sum_of_squares(&arr, 4)));
    group.bench_function("ss 2M_8", |b| b.iter(|| sum_of_squares(&arr, 8)));
    group.bench_function("ss 2M_16", |b| b.iter(|| sum_of_squares(&arr, 16)));

    //group.bench_function("ssp 1M_1", |b| b.iter(|| sum_of_squares_parallel(&arr[0..1000000], 1)));
    group.bench_function("ssp 2M_1", |b| b.iter(|| sum_of_squares_parallel(&arr, 1)));
    group.bench_function("ssp 2M_2", |b| b.iter(|| sum_of_squares_parallel(&arr, 2)));
    group.bench_function("ssp 2M_4", |b| b.iter(|| sum_of_squares_parallel(&arr, 4)));
    group.bench_function("ssp 2M_8", |b| b.iter(|| sum_of_squares_parallel(&arr, 8)));
    group.bench_function("ssp 2M_16", |b| b.iter(|| sum_of_squares_parallel(&arr, 16)));

}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
