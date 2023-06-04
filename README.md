# Aula

Esse repositório é uma demonstração uma aplicação API Rest em Go, o desenvolvimento foi gravado e diponibilizado no YouTube, no Canal [AP Tech](https://www.youtube.com/@ap_tech0). A [Playlist](https://www.youtube.com/playlist?list=PL7541a7wciveRSb3kBNPvGiom8tL_ebEO) se contra nesse link, o ideal é assistir na ordem dés do começo.

## Requisitos
- Docker 23.0.5 ou superior, versões anteriores podem funcionar também.

## Commandos 
### setup
- irá copiar o arquivo `.env.example` e renomeando para `.env`;
- irá subir os containers de Go e MongoBD;
- irá iniciar a API Rest na porta 8080.

 ### cover
- Irá realizar os testes na aplicaça;
- Após os teste irá pritnar no terminar um relatorio dos testes e suas coberturas no código.


## Testar aplicação
- importe o arquivo [Aula.postman_collection.json](doc/Aula.postman_collection.json) em seu [Postman](https://www.postman.com/downloads/) ou em seu [Insomnia]https://insomnia.rest/download). A collection importada terá todas as rotas mapenadas;

- Execute o mando cover na sua maquina, lembrando que esse precisa ter o Golang 1.18 ou superior na sua máquina instalado.

## Dúvidas, Problemas ou Sugestões
- <b>Dúvidas</b>: pode-se usar a área de comentários nos vídeos do YouTube, mesmo que não tenha haver com o vídeo em questão;
- <b>Problemas</b>: pode-se abrir um Issue nesse mesmo repositório, ficarei feliz em melhorar o projeto;
- <b>Sugestões</b>: também podem ser postadas na área de comentários nos vídeos do YouTube, ou chamar em alguma das redes sóciais, os perfeis podem ser encontradas na área de Sobre em cada plataforma.