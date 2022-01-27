# k8s-some
k8s go-client를 이용하여 Terminal 접속 및 Container의 로그를 Tailing 하는 방법에 대해서 설명 한다.

# Help
``` bash
:> .\k8s-some.exe --help
Usage of C:\Users\lsmin\Desktop\yiaw\Go\src\go-example\k8s-some\k8s-some.exe:
  -cn string
        --cn=ingress-controller (ContainerName)
  -ns string
        --ns=nginx-ingress (NameSpace) (default "nginx-ingress")
  -pn string
        --pn=ingress-controller-57b6745d-j89fm (PodName) (default "ingress-controller-57b6745d-j89fm")
  -type string
        --type=terminal or --type=log (default "terminal")
```
# Log Usage
``` bash
:> ./k8s-some.exe --type=log --ns=default --pn=ingress-controller-57b6745d-j89fm
```

# Terminal Shell
``` bash
:> ./k8s-some.exe --type=terminal --ns=default --pn=ingress-controller-57b6745d-j89fm
```