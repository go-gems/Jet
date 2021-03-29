# Jet

[![Go](https://github.com/go-gems/Jet/actions/workflows/go.yml/badge.svg)](https://github.com/go-gems/Jet/actions/workflows/go.yml)

To launch : 

```
go run .
```

By default the storage is in memory, and the generation is uuid based.


In order to change this, on launch add the following tags : 
```
  -generator.inc
        Use increment to generate keys
  -generator.short int
        Specify the number of characters to start with short generation
  -storage.dir string
        Specify directory to use Storage Directory
```
