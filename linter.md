## Использование и настройка линтера

Линтер - это статический анализатор кода. 
Мы используем [golangci-lint](https://github.com/golangci/golangci-lint) - это агрегатор линтеров.

### Зачем нужен линтер?
1. Линтер следит за качеством кода. Если вы пишете код, который не соответствует стандартам по стилю, линтер это заметит и укажет на ошибку;
Это позволит вам до ревью узнавать о том, как обрабатывать ошибки, как именовать переменные, как разбивать функции и т.п.;
2. Линтер может находить ошибки, которые вы сами не заметите. Например, игнорирование ошибки;
3. Прогон линтов - это этап, который в компаниях проводится до этапа ревью. Он позволяет сэкономить время ревьюера, т.к. забирает на себя часть рутинной работы.


### Установка и настройка

Этап установки описан [на странице линтера на GitHub](https://github.com/golangci/golangci-lint#install).
1. Выполняем команду
```sh
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.23.6
```
Для тех кто на Windows, изучаем [issue](https://github.com/golangci/golangci-lint/issues/830)   
Следуем рекомендациям от tpounds и качаем [архив для windows](https://github.com/golangci/golangci-lint/releases/tag/v1.23.6)

```
@fgagneten I think the easiest method right now is to just manually download one of the Windows archives from the releases page and extract the binary somewhere on the PATH. That said, I do think adding support for chocolatey.org or similar might make sense. Patches are welcome if someone with experience would like to add support.
```
2. В директорию с проектом добавляем файл [.golangci.yml](.golangci.yml) (точка в начале файла обязательна)
3. Убедитесь, что в $PATH добавлен путь до golangci-lint
4. Запускаем так
```sh
golangci-lint run ./...
```
или вот так (если сохранили файл с настройками для golangci-lint без точки в начале)
```sh
golangci-lint run -c golangci.yml ./...
```


Для Windows: запуск из терминала git-bash
```sh
golangci-lint.exe run ./...
```
Если папка с golangci-lint.exe не добавлен в $PATH, то нужно указать полный путь до golangci-lint.exe
```sh
$SOME_DIR/golangci-lint.exe run ./...
```

