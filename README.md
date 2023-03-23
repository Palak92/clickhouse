# Clickhouse


## Steps to build (MacOS)


### Prerequisites

1. Install [homebrew](https://brew.sh)

2. Install golang. Go version must be at least 1.18
```
brew install go
```

3. Install [bazel](https://bazel.build/install) in your system.
```
brew install bazel
```


### Clone Repo

1. Make directory and cd into it to clone the project.
```
mkdir -p ${GOPATH}/src/github.com/palak92/ 
```
Note : substitute your GOPATH

2. Change directory to workspace
```
cd ${GOPATH}/src/github.com/palak92
```

3. Checkout the code
```
git clone https://www.github.com/palak92/clickhouse.git
```
4. Change directory to project directory
```
cd clickhouse
```


### Build

1. Generate build files
```
bazel run //:gazelle
```

2. Run build command
```
bazel build //...
```

### Run program

1. Using bazel

 ```
 bazel run //cmd:cmd
 ```
2. Give absolute path to file.

### Logic 

1. File is processed line by line as file is quite large. Another way could have been doing batch processing of file data.

2. MinHeap can have maximum 10 items at one time.

### Improvements

1. More tests 
2. Batch processing of file 