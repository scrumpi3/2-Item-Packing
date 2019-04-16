### Tech

This application uses a number of open source projects to work properly:

* [Fake] - Generates fake data of specific types.
* [go] - the Go programming language

And of course this application itself can be open source with a [repository][ch_2] on GitHub and using the MIT liciense.

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
$ ./fakerData -file=<output file> -rows=<number of rows> -seed=<seed integer>
```

To find the best pair...
```sh
$ cd ${GOPATH}/src/github.com/scrumpi3/Challenge_2/find-pair
$ go build
$ ./find-pair <input file> <gift card balance>
```


[generator]: <https://github.com/scrumpi3/Challenge_2/tree/master/fakerData>
[implementation]: <https://github.com/scrumpi3/Challenge_2/tree/master/find-pair>
[test]: <https://github.com/scrumpi3/Challenge_2/tree/master/test>
[ch_2]: <https://github.com/scrumpi3/Challenge_2>
[fake]: <https://github.com/icrowley/fake>
[MIT_lic]: <https://opensource.org/licenses/MIT>
[go]: <https://golang.org>
[markdown-it]: <https://github.com/markdown-it/markdown-it>

