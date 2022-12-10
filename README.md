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
    go run clientMain/client_main.go
```

### Server 

```Bash 
    go run serverMain/server_main.go 
```
### Aclaraciones del TP:

No esta implementada la parda, el que comienza es el primero que creo la partida.

Si hay problemas para leer el cvs de las cartas, poner el path correspondiente en server/cardDealed.go linea 23
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

* Channels 

Permiten la comunicaion entre dos go rutings, es decir si yo quiero que una go ruting utilize los valores de otra se los puedo enviar por el channel establecido. El cannal debe definir un tipo. 
Son bloqueantes, lo cual permite que exista una sincronizacion. 
Se pueden definir tamaños:
* si es fijo se llama buffer y se bloquea si esta lleno o si esta vacio 

No hace falta cerrarlos. Puede hacerlo para indicar al receptor que ya no se envia nada. 