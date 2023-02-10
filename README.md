# Simple Start page
simple home page that show list link or application.

## Prerequisites
* go 1.19
* nodejs 18

## How to build portable executable
1. clone this repo
2. go to web directory
   ```cd web```
3. install frontend dependency
   ```npm install```
4. build front end
   ```npm run build```
5. go back to project root directory
   ```cd ..```
6. download backend dependency
   ```go mod download```
7. compile the application
   ```go build -o ssp cmd/ssp.go```
8. run the application
   ```./ssp -db=<database path> -secret=<app secret>```

## Todo
* ~~frontend (vue)~~
* ~~serve as embed (use go embed)~~
* ~~implement front end~~
* ~~implement api~~
* ~~integrate front end + api~~
* ~~presistance storage (diskv)~~
* ~~setting  list link~~
* ~~get list application from docker~~
* proper ui (remove alert, better app list)
