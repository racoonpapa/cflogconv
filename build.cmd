@echo off
go build -o build\conv.exe -ldflags "-s -w" conv.go