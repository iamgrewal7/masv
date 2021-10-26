
# MASV Assessment

## About
This Golang based program is part of MASV assessment and is built based on the following requirements.
- Accepts any number of parameters as command line arguments. 
- Each parameter should  be of the form <number><string> where <number> is a single digit from 0 to 9 and <string> is a lower case string composed of the characters from a to z with length of at least 1.
- Each parameter should be processed in the following way with the output from each parameter printed on its own line:
	 - The digit indicates how many character positions each character should be increased by, with wraparound from "z" to "a". 
	 - The modified string, without its digit prefix, should be printed to STDOUT.


## Usage

The project includes a `Makefile` . Follow the below steps to build, run and test the program.

1. Build the program by running `make build`
2. 
- To execute an example run `make run`. This will run the follow command:
	```sh
	./masv 0apple 1lzru 6gump 8hello 5zyo 9apoiuoqiwuero 1abcdefgh
	```
	The output should look like:
	```sh
	apple
	masv
	masv
	pmttw
	edt
	jyxrdxzrfdnax
	bcdefghi
	```
- Or run the built `./masv` binary with your own arguments.
3. To test the solutions run `make test`

There are **3** different solutions available in **solutions** folder. `make test` will test all the solutions. By default the program is using `SolutionOne`. In order to switch to different solution, uncomment the corresponding function call and comment the existing one. 

```go
result := solutions.SolutionOne(args)
// result := solutions.SolutionTwo(args)
// result := solutions.SolutionThree(args)
```
