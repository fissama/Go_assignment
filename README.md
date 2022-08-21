# Go Assignment 2022
A assignment for research Goroutines, Channel, MongoDB
## Steps research:
- Read config file
- Connect with MongoDB Atlas
- Insert a number into a collection
- Insert large number into a collection
- Use go routine, channel and waitgroup to run 10 worker for insert data
- Add echo to expose a api with POST /assignment/insert method
## Get the code
Clone the repository:
```
git clone https://github.com/fissama/Go_assignment.git
```
## Build
From the command line:
```
go build main.go
```
## Api docs
- POST **/assignment/insert**
```
{
    "FileData": FileCsv.csv
}
```