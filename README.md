# TP-Teoria-del-Lenguaje

| Alumno  | Padron | Mail | 
| -------| --------|-------|
| Agustina Segura  | 104222  | asegura@fi.uba.ar |
| Stephanie Izquierdo Osorio | 104196  | sizquierdo@fi.uba.ar |
| Maria Sol Fontenla | 103870 | msfontenla@fi.uba.ar |



# Buildear 

```Bash
    go build 
```

### Client 

```Bash 
    go run clientMain/client.go
```

### Server 

```Bash 
    go run serverMain/server_main.go 
```

### ejemplos para el video 

* Como podemos aplicar composicion:
    si sacamos de nuestra estructura el nombre del atributo y dejamos directamente el tipo de dato 
    le estamos dando la funcionadidad de ese tipo a la estructura. 
    ej: 
    tenemos struct 
    ```Go 
    type Acceptor struct {
	socketListerner net.Listener
	players []Player
    }
    ```
    Vemos que posee el atributo socketListerner que es del tipo listener, si solo dejamos el tipo 
    ahora el acceptor posee la funcionalidad del socket y directamente invocamos a aceptador para realizar los metodos del socket 
     ```Go 
    type Acceptor struct {
	net.Listener
	players []Player
    }

    aceptador.Accept()
    acceptador.Close()
    ```