# MarcelinhoCarioca
Marcelinho Carioca é um bot para Discord feito em Golang. (para fins didáticos)

### Comandos:

- help (Lista os comandos)
- ping (Mostra o ping do websockey)
- avatar (Mostra o avatar do usuário)
- ban (Bane o usuário)
- unban (Desbane o usuário)
- mute (Muta [timeout] o usuário)
- unmute (Desmuta [timeout] o usuário)

### Para usar:
  Para clonar o repositório, você precisa do [Git](https://git-scm.com) e [Golang](https://golang.org).

 > No seu Terminal:
```
# Clonando o repositório
$ git clone https://github.com/Leoff00/go-diego-bot.git

# Entrando no repositório
$ cd MarcelinhoCarioca

# Rodar o projeto localmente
$ go mod download && go mod tidy && go run main.go
```

  Bibliotecas utilizadas:

  - [discordgo](https://pkg.go.dev/github.com/bwmarrin/discordgo)
  - godotenv (development)

## License

MIT