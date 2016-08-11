# Quicksort
[![Build Status](https://travis-ci.org/jutkko/quicksort.svg?branch=master)](https://travis-ci.org/jutkko/quicksort)

This repo contains a quicksort algorithm implemented in golang.

# Benchmark
Comparison between randomised and non-randomised quicksort

```
BenchmarkLargeArray-4                   2000000000               0.07 ns/op
BenchmarkLargeArrayBuiltinSort-4        1000000000               0.26 ns/op # This is so weird, the built in sort is slower than mine?
BenchmarkLargeOrderedArray-4            2000000000               0.04 ns/op
BenchmarkLargeOrderedArrayBuiltinSort-4 2000000000               0.06 ns/op
```

# Reference
- Quicksort [tutorial](https://www.khanacademy.org/computing/computer-science/algorithms/quick-sort/a/overview-of-quicksort)
- Go test [tutorial](https://www.binpress.com/tutorial/getting-started-with-go-and-test-driven-development/160)
- Go [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)
- Randomised [quicksort](http://alg12.wikischolars.columbia.edu/file/view/QUICKSORT.pdf), cool tip!

# Project52
This is a project from my [Project52](https://github.com/jutkko/project52).
