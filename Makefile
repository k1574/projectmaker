all:
	GOOS=linux GOARCH=amd64
	go build -o ./projmaker src/advanced.go src/common.go src/create_idea.go src/editor.go src/index.go src/main.go src/login.go src/signUp.go src/profile.go

win:
	GOOS=windows GOARCH=amd64
	go build -o ./projmaker.exe src/advanced.go src/common.go src/create_idea.go src/editor.go src/index.go src/main.go src/login.go src/signUp.go src/profile.go
	