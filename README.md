# emailrecap

This program is a simple tool that generates a pdf file with the provided email address in order to be printed as a summary for the patient.

## build

 ``` bash
GOOS=windows GOARCH=amd64 go build
 ```

## usage

``` bash
./emailrecap jeantest@example.com
```

### output

``` bash
$ ls
emailrecap.pdf  go.mod  go.sum  logo.png  main.go  emailrecap  README.md
```

`emailrecap.pdf` is the document ready to be printed.