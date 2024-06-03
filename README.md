# Desafio Clean Architecture

## Criação de usecase de listagem das orders

### A listagem das orders foi disponibilizadas por três tipos de serviços:
- Endpoint REST (GET /order)
- Query ListOrders GraphQL
- Service ListOrders com GRPC

### Instruções para subir a aplicação
1. Após _git clone_ executar o comando abaixo na raiz do projeto:   
    `docker compose up --build`
2. Quando aparecer a mensagem do RabbitMQ, se aparentar que houve um freeze no console, aguarde alguns segundos que 
provavelmente irão aparecer mensagens dos serviços(web server, gRPC e GraphQL) e o sistema estará pronto.
3.  As seções seguintes serão instruções para fazer o uso da listagem das orders para cada tipo de serviço.

### Endpoint REST (GET /order)
1. Usar a URL `http://localhost:8000/order` com method `GET`
2. Isto é suficiente. Irão aparecer registros que já foram inseridos quando a aplicação subiu. 
Isto vale para os próximos serviços.

### Query ListOrders GraphQL
1. Acesse pelo browser `http://localhost:8080`
2. Será apresentado a interface do _GraphiQL_
3. Copie toda a estrutura abaixo, cole no _GraphiQL_ e execute.
   ```
   query {
       ListOrders {
           id
           Price
           Tax
           FinalPrice
       }
   }
   ```
4. Após executar a query acima os dados irão aparecer na ferramenta.

### Service ListOrders com GRPC
1. Para usar este serviço a recomendação é usar o gRPC Client [Evans - https://github.com/ktr0731/evans](https://github.com/ktr0731/evans).
2. Use o comando `evans -r repl`
3. No client digite `call` e em seguida _digite um espaço_, deverá então aparecer `ListOrders`
4. Execute então o `ListOrders` e os registros irão aparecer.