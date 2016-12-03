# Benchmarks

Can be run by invoking: go test -bench=.

This will show performance for three sorting algorithms (the internal go one,
quicksort, and bubblesort) for three different array lenghts. On my machines
this is the output:

BenchmarkInternal100-4       	  500000	      3767 ns/op
BenchmarkInternal1000-4      	   10000	    109398 ns/op
BenchmarkInternal10000-4     	   10000	    109962 ns/op
BenchmarkBubbleSort100-4     	   50000	     24491 ns/op
BenchmarkBubbleSort1000-4    	    1000	   2159768 ns/op
BenchmarkBubbleSort10000-4   	       5	 216386462 ns/op
BenchmarkQuickSort100-4      	  300000	      4704 ns/op
BenchmarkQuickSort1000-4     	   20000	     87029 ns/op
BenchmarkQuickSort10000-4    	    2000	   1096880 ns/op

