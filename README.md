# odl (Oracle Download Utility)

**Program Status:** Working, but Oracle SSO had a couple of changes in the last few weeks. Monitoring. See issue #28

[![build status](https://travis-ci.org/tschf/odl.svg?branch=master)](https://travis-ci.org/tschf/odl)

Automating tests can be a pain, I'm hoping this utility will provide developers with a quick and easy way to fetch Oracle media files.

The goal isn't to bypass the OTN license agreement, or not to log in - the download will not work if you do not provide valid OTN authentication credentials. There are two mechanisms for the username. Pass in the flag: -username &lt;username&gt; to the program; set an environment variable, `OTN_USERNAME`. The password has three mechanisms. Pass in the flag -password; set an environment variable, `OTN_PASSWORD`; Enter the password at run time, when prompted.

## Install

To grab the latest build, assuming you have `golang` installed, run:

```
go get github.com/tschf/odl
```

Then, so long as `$GOPATH/bin` is in your PATH, you should be able to run `odl`.

Or, grab the latest release: [https://github.com/tschf/odl/releases](https://github.com/tschf/odl/releases), which includes binaries for OSX, Windows and Linux, for 32 and 64-bit architecture's. Note, the compiled binaries may not include the latest features from source.

## Supported software

| Component                  | Version       | OS                           | Arch    | Lang  |
| ---                        | ---           | ---                          | ---     | ---   |
| apex                       | 4.2 - 5.1     | na                           | na      | na,en |
| db                         | 11gXE         | windows                      | x64     | na    |
| db                         | 11gXE         | linux                        | x86,x64 | na    |
| db                         | 12.1.0.2.0EE  | linux,windows                | x64     | na    |
| db                         | 12.1.0.2.0SE2 | linux,windows                | x64     | na    |
| db                         | 12.2.0.1.0    | linux                        | x64     | na    |
| instantclient-basic        | 12.2.0.1.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-basic        | 12.1.0.2.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-basic        | 11.2.0.4.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-basic        | 11.1.0.7.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-basic-lite   | 12.2.0.1.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-basic-lite   | 12.1.0.2.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-basic-lite   | 11.2.0.4.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-basic-lite   | 11.1.0.7.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-jdbc         | 12.2.0.1.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-jdbc         | 12.1.0.2.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-jdbc         | 11.2.0.4.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-jdbc         | 11.1.0.7.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-odbc         | 12.2.0.1.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-odbc         | 12.1.0.2.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-odbc         | 11.2.0.4.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-odbc         | 11.1.0.7.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-precompiler  | 12.2.0.1.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-precompiler  | 12.1.0.2.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-precompiler  | 11.2.0.4.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-sdk          | 12.2.0.1.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-sdk          | 12.1.0.2.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-sdk          | 11.2.0.4.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-sdk          | 11.1.0.7.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-sqlplus      | 12.2.0.1.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-sqlplus      | 12.1.0.2.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-sqlplus      | 11.2.0.4.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-sqlplus      | 11.1.0.7.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-wrc          | 12.2.0.1.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| instantclient-wrc          | 12.1.0.2.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-wrc          | 11.2.0.4.0    | linux,linux-rpm,windows,osx  | x64,x86 | na    |
| instantclient-wrc          | 11.1.0.7.0    | linux,linux-rpm,windows      | x64,x86 | na    |
| java-jdk                   | 8             | linux,windows                | x64,x86 | na    |
| java-jdk                   | 8             | osx                          | x64     | na    |
| java-jre                   | 8             | linux,windows                | x64,x86 | na    |
| java-jre                   | 8             | osx                          | x64     | na    |
| ords                       | 3.0 - 3.0.9   | na                           | na      | na    |
| sqlcl                      | 17.2          | na                           | na      | na    |
| sqlcl                      | 4.2           | na                           | na      | na    |
| sqldev                     | 4.1           | linux,windows,osx            | na      | na    |
| sqldev-jdk                 | 4.1           | windows                      | x64     | na    |
| sqldev                     | 4.2           | linux,windows,osx            | na      | na    |
| sqldev-jdk                 | 4.2           | windows                      | x64     | na    |

## Usage

Each column in the table above represents an argument. e.g. to download Oracle 11gXE, you would run:

```bash
odl --component db --version 11gXE --os linux --arch x64
```

notes:

* Instant client 11.1.0.7.0 not available for OS X
* No ODBC for 11.2.0.4.0 on OS X
* No ODBC for Linux x64 on 11.1.0.7.0
* No JDK/JRE for OSX 32-bit

```bash
trent@birroth:/tmp/xe$ odl --help
Usage of odl:
  -accept-license
    	Specify whether or not you accept the OTN license agreement for the nominated software.
  -arch value
    	Specify the desired architecture of the software. Should be "x86", "x64", or "na" (default na)
  -component string
    	Specify the component to grab.
  -lang string
    	Specify the language of the software. Should be "en" or "na" (default "na")
  -os string
    	Specify the desired platform of the software. Should be "linux" or "windows" (default "linux")
  -password string
    	Specify the password that corresponds to your OTN account. Alternatively, set the environment variable OTN_PASSWORD.
  -username string
    	Specify the user account that will be logging in and accepting the license agreement. Alternatively, set the environment variable OTN_USERNAME.
  -version string
    	Specify the software version.
trent@birroth:/tmp/xe$ odl --component db --os linux --version 11gXE --arch x64 --accept-license
Beginning download process for db 11gXE
oracle-xe-11.2.0-1.0.x86_64.rpm.zip: 301.26 MB / 301.26 MB [==============================] 100.00% 5m30s
Download complete.
```

## Installing golang on Linux

To install golang as part of your build process you can use the following snippet:
```bash
#!/bin/bash
set -ev

GO_FILE=go1.8.3.linux-amd64.tar.gz

# https://golang.org/doc/install
curl -O "https://storage.googleapis.com/golang/${GO_FILE}"
sudo tar -C /usr/local -xzf ${GO_FILE}
rm ${GO_FILE}
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:${GOPATH}/bin
```
