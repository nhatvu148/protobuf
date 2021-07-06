# Commands:

- protoc -I=proto --python_out=python proto/*.proto
- protoc -I=proto --java_out=java proto/*.proto
- protoc -I=proto --js_out=javascript proto/*.proto
- protoc -I=proto --csharp_out=csharp proto/*.proto
- protoc -I=proto --cpp_out=cpp proto/*.proto

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
