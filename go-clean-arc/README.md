# go-clean-arc

# /model
전체적인 Model을 정의
# /model/user.go
user.go는 user 구조체 정의

# /model/usecase.go
usercase.go  user에 대한 usecase Interface 구현   
usecaseInterface는 RepoInterface를 입력받아 생성하도록 구현(의존성 주입)   

# /model/repo.go
Usecase에서 사용할 Repo관련 Interface만 정의

# /repoim
실제 Repo Interface 구현체    
Repo는 Local Memory로 동작하도록 구현
