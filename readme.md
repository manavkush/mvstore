This is a content addressed storage that's using Golang.

*Features:*
- Having a inbuilt library for networking protocols, (TCP implemented currently)
- Using the library to build a store which allows to store the contents of the file in a content addressed manner.
- The file is split and encoded. We them partition the fileContents and store the contents seperately.
- Able to write, read, delete from the storage.

*Usage:*
- A makefile is present in the root directory with the commands present.
- Following targets are currently present:
    - build: builds the go project
    - run: runs the go project. Dependant on build target
    - test: tests the go project
