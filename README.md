# Challenge_2
Find best two items to consume the most of a gift card balance.  This is a packing problem that can be solved with a greedy algorithm.

The codeboase consists of three peices
- Sample dataset generator
- Algorithm implemtation
- Test cases

### Tech

This application uses a number of open source projects to work properly:

* [Fake] - Generates fake data of specific types.

* [markdown-it] - Markdown parser done right. Fast and easy to extend.
* [go] - duh

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
$ ./fakerData <# of rows> <output file>
```

To find the best pair...
```sh
$ cd ${GOPATH}/src/github.com/scrumpi3/Challenge_2/find-pair
$ go build
$ ./find-pair <input file> <gift card balance>
```



#### Building for source
For production release:
```sh
$ gulp build --prod
```
Generating pre-built zip archives for distribution:
```sh
$ gulp build dist --prod
```
### Docker
Dillinger is very easy to install and deploy in a Docker container.

By default, the Docker will expose port 8080, so change this within the Dockerfile if necessary. When ready, simply use the Dockerfile to build the image.

```sh
cd dillinger
docker build -t joemccann/dillinger:${package.json.version} .
```
This will create the dillinger image and pull in the necessary dependencies. Be sure to swap out `${package.json.version}` with the actual version of Dillinger.

Once done, run the Docker image and map the port to whatever you wish on your host. In this example, we simply map port 8000 of the host to port 8080 of the Docker (or whatever port was exposed in the Dockerfile):

```sh
docker run -d -p 8000:8080 --restart="always" <youruser>/dillinger:${package.json.version}
```

Verify the deployment by navigating to your server address in your preferred browser.

```sh
127.0.0.1:8000
```

### Todos

- Write MORE Tests

License
----
[MIT][MIT_lic]



[ch_2]: <https://github.com/scrumpi3/Challenge_2>
[fake]: <https://github.com/icrowley/fake>
[MIT_lic]: <https://opensource.org/licenses/MIT>
[go]: <https://golang.org>
[markdown-it]: <https://github.com/markdown-it/markdown-it>

