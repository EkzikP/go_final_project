# Файлы для итогового задания

В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.

Директория `web` содержит файлы фронтенда.

Данный проект создан показать практическое применение знаний, полученных на курсе "GO-разработчик с нуля"

Проект реализует "Планировщик задач".
Планировщик хранит задачи, каждая из них содержит дату дедлайна и заголовок с комментарием.
Задачи могут повторяться по заданному правилу: например, ежегодно, через какое-то количество дней,
в определённые дни месяца или недели. Если отметить такую задачу как выполненную,
она переносится на следующую дату в соответствии с правилом.
Обычные задачи при выполнении будут просто удаляться.

Реализованы все задачи повышенной сложности, а именно:
1. Реализована возможность определять порт веб-сервера, через переменную TODO_PORT в .env файле
2. Реализована возможность изменять путь к файлу базы данных, через переменную TODO_DBFILE в .env файле
3. Добавлена возможность установить повтор задачи:
   а. в указанные дни недели
   б. в указанные дни месяца, последний или предпоследний день месяца.
4. Реализована возможность поиска задачи по вхождению строки поиска в Заголовке или комментарии,
   также добавлена возможность поиска задач по дате.
5. Реализована аутентификация пользователя на странице http://localhost:7540/login.html
   Сессия авторизации хранится в куки 8 часов.
6. Реализована возможность сборки и запуска докер образа.

Значения .env файла, по умолчанию:
TODO_PASSWORD=1278
TODO_DBFILE=scheduler.db
TODO_PORT=7540

Параметры файла ./tests/settings.go для запуска тестов:
var Port = 7540
var DBFile = "../scheduler.db"
var FullNextDate = true
var Search = true
var Token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJoYXNoIjoiYTFlYjdjZTFiMzEzMmQyNjlmMTVhY2Q5Yzk3OGM0ODFjOTM4OTUwNiJ9.z0Jyt1rBDYUdF4Igmzr-ENZJ0-WmOEM8hhx0Ou11KEA`

Переменные среды,
TODO_PASSWORD
TODO_DBFILE
TODO_PORT
для докер образа указываются в файле Dockerfile. 

Команда для сборки докер образа:
docker build --tag go_final_project:v1 .

Команда для запуска докер образа
docker run -d -p 7540:7540 -v //f/GO/Dev/go_final_project:/usr/src/app/base go_final_project:v1

На других операционных системах необходимо изменить путь "//f/GO/Dev/go_final_project"
на путь до места хранения базы данных, в зависимости от текущей операционной системы.