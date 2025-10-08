Client (может быть несколько) получет данные от клиента в формате JSON (реализовал по REST API), делает простейшую валидацию и проксирует данные в сервис storage для получаения или изменения информации

Endpoints: \n
GET {url}/data?key={key} \n
POST {url}/data with body like {"key": "user", "value": "Roman"}
