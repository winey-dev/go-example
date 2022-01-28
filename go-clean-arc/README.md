
# Contents
[Learning Objectives](#-✔-Learning-Objectives)   
[Directory Description](#-✔-Directory-Description)


# ✔ Learning Objectives
- [x] User CRUD를 지원하는 서비스를 구현   
- [x] 별도 DB 연동 없이 Local Memory를 사용하여 DB 구현   
- [x] mockery를 통해 MockUserRepoInterface 생성
- [ ] mockRepository를 이용한 UserService Interface에 대한 Unitest 진행
- [ ] faker 사용을 통해서 Unitest 진행 시 Sample Data 제공 방법



# ✔ Directory Description
|Path        | 구현된 내용        |
|:-----------|:------------------|
|go-clean-arc|`project root path`|
|go-clean-arc/model/user.go|`User 구조체 정의`<br> `UserService Intreface 정의 및 구현` <br>`UserRepository Interface 정의`|
|go-clean-arc/model/user_test.go|`User Serivce에 대한 Unitest`|
|go-clean-arc/model/mocks/UserRepository.go|`User Repository Intreface에 대한 Mocking Repository 지원`|
|go-clean-arc/model/repo/user.go|`User Repository Interface 구현`|
