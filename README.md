# Go First APP

## Para executar
- O sistema está usando JSON Web Token para autenticação. Para ganhar agilidade não implementei a geração. Utilizei o AWS Cognito.
-Usuário e senha enviei por email

### Setup Server
- A API está na pasta **server**
- O banco de dados é MongoDB.
    - No arquivo ./infra/config.go estão os parâmetros da base.
    - Para uma carga inicial de Usuários executar o comando:
    ```
    mongoimport --db [nomebanco] --collection users --file [caminho git local]]\data\users.json  --jsonArray
    ```
    - Para Widgets não é necessário carga

- No arquivo ./infra/config.go também é possível configurar a porta da API.
- Separei a API do front end. A execução deve ser de forma separada em diferentes hosts. O CORS já está configurado.

### Setup Client
- Para executar o client é necessário configurar o endereço da API, que se encontra no arquivo *client/app/app.js* variável *$rootScope.baseServiceUrl*
