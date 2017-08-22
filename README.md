# Cipher

my cipher program

## DISCLAIMER: 

This program is a *TEST* and *TEST* only, it is not a *STRONG* cipher, and should *NOT* be used when encrypting important data

## Usage

```
NAME:
   cipher - Liam Naddell's terrible encryption program

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

