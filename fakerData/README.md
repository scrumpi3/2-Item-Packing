# Generating Fake Product Data

This application generates a file with a specific number of fake product name and price paires.

## Requirements
Go

## Usage
To build the application

'''sh
go build . 
'''

To view the application usage message

'''sh
./fakerData -help
'''

To execute the application and create an output file named products.txt with 10 products sorted by price

'''sh
./fakerData -file=products.txt -rows=10 -seed=11
'''