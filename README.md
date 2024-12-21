## Calculator on GO


### Правила ввода выражений для их вычисления

#### Доступные операнды:
1. Умножение "*"
2. Сложение "+"
3. Деление "/"
4. Вычитание "-"

#### Принцип работы:

Данный калькулятор обрабатывает выражения через алгоритм обратной польской записи [rpn][0] <br />

Сам ввод выражения происходит привычным способом `пробелов между операндами и выражениями нет`<br /> "(число/выражение в скобках)(операнд)(число/выражение в скобках)" <br />

Пример написания выражения:
1. "2+2"
2. "22/2"
3. "(34+6)/2"

#### Предупреждение
Калькулятор умеет работать с выражениями в скобках, но в некоторых случаях может некорректно выдавать ответ <br />

### Запуск

По умолчанию приложение запускается в режиме HTTP-сервиса:

```bash
go run ./cmd/calc_service/...
```
### Использование

на наш url ` /api/v1/calculate` отправляем POST-запрос
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "TYPE_HERE"}'
```
- На месте `TYPE_HERE` введите ваше выражение. 
В случае успешного ввода (т.е. если ваше выражение соответствует правилам выше) вам ретернется HTTP-ответ с телом:  
``` 
{"result":"SOME_RESULT"} 
```
На месте `SOME_RESULT` будет результат успешно выполненной работы

### Тестирование калькулятора
Юнит-тесты (простенькие) можно прогнать через команду 
```bash
go test -v ./test/... 
``` 
### Структура проекта

```
calcOnGo/
├── cmd/
│   └── calc_service/
│       └── main.go      # Точка входа в приложение
├── internal/
│   ├── handler/
│   │   └── calculate_handler.go  # Обработчик HTTP-запросов
│   ├── service/
│   │   └── calc.go  # Мозги
│   └── utils/
│       └── validator.go   # Валидность данных
├── pkg/
│   └── errors/
│       └── errors.go      # Ошибки
├── test/
│   └── unit/
│       └── calculator_test.go # Юнит-тесты
├── Dockerfile            # Docker-конфигурация (опционально)
├── go.mod                # Модуль Go
└── README.md             # Документация проекта
```

[0]: https://ru.wikipedia.org/wiki/Обратная_польская_запись "RPN"