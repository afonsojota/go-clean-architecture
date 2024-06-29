# Go Expert - Clean Architecture

## Execução

1. Executar: `docker-compose up -d`;

2. Acessar a pasta _cmd/ordersystem_;

3. Executar: `go run main.go wire_gen.go`;

#### Via Web Server

4. Acessar a pasta _api_ e o arquivo _create_order.http_;

5. Enviar um _POST request_ para criar uma ordem;

#### Via gRPC

6. Abrir um novo terminal;

7. Executar: `evans -r repl`;

8. Executar: `call CreateOrder`;

9. Preencher os valores: _id_: asdfghjklxyz / _price_: 99.9 / _tax_: 9.0;

#### Via GraphQL

10. Acessar `localhost:8080`;

11. Executar:

```
mutation createOrder {
  createOrder(input: {id: "justsomeid", Price: 80.0, Tax: 2.0}) {
    id,
    Price,
    Tax,
    FinalPrice
  }
}
```

#### RabbitMQ

12. Criar uma fila e vinculá-la ao _exchange amq.direct_. Assim que a mensagem for enviada ao _exchange_, será redirecionada para a fila;

13. Acessar `localhost:15672`;

14. Acessar aba _Queues_ / _Add a new queue_ / _Name: orders_ / clicar _Add queue_;

15. Acessar a fila _orders_ / _Add binding to this queue_ / _From exchange: amq.direct_ / clicar _Bind_;

16. Em _Get messages_ / clicar _Get Message(s)_.