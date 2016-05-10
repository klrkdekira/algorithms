# algorithms

## Test Environment

```
Intel i7-4790K 4.00 GHz
16GB DDR3 1866MHz
Ubuntu Xenial 16.04
Go Version 1.6.2
```

## Benchmark Results

**Test method is to be revised**

### Insertion Sort

```
PASS
BenchmarkInsertionSort10-8     	100000000	        16.5 ns/op
BenchmarkInsertionSort100-8    	20000000	       114 ns/op
BenchmarkInsertionSort1000-8   	 1000000	      1017 ns/op
BenchmarkInsertionSort10000-8  	  100000	     10277 ns/op
BenchmarkInsertionSort100000-8 	       1	2196498996 ns/op
BenchmarkInsertionSort1000000-8	       1	219171944305 ns/op
ok  	github.com/klrkdekira/algorithms/sort/insertion	237.431s
```

### Merge Sort

```
PASS
BenchmarkMergeSort10-8     	 1000000	      1521 ns/op
BenchmarkMergeSort100-8    	  100000	     20120 ns/op
BenchmarkMergeSort1000-8   	   10000	    232590 ns/op
BenchmarkMergeSort10000-8  	     500	   3081814 ns/op
BenchmarkMergeSort100000-8 	      50	  36666126 ns/op
BenchmarkMergeSort1000000-8	       3	 351404011 ns/op
ok  	github.com/klrkdekira/algorithms/sort/merge	39.495s
```
