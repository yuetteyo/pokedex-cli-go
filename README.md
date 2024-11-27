## POKEDEX

### Описание
Покедекс — это электронный девайс, являющийся каталогом информации о различных видах покемонов. Покедекс представлен в играх, аниме и манге покемон. Названия покедекса является игрой слов: часть от слова POKEmon и слова inDEX. В японском Покедекс просто переводится, как "Энциклопедия о Покемонах"

Данный проект реализует простое консольное приложение покедекса,  написанное на языке программирования Go с реализованной средой REPL и основанное на API PokéAPI (pokeapi.co), для получения данных с сервера.


### Задачи
- Парсинг JSON в Go
- Выполнение HTTP-запросов в Go
- Создание CLI-инструмента для взаимодействия с back-end API
- Локальная разработка на Go и настройка инструментов
- Кэширование для улучшения производительности

### Запуск и работа с программой
Cоздайте исполняемый файл с помощью команды go build:
```bash
go build
```

Запустите испольняемый файл:

```bash
.\pokedex-cli-go
```

Чтобы увидеть список доступных команд, введите help:

```
Type in a command > help
Welcome to the Pokedex help menu
Here are your available commands
 - exit: Turns off the Pokedex
 - help: Prints the help menu
 - map: Lists some location areas
 - mapb: Lists previous location areas
 - explore {location_area}: Lists the pokemons in the location area
 - catch {pokemon_name}: Attempts to a pokemon an add it to your pokedex
 - inspect {pokemon_name}: Shows information about a caught pokemon 
 - pokedex: Lists all caught pokemons in Pokedex

Type in a command > 
```