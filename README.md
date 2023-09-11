# Servindo arquivos estáticos em go

Esse é um exemplo de um servidor que serve arquivos html. Os arquivos HTML são
gerados por meio de templates e possuem estilos servidos estaticamente (pasta static).

Para rodar o projeto, tenha o go instalado, execute o comando:

```sh
go run main.go
```

Após isso, vá para a url `localhost:3030` no navegador, lá estará a página index
e as rotas para as outras páginas.
