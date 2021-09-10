# Balance Manager

## Запуск

Для работы микросервиса понадобится создать БД Postgresql с названием "balancemanager". Пользователь - "postgres", пароль "admin" (данные можно поменять в файле config/config.go, если ко времени проверки тестового я не успею поднять контейнер и вынести их в envs).

Запустить query tool, вставить содержимое файла database.sql, поднять базу.

Далее запускаем main.go. По умолчанию сервер поднимается на 0.0.0.0:8001

API хендлеров описан в файлe balanceManager.swagger.yml.

## Создание пользователя

Для создания нового пользователя отправить пустой get на 0.0.0.0:8001/createuser

В ответе получаем ID и пустой баланс

![Untitled](https://user-images.githubusercontent.com/71463390/132856031-a7c6087c-a1a0-4b56-879a-1e2a8fd61788.png)


## Пополнение счета

Для пополнения счета отправить post на 0.0.0.0:8001/deposit с body в json:

{"user_id":n, "amount":"n", "description":"n"}

![Untitled 1](https://user-images.githubusercontent.com/71463390/132856085-26a4b42b-be96-41e5-91d4-0a6be4498846.png)


## Списание средств

Для списания отправить post на 0.0.0.0:8001/withdrawal с body в json:

{"user_id":n, "amount":"n", "description":"n"}

![Untitled 2](https://user-images.githubusercontent.com/71463390/132856114-074a6244-4e91-4e21-b71a-7c0f588efd77.png)


## Перевод

Для перевода средств отправить post на 0.0.0.0:8001/sendfunds с body в json:

{"sender_id":n, "recipient_id":n, "amount":"n", "description":"n"}

![Untitled 3](https://user-images.githubusercontent.com/71463390/132856150-abf6e826-3d76-4b6b-bd57-433fc1c8b2cc.png)


## Список транзакций

Для получения списка транзакций отправить get на 0.0.0.0:8001/transactions с params:

user_id = n, limit = n, cursor = n

Для первого запроса cursor оставить пустым.

![Untitled 4](https://user-images.githubusercontent.com/71463390/132856189-4602af8e-5f5f-46aa-9a73-a1b5dfec7a6d.png)


Данные отсортированы DESC, следующие можно получить, указывая курсор из ответа сервера.

![Untitled 5](https://user-images.githubusercontent.com/71463390/132856206-2324adbc-fc4b-4d34-83a1-11b147a51dc1.png)


