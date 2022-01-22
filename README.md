# Goland REST API com Mux
Olá está é a minha primeira aplicação em _Golang_ e calhou de ser uma aplicação REST.

Olá... talvez esse readme não estaja cumprindo muito bem o seu papel. mas o fato é que eu reescrevi esse código algumas dezenas de vezes e infelizmente, por burrice minha, não subi todas as alterações no github.

Pois bem, para resolver parte disso estou subindo essa branch com algumas modificações que eu estou achando bem pertinentes, mas infelizemente ainda está longe de eu chegar no fim com elas. Uma parte do código está bagunçada e a API não está mais funcionando porque ainda não fiz uma parte que é bem importante para ela funcionar, como não quero perder o trabalho feito até aqui e caso eu mude de ideia eu volto para o inicio e continuo de lá. *rs*

Blz. esse é o readme diário =)

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
