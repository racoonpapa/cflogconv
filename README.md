# Cloudfront Log Converter

This program is a small converter utility for Cloudfront. I wrote this one for my personal purpose.

It was very annoying for me to check the logs while I was developing my app with Cloudfront. To check logs from Cloudfront, I had to download compressed files,  uncompress them, and open it with a text editor.

I know that I can use `Athena` or some other solutions to help logging and they give useful functions, but I don't need those functions.

This program is very simple as you can see the source code.

Just download your gzipped log files from S3 into one folder. And run this program with an argument to specify input directory. All .gz files will be converted into a single csv-comma seperated- file.

I cannot sure whether this program works for all conditions. Just use it freely for your purpose.

## Build

```
go build -o build\conv.exe -ldflags "-s -w" conv.go
```

## Usage
```
conv.exe -i <input directory> -o <output file. csv filename>
```
