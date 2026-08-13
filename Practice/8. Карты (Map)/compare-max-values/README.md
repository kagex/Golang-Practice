# Сравнение максимальных значений в срезах

Напишите функцию `CompareMaxValues`, которая принимает два параметра: два `map` типа `map[string][]int`.

Функция должна сравнить максимальные значения в срезах для каждого ключа. Если максимальные значения для всех соответствующих ключей в обоих `map` равны, функция должна вернуть `true`, в противном случае — `false`.

## Примечание

- Функция `CompareMaxValues` не должна ничего выводить, она должна вернуть `true` или `false`.
- Убедитесь, что ваша функция корректно обрабатывает пустые значения.

## Тестовые данные

| № Теста | Входные данные | Выходные данные |
| :---: | :--- | :---: |
| 1 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [0, 2, 3],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [4, 5, 6]`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [3, 1, 2],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [6, 5, 4]`<br>&nbsp;&nbsp;`}`<br>`}` | `true` |
| 2 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [1, 2, 3],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [4, 5, 6]`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [3, 1, 2],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [6, 5, 7]`<br>&nbsp;&nbsp;`}`<br>`}` | `false` |
| 3 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [1, 2, 3],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [4, 5, 6]`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [3, 1, 2]`<br>&nbsp;&nbsp;`}`<br>`}` | `false` |
| 4 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [1, 2, 3]`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [3, 2, 1]`<br>&nbsp;&nbsp;`}`<br>`}` | `true` |
| 5 | `{`<br>&nbsp;&nbsp;`"m1": {},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [1, 2, 3]`<br>&nbsp;&nbsp;`}`<br>`}` | `false` |
| 6 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [1, 2, 3]`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {}`<br>`}` | `false` |
| 7 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [-5, -3]`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [-5, -1]`<br>&nbsp;&nbsp;`}`<br>`}` | `false` |
| 8 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [-5, -3, 0]`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": []`<br>&nbsp;&nbsp;`}`<br>`}` | `false` |
| 9 | `{`<br>&nbsp;&nbsp;`"m1": {},`<br>&nbsp;&nbsp;`"m2": {}`<br>`}` | `true` |
| 10 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [1, 2, 3],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [4, 5]`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [3, 1],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [5, 4, 6]`<br>&nbsp;&nbsp;`}`<br>`}` | `false` |
| 11 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"x": [1, 2, 3]`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"x": [9, 9, 9]`<br>&nbsp;&nbsp;`}`<br>`}` | `false` |
| 12 | `{`<br>&nbsp;&nbsp;`"m1": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [1, 2, 3],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [4, 5, 6],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"c": [7, 8, 9]`<br>&nbsp;&nbsp;`},`<br>&nbsp;&nbsp;`"m2": {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"a": [3, 2, 1],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"b": [6, 5, 4],`<br>&nbsp;&nbsp;&nbsp;&nbsp;`"c": [99, 99, 99]`<br>&nbsp;&nbsp;`}`<br>`}` | `false` |
