## How to run the program

To run the program, you need to ensure you have Go installed. You can download Go [here](https://golang.org/dl/).

The commands provided below will be in bash, but should be the same in any command language, though the paths might be formatted a bit differently.

Then run the command: `go mod tidy`.

Next, open four separate terminals, one for the hospital and three terminals for the different clients:

1. **Hospital**: `go run ./hospital/hospital.go`.

2. **Client0**: `go run ./patient/patient.go -id=0 -input=1`.

3. **Client1**: `go run ./patient/patient.go -id=1 -input=2`.

4. **Client2**: `go run ./patient/patient.go -id=2 -input=3`.

You will be able to follow the progress of each process in their terminal, and the program will be done once the hospital writes out: `Sum of shares: 6`.

Note that in this example the sum is 6, but that can change depending on which numbers are given as inputs to the clients.

The program does not terminate by itself, as they are left running to ensure that each process can do what it needs to without the others shutting down prematurely.