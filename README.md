# Cipher

my cipher program

## DISCLAIMER: 

This program is a _TEST_ and _TEST_ only, it is not a _STRONG_ cipher, and should _NOT_ be used when encrypting important data

## Installing

* firstly, we need to get the software, you can do this by entering `go get -d github.com/liamnaddell/cipher` 
* we change directories to the source directory with `cd $(GOPATH)/src/github.com/liamnaddell/cipher`
* we get the official golang dependancy management tool with `go get -u github.com/golang/dep/cmd/dep`
* Compile the software with `make`
* Install it to `/usr/local/bin` with `sudo make install`
* Profit


## Usage

```
NAME:
   cipher - Liam Naddell's basic encryption program

USAGE:
   cipher [global options] command [command options] [arguments...]

COMMANDS:
     genkey, g, -g  generate a key and store it in a file or print to stdout
     help, h        Shows a list of commands or help for one command
   Text Options:
     encode, e, -e  encode a string
     decode, d, -d  decode a encrypted string

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   ©️ Liam Naddell github.com/liamnaddell/cipher
```

### Example

This example will generate an encrypted message, then show you how to decrypt it.

```
mkdir test_keys
cd test_keys
cipher genkey -k mykey.lek
cipher encode "my (hopefully not important) message is this text, this text is self aware" -k mykey.lek
cipher decode $(whatever was output by the last command) -k mykey.lek
```

## Uninstalling

`sudo make uninstall`

Please if you could be so kind, tell me if there are issues with the software thus causing you to uninstall it

## Fin

If you wish to know anything else about this software, please leave an issue about what needs to be explained, who knows, maybe it will be implemented into these docs in the future.
