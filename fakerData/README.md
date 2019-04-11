# Generating Fake Product Data

This application generates a file with a specific number of fake product name and price pairs.

## Requirements
Go
github.com/icrowley/fake

## Usage
To install fake package

```sh
go get github.com/icrowley/fake
```

To build the application

```sh
go build . 
```

To view the application usage message

```sh
./fakerData -help
```

To test the application

```sh
go test
```

To execute the application and create an output file named products.txt with 10 products sorted by price

```sh
./fakerData -file=products.txt -rows=10 -seed=11
```


<details><summary>Example file produced</summary>
<p>

```
Dabtype Input Case,60
Agimba Direct Audible Case,111
Kwinu Performance Performance Compressor,124
Buzzbean Power Adapter,160
Devcast Disc Transmitter,194
Bluezoom Auto Disc Bridge,199
Plambee Power Amplifier,228
Yombu Direct Power Receiver,287
Oloo GPS Kit,371
Skinix Air Output Compressor,389
```

</p>
</details>

