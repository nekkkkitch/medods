# Installation and launching
1. Клонируйте проект куда-нибудь
   ```
   git clone https://github.com/nekkkkitch/medods
   ```
2. Перейдите в папку medods
   ```
   cd .../medods
   ```
3. Запустите в терминале следующую команду(Должен быть запущенный Docker engine и Makefile должен быть установлен)
   ```
   make buildbuilder
   ```
4. Запустите в терминале ещё одну команду(100% сначала контейнер с tokens остановится из-за того, что БД запускается довольно долго. Подождите и потом запустите контейнер с gateway)
   ```
   make start
   ```

# Дополнительно
В проекте лежит папка test - можно протестировать сервис, запуская различные go файлы оттуда
При инициализации в БД закидывается пользователь с login=123, который вы можете использовать в тестах для получения uuid

Если что-то не работает - напишите мне пожалуйста на телеграм @nekkkkitch
