package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// name := r.Header["User-Agent"]
	name := r.Header
	verb := os.Getenv("VERB") // glad then delighted
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	  <head>
		<style>
		  body {
			font-family: 'Courier New';
			background-image: url('https://www.open-systems.com/wp-content/uploads/2021/05/Home_Sphere_Tripplets_Green_R1-1.png');
			background-repeat: no-repeat;
			background-attachment: fixed;
			background-size: 100%% 100%%;
		  }
		</style>
	  </head>
	  <body>
		<p id="d"></p>
		<p>OSDP features on display here:</p>
		<ul>
			<li>Speed of setting up new services / products / namespaces</li>
			<li>Fully integrated DNS and SSL setup (https://)</li>
			<li>Full Microsoft Azure Active Directory Integration for Authentication and Authorization (you are %s, right?)</li>
			<li>External Secrets Integration (hcv)</li>
			<li>And much more (Integrated monitoring and alerting, dashboards, auto reload on config change, etc) </li>
		</ul>
		<p>
		  Want to start your own project?<br />
		  --> <a href="https://go.co">https://g.co</a>
		</p>
		<script>
		  var i=0; 
		  var t='Hi %s, %s to see you. OSDP whishes you an awesome day!';
		   var s=60; 
		   function tW(){if (i < t.length){document.getElementById('d').innerHTML +=t.charAt(i); i++; setTimeout(tW, s);}}
		   window.onload=tW;
		</script>
	  </body>
	</html>
	
	`, name, name, verb)
	fmt.Printf("Serving: %s \n", name)
}

func main() {
	http.HandleFunc("/", helloHandler) // Update this line of code

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
