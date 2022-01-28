# go-example
Golang을 사용하면서 Example한 소스코드 모음 집이다.

## <a href="https://github.com/yiaw/go-example/tree/main/go-clean-arc" target="_blank">1. Go Clean Architecure</a>
 `Go Clean Architecture`에 대해 궁금증이 생겼고 어떤식으로 구현되는지 파악하기 위해 Example을 만들었다.<br> 
 `Go Clean Architecture`를 연습하면서 평소 미루고 미루던 Unitest 방법에 대해서도 같이 학습 하였다.<br>

## <a href="https://github.com/yiaw/go-example/tree/main/gorm" target="_blank">2. gorm</a>
 C언어를 사용하다보니 ORM이라는 개념이 별로 중요하지 않다고 생각했지만 이번 Golang을 학습하면서 새로 ORM에 대해서 학습을 진행했다.<br>
 Golang에서 사용하는 ORM은 `gorm`으로 대강 동작 방법을 익히기 위해 Example을 만들었다.

## <a href="https://github.com/yiaw/go-example/tree/main/k8s-client/get-resource" target="_blank">3. k8s-client</a>
 Kubernetes에서 제공하는 `OpenAPI`를 통해서 `kube-apiserver`와 통신하여 원하는 값을 얻을 수는 있다. <br>
 REST API를 직접 사용하려면 까다로운 절차가 필요하다. <br>
 kubernetes에서 제공해주는 Client PKG인 `client-go`를 이용하여 `kube-apiserver`와 통신 하는 방버에 대해서 알아 본다.

## <a href="https://github.com/yiaw/go-example/tree/main/k8s-some" target="_blank">4. k8s-some</a>
 `kubernetes-dashboard`, `okd`, `octant`, `data-dog`등 kubernetes 진영에서 사용되는 `Monitoring Solution` 에서는 직접 Container에 접속할 수 있는 화면을 제공 한다. <br>
 어떤 방법으로 구현되어 있는지 서버 입장에서 동작 방식을 파악해 본다.
