# Greatest Sum of Two Items that is Less then Maximum
Find best two items to consume the most of a gift card balance. 
This is a packing problem that can be solved with a greedy algorithm.

The codebase consists of two components
- [Sample dataset generator][generator]
- [Algorithm implemtation][implementation]

## Approach

A shrinking window gready algorithm was used to solve this problem.  
The window is initialized to the __top__ and __bottom__ of the product list.
The window is then shrunk, starting from the top, until a two product cost can consume all or part of the gift card.
This potential solution is then fine tuned by shrinking the window again, this time from the bottom, until the next value is greater than the gift card. 
The algorithm has now found a valid solution, which is tracked as the currect best product pair.
The __top__ of the window in then decreased and the process continues until the window is closed.

## Questions
### Big-O: What is the big O notation for your program?

> This program will require O(n) runtime and O(n) space.

### Bonus Questions
You are considering giving gifts to more people. Instead of choosing exactly 2 items, allow for 3
gifts.

> To allow for three gifts, I would add an additional product pointer to the algorithm, which incrementally increases in product price for products between the current high and low priced products.  
> I would follow that same process of calculating and minipulating the shrinking window, with a minor addition of the third pointer, used to fine tune the total price of the three selected products.

How would you optimize your solution if you could not load the file in memory?
> I am going to assume the file cannot be loaded into memory becasue it is too large.  
> Instead of using a single large file, I would subdivide the file into an ordered set of files that all follow the prerequsits as are defined in the orginal problem.
> I would use a data structure, like a B-tree, to store the first value of each of the subfiles.
> A stored value would be paired with a pointer to the value's source file.
> I would then implement the same general pointer movement, at the file-level, to quickly determine which two files contain a potentential solution. 
> Then I would transition from the file-level to the product-level, continuing to executing the same algorithm to find a potential solution.  



### Todos

- Write MORE Tests

License
----
[MIT][MIT_lic]


[generator]: <https://github.com/scrumpi3/Challenge_2/tree/master/fakerData>
[implementation]: <https://github.com/scrumpi3/Challenge_2/tree/master/find-pair>
[ch_2]: <https://github.com/scrumpi3/Challenge_2>
[fake]: <https://github.com/icrowley/fake>
[MIT_lic]: <https://opensource.org/licenses/MIT>
[go]: <https://golang.org>
[markdown-it]: <https://github.com/markdown-it/markdown-it>

