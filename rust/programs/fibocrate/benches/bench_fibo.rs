use criterion::{black_box, criterion_group, criterion_main, Criterion};
use fibocrate::fibonacci;

pub fn criterion_benchmark(c: &mut Criterion) {
    let mut group = c.benchmark_group("fibo");
    group.bench_function("fib 50", |b| b.iter(|| fibonacci(black_box(50))));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
