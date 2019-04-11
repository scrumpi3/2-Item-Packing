### Tech

This application uses a number of open source projects to work properly:

* [Fake] - Generates fake data of specific types.
* [go] - the Go programming language

And of course this application itself can be open source with a [repository][ch_2] on GitHub and using the MIT l\
iciense.

### Installation

The application requires [go] and has been tested using version 1.12.1.

Clone the repositoty into your GOPATH
```sh
$ cd ${GOPATH}/src/github.com/scrumpi3
$ git clone https://github.com/scrumpi3/Challenge_2.git
```

To generate demonstration datasets, with fake data...
```sh
$ cd ${GOPATH}/src/github.com/scrumpi3/Challenge_2/fakerData
$ go build
$ ./fakerData <# of rows> <output file>
```

To find the best pair...
```sh
$ cd ${GOPATH}/src/github.com/scrumpi3/Challenge_2/find-pair
$ go build
$ ./find-pair <input file> <gift card balance>
```

## Questions
### Big-O
What is the big O notation for your program?

> This program will require O(n) runtime and O(n) space.

## Bonus Questions

You are considering giving gifts to more people. Instead of choosing exactly 2 items, allow for 3
gifts.

> To allow for three gifts, I would add on additional product pointer to the algorithm, which incrementally incr\
eases in product price for products between the current high and low priced products.  I would follow that same \
process of calculating and minipulating the shrinking window, with a minor addition of the third pointer, used t\
o fine tune the total price of the three selected products.

How would you optimize your solution if you could not load the file in memory?
> I am going to assume the file cannot be loaded into memory becasue it is too large.
> Instead of using a single large file, I would subdivide the file into an ordered set of files that all follow \
the prerequsits as are defined in the orginal problem.
> I would use a data structure, like a B-tree, to store the first value of each of the subfiles.
> A stored value would be paired with a pointer to the value's source file.
> I would then implement the same general pointer movement, at the file-level, to quickly determine which two fi\
les contain a potentential solution.
> Then I would transition from the file-level to the product-level, continuing to executing the same algorithm t\
o find a potential solution.



### Todos

- Write MORE Tests

License
----
[MIT][MIT_lic]


[generator]: <https://github.com/scrumpi3/Challenge_2/tree/master/fakerData>
[implementation]: <https://github.com/scrumpi3/Challenge_2/tree/master/find-pair>
[test]: <https://github.com/scrumpi3/Challenge_2/tree/master/test>
[ch_2]: <https://github.com/scrumpi3/Challenge_2>
[fake]: <https://github.com/icrowley/fake>
[MIT_lic]: <https://opensource.org/licenses/MIT>
[go]: <https://golang.org>
[markdown-it]: <https://github.com/markdown-it/markdown-it>

