Install on mac from brew ``` brew install dhnikolas/tools/cloner``` <br>

Put config in home directory ```~/.clonerconfig```<br>

```layouts, namespaces, projects_dir - optional parameters``` <br>
```default projects_dir = ~/mygo``` <br>
```default namespace, layouts = the values correspond to those in the example below```

```json
{
  "git": {
    "user": "nnaumchenko",
    "password": "somepassword"
  },
  "projects_dir": "/Users/dhnikolas/mygo",
  "namespaces": [
    "common-bank-services"
  ]
}
```
