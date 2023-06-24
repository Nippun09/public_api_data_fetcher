# Periodic Public API Data Fetcher

## General Info
 All source files are in "workspace" folder.
 Default configured api in this project is "https://api.chucknorris.io/jokes/random".
 Above api return new data every time it is called.
 In case wish to use other API, you will need to change "Data" struct accordingly.
 Default fetch period configured is 5000000000 i.e. 5s.
 Received Json results are stored in data.txt as configured.

## Run Command for Windows
 go run .\main.go
