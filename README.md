
```Put config in home directory ~/.clonerconfig ```<br>
```layouts, namespaces, projects_dir - optional parameters``` <br>
```default projects_dir = ~/mygo``` <br>
```default namespace, layouts = the values correspond to those in the example below```

```json
{
  "git": {
    "user": "nnaumchenko",
    "password": "somepassword"
  },
  "layouts": {
    "list": [
      {
        "name": "layout-grpc",
        "namespace": "examples",
        "url": "https://qcm-git.mbrd.ru/service-platform/examples/layout-grpc",
        "description": "Grpc MTSBF template"
      },
      {
        "name": "layout-http",
        "namespace": "examples",
        "url": "https://qcm-git.mbrd.ru/service-platform/examples/layout-http",
        "description": "Http MTSBF template"
      },
      {
        "name": "layout-pub-sub",
        "namespace": "examples",
        "url": "https://qcm-git.mbrd.ru/service-platform/examples/layout-pub-sub",
        "description": "Pub-sub MTSBF template"
      }
    ]
  },
  "projects_dir": "/Users/dhnikolas/mygo",
  "namespaces": [
    "common-bank-services"
  ]
}
```