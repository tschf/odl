# odl (Oracle Download Utility)

![build status](https://travis-ci.org/tschf/odl.svg?branch=master)

Automating tests can be a pain, I'm hoping this utility will provide developers with a quick and easy way to fetch Oracle media files.

The goal isn't to bypass the OTN license agreement, or not to log in - the download will not work if you do not provide valid OTN authentication credentials. There are two mechanisms for the username. Pass in the flag: -username <username> to the program; set an environment variable, `OTN_USERNAME`. The password has three mechanisms. Pass in the flag -password; set an environment variable, `OTN_PASSWORD`; Enter the password at run time, when prompted.

Supported software:

| Description                | OS                | Architecture | Language | Version                          | Arguments |
|---                         | ---               | ---          | ---      | ---                              | ---       |
| Oracle Database 11g XE     | linux             | x64          | na       | 11gXE                            | --component db --os linux --version 11gXE --arch x64 |
| Oracle Database 11g XE     | windows           | x86,x64      | na       | 11gXE                            | --component db --os windows --version 11gXE --arch x86 |
| APEX 5.1                   | na                | na           | en,na    | 4.2,5.0,5.1                      | --component apex --os na --version 5.1 --arch na --lang na |
| SQLcl                      | na                | na           | na       | 4.2                              | --component sqlcl --os na --version 4.2 |
| SQL Developer              | windows,linux     | na           | na       | 4.1                              | --component sqldev --os na --version 4.1 |
| SQL Developer (incl JDK)   | windows           | x64          | na       | 4.1                              | --component sqldev-jdk --os windows --version 4.1 --arch x64 |
| Instant client (basic)     | windows,linux,osx | x86,x64      | na       | 12.1.0.2.0,11.2.0.4.0,11.1.0.7.0 | --component instantclient-basic --os windows --version 12.1.0.2.0 --arch x64 |
| Instant client (basic lite)| windows,linux,osx | x86,x64      | na       | 12.1.0.2.0,11.2.0.4.0,11.1.0.7.0 | --component instantclient-basic-lite --os windows --version 12.1.0.2.0 --arch x64 |
| Instant client (jdbc)      | windows,linux,osx | x86,x64      | na       | 12.1.0.2.0,11.2.0.4.0,11.1.0.7.0 | --component instantclient-jdbc --os windows --version 12.1.0.2.0 --arch x64 |
| Instant client (sqlplus)   | windows,linux,osx | x86,x64      | na       | 12.1.0.2.0,11.2.0.4.0,11.1.0.7.0 | --component instantclient-sqlplus --os windows --version 12.1.0.2.0 --arch x64 |
| Instant client (sdk)       | windows,linux,osx | x86,x64      | na       | 12.1.0.2.0,11.2.0.4.0,11.1.0.7.0 | --component instantclient-sdk --os windows --version 12.1.0.2.0 --arch x64 |
| Instant client (odbc)      | windows,linux,osx | x86,x64      | na       | 12.1.0.2.0,11.2.0.4.0,11.1.0.7.0 | --component instantclient-odbc --os osx --version 12.1.0.2.0 --arch x64 |
| Instant client (wrc)       | windows,linux,osx | x86,x64      | na       | 12.1.0.2.0,11.2.0.4.0,11.1.0.7.0 | --component instantclient-wrc --os osx --version 12.1.0.2.0 --arch x64   |
| Java JDK,JRE               | windows,linux,osx | x86,x64      | na       | 8                                | --component java-jdk --os osx --version 8 --arch x64   |


notes:

* Instant client 11.1.0.7.0 not available for OS X
* No ODBC for 11.2.0.4.0 on OS X
* No ODBC for Linux x64 on 11.1.0.7.0
* No JDK/JRE for OSX 32-bit

Sample usage:

```bash
trent@birroth:/tmp/xe$ odl --help
Usage of odl:
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
trent@birroth:/tmp/xe$ odl --component db --os linux --version 11gXE --arch x64
demo.user@gmail.com
Do you accept the XE license agreement?
Full terms found here: http://www.oracle.com/technetwork/licenses/database-11g-express-license-459621.html
Enter Y for Yes, or N for No: Y
Enter your OTN password (demo.user@gmail.com):
The file being requested is https://edelivery.oracle.com/akam/otn/linux/oracle11g/xe/oracle-xe-11.2.0-1.0.x86_64.rpm.zip
Download complete.
trent@birroth:/tmp/xe$ ls -Alth
total 302M
-rw-rw-r-- 1 trent trent 302M Dec 16 16:25 oracle-xe-11.2.0-1.0.x86_64.rpm.zip
```
