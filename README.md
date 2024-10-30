# qdrantURItoClient

Хочется чтобы все было модно и молодежно и поэтому затаскиваем database url на территорию qdrant. Этот модуль содержит одну фукнцию UriToClient с одним параметром. Задача этой функции из строки "qdrant://api_key-11111@11111.europe-west3-0.gcp.cloud.qdrant.io:6334?UseTLS=1", получить просто указатель на обьект клиента. Вместо имени пользователя указываем api key . Порт по умолчанию 6333.

### install
go get github.com/saintbyte/qdrantURItoClient
Буду рад вашим пулл реквестам.
