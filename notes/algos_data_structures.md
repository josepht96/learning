# Lecture 2 
Program <--> algorithm
Programming language <--> pseudo code
Computer <--> model of computation 
Model of Computation: specifies what operations an algorithm is allowed to do, what is the cost of each operation. 
Random Access Machine (arrays), operations take constant time
Random Access Memory (just giant array all the way up to however many gigs)
Computing with RAM,  in constant time an algorithm:
load words O(1) into registers (word here means w bits, some type of data) 
do O(1) operations
write words back O(1)
Pointer Machine (pointers and references)
Dynamically allocated objects
object has constant number of fields
a field is either a word or a point (points to another object or nil)
operations take constant time still
Constant time = O(1)
Linear time  = O(n)
Consider the time complexity of certain operations.
In python:
	list.append() = O(1)
	list.len() = O(1)
	list = list1 + list2  = O(n)
	list.Sort() = O(nlog(n))

Document distance:
d(D1, D2) - how similar are documents D1 and D2, sequences of words
word = string of alphanumeric chars
idea: look at shared words
document is a vector, D[w] = number of occurrences of word in document
multi dimensional representation of the word frequency of each document
each word gets its own dimension with a length of frequency of that word
can take the dot product of the vectors to compare, { d1[w] . d2[w] 
a little closer if you divide dot product by the combined length of the vectors
({ d1[w] . d2[w] )/ |d1|+|d2|
split document into words
compute word frequencies (get vectors)
get dot product
Regex is not always linear time. 

# Lecture 3
Why is sorting important? Some applications logically require sorting, phonebook or mp3 player.
Some problems become easier once items are sorted, like finding a median.
Binary search: Looking for an item in [0:n], if unsorted you can find it in O(n). If sorted you can find it in O(log(n)). Data compression, computer graphics, etc… all use sorting. There are points where high time complexity algorithms are faster than lower complexity ones because of constant factors (c). 
Insertion sort:
	[5, 2, 4, 6, 1, 3]
First num is sorted, second num is called key (at first)
Insertion sort:
Swap the key against the values before it. At each point if the key
is less than the value adjacent to it, swap them
So a key might be swapped many times to get to proper spot
[5, 2, 4, 6, 1, 3] > [2, 5, 4, 6, 1, 3] > [2, 4, 5, 6, 1, 3]
-> [2, 4, 5, 1, 6, 3] -> [2, 4, 1, 5, 6, 3] etc...
Each step is O(n), so worst case is O(n^2)
Can be turned into O(nlgn) if you replace comparison with a binary search
Merge sort:
	Get two sorted arrays and compare them from least to greatest
	[5, 2, 4, 6, 1, 3] > [5, 2, 4], [6, 1, 3]
	T(n) = c1 + 2T(n/2) + c * n
	While merge sort is faster than insertion sort, merge sort requires more memory as the
split arrays are new areas. In place merge sort solves this problem, but the operations
take a significantly longer than normal merge sort so nobody uses it.
	Given T(n) = 2T(n/2) + cn^2, the merge routine takes n^2. Where the work is done is
determined by the cn^2. When that portion is small, say c(1), the work is done by the 
	leaves as you’d have n * c(1). 
Priority Queue: Heap, max heap, min heap
Structure that implements a set S of elements, each element is associated with a key. 
Properties, methods:
	Insert(S, x) element into set S
	max(S) - return element with largest key
	extract-max(S) - returns largest key and removes it from S
	increase-key(S, x, k) - increase value of x.key by k
Heap is an implementation of a priority queue. Implemented is an array but it's a ‘nearly complete’ binary tree. Max-heap - the key of a node is greater than or equal to the keys of its child nodes. Needs to be true for every node of the tree. Min-heap is the opposite. Max value is trivially performed on a Max-heap tree. Extract max would not be trivially performed as you'd need to rearrange the tree (the top element is the max). If you can maintain max heap property you can return values in sorted order. 
Heap operations: build max heap - produces max heap from unordered array. 
max_heapify: Assume trees rooted at left(i), right(i) are max heaps. Bottom up approach as leaves are by definition max heaps. Exchange node value with left child node, if child node becomes non max heap, then do the same for the child node, repeat indefinitely. 
Max heapify has a time complexity O(lgn) - the height of the tree is bounded by lgn, so in the worst case you have to go the height of the tree, which would be lgn. 
max heapify takes O(n). The lowest nodes, n/2, going up to n/4, n/16… n/n
0c(n) + 1c(n/2) + 2c(n/4) … = c(n) = O(n)

# Lecture 5: Binary Search Trees, BST Sort
Runway reservation system:
single runway
reservations for future landings
reserve requests specify a time
add t to the set R if no landings are within k time
remove from set R after the plane lands
Goal: |R| = n, O(logn)
Options: 
Unsorted list/array - value cannot be found O(n) (everything is linear)
Sorted list/array - Finding is O(lgn) but shifting values for insertion is O(n)
Linked list - searching is O(n), insertion is O(1)
Heap - bad search time, comparison takes O(n)
Solution for fast search and fast insert: Binary search tree
Array, visualized as a tree, with pointers to parent(x), left(x), right(x)
For the root node, all values on the left are less, all values on the right are greater. 
If the insert value is greater than a node, go right, else go left. If there is no child node, insert on the proper side. This means insertion time is O(height of the tree).
Find min: go all the way left
Find max: go all the way to the right
Fairly flexible data structure (augmentation)
If you number the number of times a node is passed through (including one for itself) you can track how many sub nodes a node has. Sub tree size = nodes beneath + 1 (itself)
Search time is O(n) if the insertion is done with an ordered list. 
# Lecture 6: AVL trees
# Lecture 7: Counting Sort, Radix Sort, Lower Bounds for Sorting
Linear sorting: Sorting requires O(nlgn) in worst case, and sometimes O(n)
Decision tree:
internal nodes
leaf
root-to-leaf
path length
worst case run time: height of the tree
Algorithm:
binary decisions comparisons
find the answer
worst case run time
Searching lower bound:
n preprocessed items
finding given item among items in comparison model G(lgn) (lower bound, omega) in the worst case
Height will always be at least logn G(lgn)

Sorting around O(n) time. As long as n isnt giant, you can use Counting Sort to get pretty close to O(n)

# Lecture 8 - Hashing with Chaining
Used everywhere, constant lookup time O(1)
Storing at index is bad because keys might not be easily made into integers.
Giant memory hog
Solution to non integer values is prehash: maps keys to non-negative integers. Keys should be finite and discrete
reduce all possible keys down to smaller set of integers\
Want size of array to be near the size of the input m = L(n)
enter changing to deal with collisions
Chaining: Idea is simple: if there are multiple items with the same hash, store them as a list (linked list for example). Worst case is all values hash to the same value/index so enter input is just stored in a linked list, but doesn’t really happen in reality. 
Analysis: expected length of a chain
n keys, m slots: probably is n/m
Hash functions:
Division: h(k) = k mod m
Multiplication: h(k) = [(a*k) mod 2^w] >> (w-r) -- basically just generates random number from mixing up k with a bit shift for every 1 value, then taking ~the middle values (w-r)
Universal hashing (cool method?): h(k) = [(ak + b) mod p] mod m. The Probability of two h(k) equals each other is 1/m. 

# Lecture 9 - Table Doubling, Karp-Rabin
Want the size of the hash array to be relatively close in size to the size of the input, grow and shrink the table as necessary. Hash function has to output values within the range of the array.
Allocate memory
Rehash values into new array O(n) - Not great if this happens a lot
Bad idea: m’ = m + 1 -> constant of n insertions = O(n^2)!
Better idea: m’ = 2m (Table doubling). Only have to go through O(n) a few times. O(n) time, O(nlogn) for the entire run (Table doubling). Most inserts are still constant, but every doubling requires O(n). 
Amortization: 
Operation takes “T(n) amortized” if k operations take <= k * T(n) time
spreading out total run cost over individual operations
All operations/total time. O(1) is lower bound.
In general, expensive operations are acceptable if most operations are cheap as n grows
Deletion: if m = n/2 then shrink the table, not great if you’re inserting/deleting at a value that requires the table to double O(n). 
if m = n/4 then shrink to m/2. Amortized time shrinks to O(1)
Python lists: Array but a number of operations happen in O(1). Table doubling solves this problem. Could also start building the hash table of size 2m while the original m is filling up. 

String matching (grep): O(s * t) because in the worst case you have to compare each char against the string chars until it matches. t = “aaaaaaaaab”, s = “aab” = ~ s * t operations.
Could split t into substrings of size len(s), call them r, hash each, see if s has a match. Since r contains len(s-1) of the same characters as the r before it, can just remove the first character of r and append the next character, this would take O(n), O(len(t). This is the Karp-Rabin string matching algorithm, using the rolling hash ADT method. 

# Lecture 10 - Open Addressing, Cryptographic hashing
Open addressing is the simplest way to implement a hash table. 
One item per slot, m >= n, slots greater than elements (ideally the same I guess)
Probing - check if there will be a collision, if there will be then adjust the hash function slightly. 
When a value will cause a collision, try h(x, 1) -> h(x, 2) to see if a different hash function yields a non collision insertion. 
Insertion works the same: while h(k) returns nil values && values not equal to k, keep probing/incrementing the hash function. 
Problem with delete is if you remove a key, the search function will be thrown off (as it will hit null even though the probe would return: key key key nil key key… etc. 
Solution to this issue is to replace the delete item with a value different than nil, but one that still indicates a value is not assigned to that key. 
Insert can insert into nil or deleted-flag, search only looks for nil ends, but passes through deleted-flags, delete replaces value with deleted-flag. 
Linear probing: h(k, i) = (h’(k) + i ) mod m - suspect to clustering (bunch of sequential collisions) 
Better: Double hashing - h(k, i) = (h1((k) + i h2(k)) mod m

# Lecture 11 -  Integer Arithmetic, Karatsuba Multiplication
Irrationals:
Catalan numbers: Have a recursive definition

# Lecture 13 - Breadth-First Search (BFS)
Graph searching is about exploring a graph
G = (V, E), size is about O(V+E)
V = set of vertices
E = set of edges
Undirected, directed graphs
Undirected, E = {all edges}
Usage: Web crawling, social networking, network broadcast, garbage collection, model checking, mathematical conjectures, solving puzzles and games
diameter of the graph is the number of branches off a root choice



Graph representation: 
Adjacency lists:
array adj of |V|
linked list
for each vertex u E V, adj[u]
for every vertex, store its neighbors
store pointer to node at end of edge

Breadth-first search:
Find all vertices that are reachable from a node, in O(V+E) time
Look at nodes reachable in zero moves (root), then 1, then 2, etc…
Careful to avoid duplicates, edges that point to the same vertices.
