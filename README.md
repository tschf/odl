# odl (Oracle Download Utility)

Automating tests can be a pain, I'm hoping this utility will provide developers with a quick and easy way to fetch Oracle media files.

The goal isn't to bypass the OTN license agreement, or not to log in - the download will not work if you do not provide valid OTN authentication credentials. There are two mechanisms for the username. Pass in the flag: -username <username> to the program; set an environment variable, `OTN_USERNAME`. The password has three mechanisms. Pass in the flag -password; set an environment variable, `OTN_PASSWORD`; Enter the password at run time, when prompted.

Supported software:

| Description            | OS      | Architecture | Arguments |
|---                     | ---     | ---          | ---       |
| Oracle Database 11g XE | linux   | x64          | --component db --os linux --version 11gXE --architecture x64 |
| Oracle Database 11g XE | windows | x86          | --component db --os windows --version 11gXE --architecture x86 |
| Oracle Database 11g XE | windows | x64          | --component db --os windows --version 11gXE --architecture x64 |
| APEX 5.1               | na      | na           | --component apex --os na --version 5.1 --architecture na  |
| SQLcl                  | na      | na           | --component sqlcl --os na --version 4.2  |

Sample usage:

```bash
trent@birroth:/tmp/xe$ odl --help
Usage of odl:
  -arch string
    	Specify the desired architecture of the software. Should be "x86", "x64", or "na" (default na)
  -component string
    	Specify the component to grab. Should be "db" (default "db")
  -os string
    	Specify the desired platform of the software. Should be "linux" or "windows" (default "linux")
  -password string
    	Specify the password that corresponds to your OTN account. Alternatively, set the environment variable OTN_PASSWORD.
  -username string
    	Specify the user account that will be logging in and accepting the license agreement. Alternatively, set the environment variable OTN_USERNAME.
  -version string
    	Specify the software version. Should be "11gXE" (default "11gXE")
trent@birroth:/tmp/xe$ odl
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

Keep in mind, at this stage it is a very early prototype. It needs to be fleshed out a lot more. The next tasks on the pipeline are to modularise the program more, print the license agreement to the screen and then to include other installation media, with the next priority media files being APEX and SQLcl.
