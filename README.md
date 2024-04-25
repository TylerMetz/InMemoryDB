# InMemoryDB

## Setup
- This code was generated in GoLang, to properly setup this code for running you must first install [Go](https://go.dev/doc/install).
- Once installed, clone the repo and follow the instructions for using the makefile

## How to Run
- To run this code, in the root of the repository there is a `Makefile` that will be able to properly build and run the code with the `src` folder.
- `build` in the makefile generates the Go executable that outputs the demo main.go showing the usage of the InMemoryDB
- `run` first builds, then runs the executable within the `bin` folder, however within our `.gitignore` those files are never committed

## Modifications to become an Offical Assignment
- To modify this assignment to become an official assignment I think that the addition of unit tests in the chosen language would be a very benefical perspective to help tie in other topics we have visited in the course. Also perhaps a new method for generating logs after the completion of a transaction to files would be a new perspective that we could learn. Lastly, a method that could output the Database to a file that could then be loaded if specified at the start of the program would be great improvements to bring this to the level of a great assignment.