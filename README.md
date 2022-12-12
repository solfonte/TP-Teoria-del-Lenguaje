# TP-Teoria-del-Lenguaje

| Alumno  | Padron | Mail | 
| -------| --------|-------|
| Agustina Segura  | 104222  | asegura@fi.uba.ar |
| Stephanie Izquierdo Osorio | 104196  | sizquierdo@fi.uba.ar |
| Maria Sol Fontenla | 103870 | msfontenla@fi.uba.ar |

# Link al video

[Go - Nosotras](https://drive.google.com/drive/folders/1wnyuQaBviVRMP4Hs4lXqXphPqfuFxGzC)

# Lenguaje Go 

# Buildear 

```Bash
    go build 
```

### Client 

```Bash 
    go run clientMain/client_main.go
```

### Server 

```Bash 
    go run serverMain/server_main.go 
```

* Acepta multiples jugadores 
* Se desconeta al servidor con el comando Q

### Aclaraciones del TP:

* No esta implementada la parda
* El que creo la partida es el que comienza primero en el juego
* Se implemento truco y retruco 
* Se implemeneto envido, envido - envido 
* Si ambos jugadores tiran una carta del mismo valor gana el que la tiro primero
* Cuando un jugador esta en espera de que el otro jugador espere este puede irse al   mazo o ver que cartas tiene. En caso que no responde por 6 segundos no podra solicitar mas ninguna opcion y debe esperar a que su oponenete juege.
* Si hay problemas para leer el cvs de las cartas, poner el path correspondiente en server/cardDealed.go linea 23

# Recursos 
La siguiente carpeta contiene documentacion realizada durante el desarrollo del trabajo practico  
[Recursos](https://drive.google.com/drive/folders/1uluZtbpqz5h7tG4S_XXy6tW1SR6HyGTK)