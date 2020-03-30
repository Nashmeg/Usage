# Usage
A service that gives the disk usage of the files in a given directory
- Query -> Files

## Test
- open main_test.go, click on "run package tests"

## Different Ways To Run
Clone the project on your local machine
- Run "go run ./cmd/usage/main.go" in vscode's terminal
OR 
- Click on the Debug button in vscode
OR
- Build and then run the image -> "docker build -t usage ." and then 
  "docker run -e PORT=80 -p 8082:80 usages"

## See the result
- In insomnia open a POST reest with "localhost:8082/api/usage"
OR 
- In a web browser open "http://localhost:8082/api/usage/playground"
Body:
```
  {
     files(where:{mountPoint:"/tmp"}) {
       name
       usage
     }
  }
```
OR
- Just health check with "http://localhost:8082/api/usage/health"

 







