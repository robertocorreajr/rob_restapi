# Goland REST API com Mux
Olá está é a minha primeira aplicação em _Golang_ e calhou de ser uma aplicação REST.

![postman](https://i.imgur.com/MRV5BoG.png)

## E uma API de livros
Como é um exemplo bem simples eu resolvi "soldar" no código os 2 primeiros livros inicias, mas eles podem ser apagados utilizando a API. Mas é claro que se reiniciar o servidor web eles voltarão. :)  


### O que podemos fazer:  
| URL | Metodo | Descrição |
|---|---|---|
| "/api/books" | GET | Listar todos os livros |
| "/api/books/{id}" | GET | Listar o livro do id |
| "/api/books" | POST | Criar um livro |
| "/api/books/{id}" | PUT | Atualizar um livro |
| "/api/books/{id}" | DELETE | Deletar um livro |

> OBS: Se alguém está vendo essa documentação e sabe uma forma melhor de documentar (dicas) por favor me diga. Ainda sou novo nessa área e toda ajuda é bem vinda. :)  



zero log
https://pkg.go.dev/github.com/rs/zerolog

retirar o ponteiro de autores
executar update sem alterar o array
proteger as operações do slice com mutex - 

pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

var jsonData = pm.response.json();
pm.test("Ensure Result", function () {
    pm.expect(jsonData.Total).to.eql(1);
    pm.expect(jsonData.Nfes[0].AccessKey).to.eql("35200472381189001001550010018766061478420160");
});