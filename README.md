# Advent-of-Code-2017
Personal quick and dirty solutions to [Advent Of Code's](http://adventofcode.com) 2017 problems.


## Downloading and building the solutions
This project is structured as multiple Go projects in a single repository. The easiest way to get 
started is to let Go download the projects and dependencies for you.
 
```
# Download projects as well as dependencies
go get -u -v github.com/devosray/Advent-of-Code-2017/...
```

To build and run a project,`cd` to its folder and build it like any go application. For example,
to build and run `day-1`'s solution:
```
# cd into project folder
cd $GOPATH/src/github.com/devosray/Advent-of-Code-2017/
cd day-1 

# Build project
go build

# Run solution
./day-1 --input=1221 --part=1
```
