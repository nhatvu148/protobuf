# Reference:

- https://developers.google.com/protocol-buffers/docs/proto3

# Setup protoc:

- Mac: brew install protobuf
- Windows:

```
Download the windows archive: https://github.com/google/protobuf/releases

Example: https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-win32.zip

Extract all to C:\proto3

Your directory structure should now be

C:\proto3\bin

C:\proto3\include

Finally, add C:\proto3\bin to your PATH:

From the desktop, right-click My Computer and click Properties.
In the System Properties window, click on the Advanced tab
In the Advanced section, click the Environment Variables button.
Finally, in the Environment Variables window (as shown below), highlight the Path variable in the Systems Variable section and click the Edit button. Add or modify the path lines with the paths you wish the computer to access. Each different directory is separated with a semicolon as shown below.

C:\Program Files; C:\Winnt; ...... ; C:\proto3\bin
(you need to add ; C:\proto3\bin  at the end)
```

# Commands:

- protoc -I=proto --python_out=python proto/\*.proto
- protoc -I=proto --java_out=java proto/\*.proto
- protoc -I=proto --js_out=javascript proto/\*.proto
- protoc -I=proto --csharp_out=csharp proto/\*.proto
- protoc -I=proto --cpp_out=cpp proto/\*.proto

# Git default name:

- hint: Using 'master' as the name for the initial branch. This default branch
  name
- hint: is subject to change. To configure the initial branch name to use in all
- hint: of your new repositories, which will suppress this warning, call:
- hint:
- hint: git config --global init.defaultBranch <name>
- hint:
- hint: Names commonly chosen instead of 'master' are 'main', 'trunk' and
- hint: 'development'. The just-created branch can be renamed via this command:
- hint:
- hint: git branch -m <name>

# Golang:

- go get -u github.com/golang/protobuf/protoc-gen-go
- go get -u github.com/golang/protobuf/proto
- Add to proto file an output folder: option go_package = "./simple";

# Import local modules:

```
In async folder: go mod init github.com/nhatvu148/go-lang/async
Enter hackerrank folder: go mod edit -replace=github.com/nhatvu148/go-lang/async=../async
go get github.com/nhatvu148/go-lang/async
Reference: https://golang.org/doc/tutorial/call-module-code
```
