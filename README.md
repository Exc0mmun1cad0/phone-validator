## Пример CI/CD для простого приложения

#### На push в dev (CI):
Сборка, тесты и 2 линтера

#### На push в main (CD):
1. Всё то же самое что и в dev
2. Сборка Docker-образа
3. Пушим его в ***реестр***
4. Деплой

## Описание работы приложения (задача с Yandex Cup направления Backend)
В спецификации указан регистрозависимый path. Сервис должен обрабатывать параметры query, которые задекларированы в спецификации, остальные параметры должны быть проигнорированы.
Запросы закодированы в версии HTTP/1.0. В ответ на запрос с неверным path возвращаем код 404. В ответ на запрос с неверными параметрами query возвращаем код 400. В остальных случаях возвращаем код 200 с детализированным ответом внутри.

Номера телефонов могут иметь форматы:

    +7 code ### ####
    +7 (code) ### ####
    +7code#######/code>
    8(code)###-####
    8code#######

Здесь # — цифра, <code — код оператора из списка [982, 986, 912, 934]
