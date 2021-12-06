use criterion::{criterion_group, criterion_main, Criterion};
use binarytree::binarytree;

pub fn criterion_benchmark(c: &mut Criterion) {
    let mut group = c.benchmark_group("fibo");
    group.bench_function("binary", |b| b.iter(|| binarytree()));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
