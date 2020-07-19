# How to Debug Golang with VS Code

## Summary

- [Basic](#basic)
- [Spec](#spec)
- [Instruction](#instruction)
- [debugging unit test](#debugging-unit)
- [debugging executable file](#debugging-executable-file)
- [debugging remote process](#debugging-remote-debug)
- [debugging running process](#debugging-running-process)

## Basic

- [The Go Programming Language](https://golang.org/)
- Extension: [Go](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go)
- Debugger: [delve](https://github.com/derekparker/delve)
- module code: [bubbleSort.go](https://github.com/74th/vscode-debug-specs/blob/master/golang/bubbleSort.go)

## Spec

- OS
  - ✅ MacOS
  - ✅ Windows
  - ✅ Linux
- Break Point
  - ✅ break point
  - ✅ condition break point
  - ❌ function breakpoint
- Step Execution
  - ✅ Step Over
  - ✅ Step Into
  - ✅ Step Out
  - ✅ Continue
- Variables
  - ✅ variables views
  - ✅ watch variables
- Call Stack
  - ✅ call stack
- Evaluation
  - ✅ eval expression to show variables
  - ✅ eval expression to change variables
- Type of Execution
  - ✅ debug unit test
  - ✅ debug executable package
  - ✅ remote debugging

## Instruction

### MacOS

1. install go : `brew install golang`
1. add go/bin to PATH
1. [install extension "Go"](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go)
1. install other tools: `F1`->`Go: Install/Update Tools`

### Windows

1. install go and add go/bin to PATH
2. [install extension "Go"](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go)
3. install other tools: `F1`->`Go: Install/Update Tools`

### Linux

1. install go and add go/bin to PATH
2. [install extension "Go"](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go)
3. install other tools: `F1`->`Go: Install/Update Tools`

## unit test

source : [bubbleSort_test.go](https://github.com/74th/vscode-debug-specs/blob/master/golang/bubblesorter/bubbleSort_test.go)

### inline

![inline unit test](inline_unit_test.png)

### launch json

menu:`Go: Launch test function`

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch test function",
      "type": "go",
      "request": "launch",
      "mode": "test",
      "program": "${workspaceFolder}",
      "args": [
        "-test.run",
        // test function name
        // * can use reguler expression
        // * NOT include "Test"
        // * the first charactor MUST be small
        "bubblesort"
      ]
    }
  ]
}
```

- `program` must be package folder

### using Test Explorer

install [Go Test Explorer](https://marketplace.visualstudio.com/items?itemName=ethan-reesor.vscode-go-test-adapter)

```
ext install ethan-reesor.vscode-go-test-adapter
```

## debugging executable file

source: [bubblesorter/cmd/bubbleSorter/main.go](https://github.com/74th/vscode-debug-specs/blob/master/golang/bubblesorter/cmd/bubbleSorter/bubbleSorter.go)

### launch.json

menu:`Go: Launch package`

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}/bubblesorter/cmd/bubbleSorter"
    }
  ]
}
```

- `program` must be main package folder or \*.go file

### start debugging

▶︎ Launch Package

## debugging at local process

source: [bubblesorter/cmd/bubbleSorter/main.go](https://github.com/74th/vscode-debug-specs/blob/master/golang/bubblesorter/cmd/bubbleSorter/bubbleSorter.go)

### prepare

```sh
# build executable file
cd bubblesorter/cmd/bubblesorter
go build

# enable ptrace scope (for Linux)
echo 0 | sudo tee /proc/sys/kernel/yama/ptrace_scope
```

### start process

```sh
./bubblesorter -sleep 30 7 4 2 6 &

[1] 1859211
```

### edit launch.json

Add processId to launch.json.

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Attach local process",
      "type": "go",
      "request": "attach",
      "mode": "local",
      "processId": 1859211,
      "apiVersion": 2,
      "showLog": true
    }
  ]
}
```

### start debugging

▶︎ Attach local process

## debugging running remote process

### prepare

```sh
cd cmd/bubbleSorter/
go build

# enable ptrace scope (for Linux)
echo 0 | sudo tee /proc/sys/kernel/yama/ptrace_scope
```

### execute and dlv attach

```sh
cd cmd/bubbleSorter/

# runnning process
./bubbleSorter -sleep 30 &
PID=$!
dlv attach $PID ./bubbleSorter --headless --listen=0.0.0.0:2345 --log --api-version 2
```

### edit launch.json

Edit host to remote server address.

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Attach remote process",
      "type": "go",
      "request": "attach",
      "mode": "remote",
      "port": 2345,
      "host": "127.0.0.1",
      "apiVersion": 2,
      "showLog": true
    }
  ]
}
```

### start debugging

▶︎ Attach remote process
