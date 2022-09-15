# go-workspaces
Workspaces in Go

## Монорепозиторий + поддержка IDE
Данный пример по шагам показывает использование Go Workpaces в монорепозиториях

### Шаги реализации

- Убедиться, что изначально все работает
```shell
cd library/client
go run main.go
```

- Создать `go.work` файл в корне репозитория
```shell
cd ../..
go work init ./library ./library/client
```

- Добавить новую функцию в `library/library.go`
```go
func Bye(name string) {
	logrus.Infof("Bye, %s!\n", name)
}
```

- Добавить вызов функции в `library/client/main.go`
```go
	library.Bye("Boba")
```

- Запустить `client` из корня репозитория, убедиться, что изменения отображаются
```shell
go run library/client/main.go
```

- Опубликовать новую версию library
```shell
git add library/library.go
git commit -q -m "Add Bye function"
git push -q
git tag library/v0.0.2
git push -q origin library/v0.0.2
```

- Поменять версию зависимости в `library/client/go.mod` на `library/v0.0.2`
- Проверить, как работает client вне режима рабочей области
```shell
cd library/client
GOWORK=off go run main.go
```

- Опубликовать новую версию `client`
```shell
git add library/client
git commit -q -am "Add Bye function"
git push -q
git tag library/client/v0.1.2
git push -q origin library/client/v0.1.2
```

- Вы замечательны!

