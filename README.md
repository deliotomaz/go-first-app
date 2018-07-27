# Go First APP

## Para executar
- O sistema está usando JSON Web Token para autenticação. Para ganhar agilidade não implementei a geração. Utilizei o AWS Cognito.
-Usuário e senha enviei por email

### Setup Server
- A API está na pasta **server/api**. Nesta pasta que se encontra o *main*
- O banco de dados é MongoDB.
    - No arquivo **./infra/config.go** estão os parâmetros da base.
    - Para uma carga inicial de Usuários executar o comando:
    ```
    mongoimport --db [nomebanco] --collection users --file [caminho git local]]\data\users.json  --jsonArray
    ```
    - Para Widgets não é necessário carga

- No arquivo **./infra/config.go** também é possível configurar a porta da API.
- Separei a API do front end. A execução deve ser de forma separada em diferentes hosts. O CORS já está configurado.

### Setup Client
- Para executar o client é necessário configurar o endereço da API, que se encontra no arquivo *client/app/app.js* variável *$rootScope.baseServiceUrl*
- O Client deve rodar em um host separado

## Estrutura da API
- Como se trata de um teste implementei uma arquitetura de monolito. Em uma aplicação real consideraria utilizar microservices publicados em Docker ou Servless App.
- Devido ao tempo não consegui implementar testes na aplicação, mas sua estrutura permite testes em todas as camadas.
- Estrutura
    - Pasta **api**:
        - **main.go**: initializer da aplicação. Cria as interfaces de services e repositório
        -   **controller.go**: contém as controllers que recebem os requests os endpoints.
     - Pasta **domain**: 
        - **models.go**:Contém as models e declaração dos repositórios
      - Pasta **infra**:
        - **config.go**:Contém as configurações da api
        - **repository.go**:Contém as implementações dos repositórios da domain
     - Pasta **service**:
        - **interface.go**:Contém as declarações das interfaces (use cases) de AppService
        - **service.go**:Implementação dos use cases
- 

## Considerações

Entendo que o teste é simples e uma boa forma de avaliação. Contudo em uma aplicação real eu faria algumas coisas de forma diferente (não apliquei para não descaracterizar o que foi pedido):
 - Para listagem de usuários e widgets retornaria os dados paginados, assim como sua apresentação.
 - Nas AppServices faria validação das informações enviadas (tipo, obrigatoriedade, etc)
 - A busca não faria no client.
 - Criaria um objeto de resposta padrão que conteria os resultados e informações do response para que um middleware no client possa fazer alguns tratamentos padrão.
  - Como GO é muito novo pra mim, tentei ao máximo implementar uma estrutura consistente e desacoplada (minha preocupação constante), mas devido ao pouco tempo de contato com a linguagem sei que alguns pontos poderiam ser melhores. O fator tempo foi um complicador pra mim.
 - GO é fantástico e obrigado pela oportunidade de aprender mais!!!