# Работа с проектом

`vscode` при старте предложит установить рекомендованныек расширения - с ним лучше согласиться.


## Настройка окружения

### Golang

Установите `golang` для соответствующей операционной системы. 

Запустите выкачивание зависимостей

```shell
go mod tidy
```

Для локальной проверки линтером следует установить линтер. Устанавливать следует из git bash терминала в случае `windows`

```shell
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
```

Для запуска линтера следует выполнить следующую команду
```shell
$(go env GOPATH)/bin/golangci-lint run -vvv --out-format tab --print-resources-usage --timeout=180s --config=.golangci.yml
```

### Запуск тестов

```shell
go test -count=1 -p 1 -v ./...
```