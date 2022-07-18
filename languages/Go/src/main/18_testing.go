package main

/*
https://github.com/stretchr/testify
https://pkg.go.dev/github.com/stretchr/testify/assert
https://pkg.go.dev/github.com/stretchr/testify/suite

Тесты размещаются рядом с тестируемыми файлами и имеют суффикс _test.go.
Тесты могут также размещаться в директории с суффиксом _test, это помогает
избавиться от циклических импортов.
Вспомогательные для тестов файлы называют harness_test или common_test.
Файлы *_test.go не используются при компиляции приложения.

Тестовые функции именуются с префиксом Test*.
См. примеры в пакете foo: go test foo -count 1
-count 1 запустит тесты без кеширования либо можно очистить go test clear.

Тесты можно запускать по регулярным выражениям, например:
go test foo -run ^TestFoo

Проверить покрытие: go test foo -cover -coverprofile=coverage.out
Проанализировать: go tool cover -html=coverage.out
*/
