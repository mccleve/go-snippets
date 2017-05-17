# File Replace:


The goal of this script is to create a user friendly sed replacement for inline case insensitve string substitution.


## Build & Installation:


Requires go-lang 1.7+


build:

`make`


install:

`make install`


Installs to `/usr/local/bin`


## Usage:


```bash
$ ./file-replace
2017/05/17 17:07:03 Usage:
  -match string
    	String to match
  -path string
    	Path to scan for files
  -replace string
    	Replace matched string with this value

```

