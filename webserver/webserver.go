package webserver

import "net/http" //paquete importado para trabajar con servidores web

func MiWebServer() {
	http.HandleFunc("/", home) //cuando el usuario viene a / ejecuta la funcion home, cuando viene a /login ejecuta login y asi sucesivamente la logica, direcciona la página a home (localhost/ muestra todo lo de home)
	//HandleFunc es un manejador que recibe la peticion y le decimos que lo que llegue a localhost/ ejecute home

	//una vez definida la ruta le decimos que se quede escuchando el puerto 3000
	http.ListenAndServe(":3000", nil) //la función "ListenandServe" escucha y sirve los datos en el puerto indicado, en este caso 3000, la sintaxis es ":puerto", nil (porque el segundo parametro no importa)
}

//ésta es la sintaxis de webserver cuando una funcion tiene que mostrar/procesar peticiones de http
func home(w http.ResponseWriter, r *http.Request) { //el segundo es con puntero *
	//la funcíon recibe 2 parámetros que son 2 interfaces estándar de net/http (w de writer y r de request)

	http.ServeFile(w, r, "./webserver/index.html")
	//la funcion "ServeFile" le sirve a nuestra web un archivo (en este caso "index.html") es ./ porque está en el mismo package, se pone (w,r, "direcion del archivo")

}
