// https://go.dev/ref/mod#go-mod-file-require
package main

/*
Есть два способа установить внешнюю зависимость:
1. Вручную с помощью `go get github.com/username/packagename`
Скачанные пакеты будут помещены в `$GOPATH/src/username/packagename`.

2. Указав список в разделе require файла `go.mod`.
При компиляции модули будут скачены и сохранены в `$GOPATH/pkg/mod`.
Также будет создан файл `go.sum` с хеш-суммами скаченных модулей.

В Go используется семантическое версионирование.
https://semver.org/spec/v2.0.0.html

Когда необходимо подменить зависимость, можно воспользоваться replace в `go.mod`
replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5

Защита от проблем с зависимостями.
1. Вендоринг - хранение зависимостей в модуле, хранятся в директори vendor/.
Чтобы "завендорить" зависимости необходимо выполнить `go mod vendor`.
Чтобы использовать такие зависимости запуск требуется выполнять с флагом -mod=vendor.
2. Прокси-серверы модулей.
Указываются через переменную окружения GOPROXY:
GOPROXY=https://proxy.golang.org/,direct
Попытки скачать пакет выполняются по порядку, указанному через запятую.
direct - скачать пакет из оригинального репозитория.
GOPRIVATE=*.internal.company.com
В GOPRIVATE указывается для каких зависимостей необходимо
ходить непосредственно в репозитории.
*/

func main() {
	// fmt.Println("testify", mock.Call{})
}
