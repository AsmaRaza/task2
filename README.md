##  Task2
Read the file containing the numbers in the format where each line consists of one number.
The size of the numbers would be minimum 1000 and you have to calculcate the sum of numbers via go routines.
In this problem you have to split the data into multiple chunks and have to run multiple routines to calculate the sum.
The number of routines would be dynamic and will be taken from the command line and you have to make sure that the work is distributed  equally to each routine.
There will be one extra routine that will be respoisble for fetching the results from all of the routines and will sum up the results and print the sum.
## Description How to run with flags 
we take  routines and path from flags direcctly as a command line argument to set the path.
Run through the command like
go run task2.go --routines=5  --path=./numbers.txt

## output
![chunk](https://user-images.githubusercontent.com/93153939/179921220-7e255d51-568b-43ae-b7b9-09b18251cb1b.PNG)
