# Desafio: Clean Architecture 

## Executando o Projeto

1. Clone o repositório para o seu ambiente local:
   ```sh
   git clone git@github.com:afonsojota/go-clean-architecture.git
    ```
2. Abra um terminal na raiz do projeto.
3. Execute o seguinte comando para construir e iniciar os serviços definidos no arquivo docker-compose:
   ```sh
   docker-compose build
   ```
4. Execute o seguinte comando para iniciar a aplicação
   ```sh
   docker-compose up
   ```
   
###  API HTTP/Web

- **Criar um novo pedido:**

  ```sh
  curl --location 'http://localhost:8000/order' \
  --header 'Content-Type: application/json' \
  --data '{
      "id": "9",
      "price": 50.5,
      "tax": 0.12
  }'
  ```

- **Listar pedidos:**

  ```sh
  curl --location 'http://localhost:8000/order'
  ```


###  GraphQL

- **Criar pedido:**

  ```graphql
  mutation createOrder {
    createOrder(input: {id: "8", Price: 50.5, Tax: 0.12}) {
      id
      Price
      Tax
      FinalPrice
    }
  }
  ```

- **Listar pedidos:**

  ```graphql
  query  {
    listOrders {
      id
      Price
      Tax
      FinalPrice
    }
  }
  ```

### GRPC 

**Criar pedido:**

      grpcurl -plaintext -d '{"id":"7","price": 100.5, "tax": 0.5}' localhost:50051 pb.OrderService/CreateOrder


**Listar pedidos:**

    grpcurl -plaintext -d '{}' localhost:50051 pb.OrderService/ListOrders
      
