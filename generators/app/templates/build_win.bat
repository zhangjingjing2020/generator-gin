set GOOS=windows
set GOARCH=386
set CGO_ENABLED=0
go build -ldflags="-s -w" -o api-service32.exe main.go 
pause