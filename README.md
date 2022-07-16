# Install on Mac OS with brew
```
 brew install dhnikolas/tools/cloner
```

# Install on Linux
```
$ curl -s https://api.github.com/repos/dhnikolas/mtsbf-cloner/releases \
| grep "browser_download_url.*linux.tar.gz" | head -n1 \
| cut -d : -f 2,3 | tr -d \" | wget -qi - -O -| sudo tar -xz -C /usr/local/bin
 
$ sudo chmod +x  /usr/local/bin/cloner
```

# Configuration
Put config in home directory ```~/.clonerconfig.json```<br>

```layouts, namespaces, projects_dir - optional parameters``` <br>
```default projects_dir = ~/mygo``` <br>
```default namespace, layouts = the values correspond to those in the example below```

```json
{
  "git": {
    "user": "username",
    "password": "somepassword"
  },
  "projects_dir": "/Users/dhnikolas/mygo",
  "namespaces": [
    "core"
  ]
}
```
