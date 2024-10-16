
HTTP GET data:
curl -sS http://localhost:8080/clients | nix-shell -p jq --run jq

HTTP POST data:
curl -sS http://localhost:8080/clients/ -d '{
    "bookmark": false,
    "tags": "",
    "full_name": "Сорокин Егор Львович",
    "phone": "+7 (000) 000 0000",
    "address": "",
    "email": "none@none.ru",
    "date": "2024/08/28"
}'
