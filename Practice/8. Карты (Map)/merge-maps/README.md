# Объединение мап

Необходимо реализовать функцию `mergeMaps`, которая принимает два аргумента типа `map[string]int` и возвращает новый мап типа `map[string]int`. Функция должна объединять значения из двух мап, суммируя значения для одинаковых ключей. Если ключ присутствует только в одной из мап, он должен быть добавлен в результирующую мапу с соответствующим значением.

## Пример

```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"b": 3, "c": 4, "d": 5}

result := mergeMaps(m1, m2)
// result будет равен map[string]int{"a": 1, "b": 5, "c": 7, "d": 5}
```

## Примечание

- Внутри функции `mergeMaps` не нужно менять переданные в нее значения.

## Тестовые данные

| № Теста | Входные данные | Выходные данные |
| :---: | :--- | :--- |
| 1 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": 1,`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": 2,`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"c": 3`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": 3,`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"c": 4,`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"d": 5`<br>&nbsp;&nbsp;`}`<br>`}` | `{`<br>&nbsp;&nbsp;`"a": 1,`<br>&nbsp;&nbsp;`"b": 5,`<br>&nbsp;&nbsp;`"c": 7,`<br>&nbsp;&nbsp;`"d": 5`<br>`}` |
| 2 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": 1,`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": 2`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {}`<br>`}` | `{`<br>&nbsp;&nbsp;`"a": 1,`<br>&nbsp;&nbsp;`"b": 2`<br>`}` |
| 3 | `{`<br>&nbsp;&nbsp;`"m1": {},`<br>&nbsp;&nbsp;`"m2": {}`<br>`}` | `{}` |
| 4 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": 1,`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": 2`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"c": 3,`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"d": 4`<br>&nbsp;&nbsp;`}`<br>`}` | `{`<br>&nbsp;&nbsp;`"a": 1,`<br>&nbsp;&nbsp;`"b": 2,`<br>&nbsp;&nbsp;`"c": 3,`<br>&nbsp;&nbsp;`"d": 4`<br>`}` |
| 5 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": 1,`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": 2`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": null`<br>`}` | `{`<br>&nbsp;&nbsp;`"a": 1,`<br>&nbsp;&nbsp;`"b": 2`<br>`}` |
| 6 | `{`<br>&nbsp;&nbsp;`"m1": null,`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": 1,`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": 2`<br>&nbsp;&nbsp;`}`<br>`}` | `{`<br>&nbsp;&nbsp;`"a": 1,`<br>&nbsp;&nbsp;`"b": 2`<br>`}` |
