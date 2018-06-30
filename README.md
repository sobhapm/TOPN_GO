### Short Description

The program topN will output the largest N numbers, highest first, given an arbitrarily large file  containing individual numbers on each line (e.g. 200Gb file) and a number N.

### How to compile

go build topN.go

### How to run the program

At the command-line type the following - 

./topN -input \<Input filename\> -N \<Number of largest numbers\>

For example,

./topN -input ./testcase1.txt -N 10

Expected Output---

Top N -----  
9999999  
9999998  
9999997  
9999996  
9999995  
9999994  
9999993  
9999992  
9999991  
9999990  

### To run the test script

At the command-line type the following -

./test.sh 
