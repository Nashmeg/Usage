# Usage
A service that gives the disk usage of the files in a given directory
- Query -> Files

## Test
- open main_test.go, click on "run package tests"
- After cloning the project on your local machine -> "go run ./cmd/usage/main.go" in    
  vscode's terminal
- OR Click on the Debug button in vscode
- OR Build and then run the image -> "docker build -t usage ." and then 
  "docker run -e PORT=80 -p 8082:80 usages"

## Different Ways To Run
Cloning the project on your local machine
- "go run ./cmd/usage/main.go" in vscode's terminal
> in insomnia open a POST reest with localhost:8082/api/usage and the body:
> - {
> -    files(where:{mountPoint:"/tmp"}) {
> -      name
> -      usage
> -    }
> - }





