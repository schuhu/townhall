package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// name := r.Header["User-Agent"]
	header := r.Header["X-Jwt-Payload"][0]

	nameB64,_ := base64.StdEncoding.DecodeString(header)
	var jwtPayload map[string]interface{}
	json.Unmarshal([]byte(nameB64), &jwtPayload)
	name := jwtPayload["name"]

	verb := os.Getenv("VERB") // glad then delighted
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	  <head>
		<style>
		  body {
			font-family: 'Courier New';
			min-height: 100vh; 
			min-width: 100vh;
			background-image: url('https://sso.osdp.open.ch/static/favicon.svg');
			background-repeat: no-repeat;
			background-size: 50%% 50%%;
			background-position: bottom 10px right 10px;
		  }
		</style>
	  </head>
	  <body>
		<p id="d"></p>
		<p>OSDP features on display here:</p>
		<ul>
			<li>Speed of setting up new services / products / namespaces</li>
			<li>Fully integrated DNS and SSL setup (https://)</li>
			<li>Full Microsoft Azure Active Directory Integration for Authentication and Authorization (you are %q, right?)</li>
			<li>External Secrets Integration (%q): <a href="https://hcv.dev.open.ch/ui/vault/secrets/central-dev/show/_shared/osdp-generic-ns/_public/townhall-verb">hcv</a></li>
			<li>And much more (Integrated monitoring and alerting, dashboards, auto reload on config change, etc) </li>
		</ul>
		<p>
		  Want to start your own project?<br />
		  --> <a href="https://docs.open.ch/docs/display/DEV/OSDP+Documentation">https://link to a tutorial</a>
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
	
	`, name, verb, name, verb)
	fmt.Printf("Serving: %s \n", name)
}

func main() {
	http.HandleFunc("/", helloHandler) // Update this line of code

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
