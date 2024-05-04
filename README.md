# test-cache-module
Учебный модуль для хранения данных в кеше

## Установка:
Если вы не инициализировали модуль, необходимо выполнить команду

```bash
  go mod init <module_name> 
```

затем выполнить команду для получения модуля из GIT

```bash
go get github.com/RomanSkriabin/test-cache-module
```

После этого можно импортировать модуль в рабочий файл и использовать

### Инициализация кеша

```go
cache := cache.NewCache()
```

### Добавить значение в кеш или перезаписать значение
```go
cache.Set("user1", 20)
```

### Удалить значение из кеш
```go
cache.Delete("user2")
```

Получить значение из кеш
```go
gocache.Get("user2")
```


### example:

```go
package main

import (
	"fmt"

	cache "github.com/RomanSkriabin/test-cache-module"
)

func main() {
	fmt.Println("hi")
	cache := cache.NewCache()

	cache.Set("user1", 20)
	user1, _ := cache.Get("user1")

	cache.Set("user2", 30)
	user2, _ := cache.Get("user2")

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println("-------------------")

	cache.Set("user2", "40")
	user2, _ = cache.Get("user2")
	fmt.Println(user1)
	fmt.Println(user2)

	cache.Delete("user2")
	user2, _ = cache.Get("user2")
	fmt.Println("-------------------")
	fmt.Println(user1)
	fmt.Println(user2)


}

```
