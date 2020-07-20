# sorting_algorithms

## Implementation of sorting algorithms in Python3 and Go



### TODO
 - [ ] make comparison charts for execution time
 - [ ] add automation to execute everything



Quicksort	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	{\displaystyle n^{2}}n^{2}	{\displaystyle \log n}\log n	No	Partitioning	Quicksort is usually done in-place with O(log n) stack space.[5][6]
Merge sort	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	n	Yes	Merging	Highly parallelizable (up to O(log n) using the Three Hungarians' Algorithm).[7]
In-place merge sort	—	—	{\displaystyle n\log ^{2}n}n\log ^{2}n	1	Yes	Merging	Can be implemented as a stable sort based on stable in-place merging.[8]
Introsort	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	{\displaystyle \log n}\log n	No	Partitioning & Selection	Used in several STL implementations.
Heapsort	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	1	No	Selection	
Insertion sort	n	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	1	Yes	Insertion	O(n + d), in the worst case over sequences that have d inversions.
Block sort	n	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	1	Yes	Insertion & Merging	Combine a block-based {\displaystyle O(n)}O(n) in-place merge algorithm[9] with a bottom-up merge sort.
Quadsort	n	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	n	Yes	Merging	Uses a 4-input sorting network.[10]
Timsort	n	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	n	Yes	Insertion & Merging	Makes n comparisons when the data is already sorted or reverse sorted.
Selection sort	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	1	No	Selection	Stable with {\displaystyle O(n)}O(n) extra space or when using linked lists.[11]
Cubesort	n	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	n	Yes	Insertion	Makes n comparisons when the data is already sorted or reverse sorted.
Shellsort	{\displaystyle n\log n}n\log n	{\displaystyle n^{4/3}}{\displaystyle n^{4/3}}	{\displaystyle n^{3/2}}n^{3/2}	1	No	Insertion	Small code size.
Bubble sort	n	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	1	Yes	Exchanging	Tiny code size.
Tree sort	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n(balanced)	n	Yes	Insertion	When using a self-balancing binary search tree.
Cycle sort	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	1	No	Insertion	In-place with theoretically optimal number of writes.
Library sort	n	{\displaystyle n\log n}n\log n	{\displaystyle n^{2}}n^{2}	n	Yes	Insertion	
Patience sorting	n	—	{\displaystyle n\log n}n\log n	n	No	Insertion & Selection	Finds all the longest increasing subsequences in O(n log n).
Smoothsort	n	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	1	No	Selection	An adaptive variant of heapsort based upon the Leonardo sequence rather than a traditional binary heap.
Strand sort	n	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	n	Yes	Selection	
Tournament sort	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	n[12]	No	Selection	Variation of Heap Sort.
Cocktail shaker sort	n	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	1	Yes	Exchanging	
Comb sort	{\displaystyle n\log n}n\log n	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	1	No	Exchanging	Faster than bubble sort on average.
Gnome sort	n	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	1	Yes	Exchanging	Tiny code size.
UnShuffle Sort[13]	n	kn	kn	n	No	Distribution and Merge	No exchanges are performed. The parameter k is proportional to the entropy in the input. k = 1 for ordered or reverse ordered input.
Franceschini's method[14]	—	{\displaystyle n\log n}n\log n	{\displaystyle n\log n}n\log n	1	Yes	?	
Odd–even sort	n	{\displaystyle n^{2}}n^{2}	{\displaystyle n^{2}}n^{2}	1	Yes	Exchanging	Can be run on parallel processors easily.
