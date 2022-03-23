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

	name := "Open Systems"

	if val, ok := r.Header["X-Jwt-Payload"]; ok {
		nameB64,_ := base64.RawURLEncoding.DecodeString(val[0])

		var jwtPayload map[string]interface{}
		json.Unmarshal([]byte(nameB64), &jwtPayload)
		name = fmt.Sprintf("%v", jwtPayload["name"])

	}

	verb := os.Getenv("VERB") // glad then delighted

	fmt.Fprintf(w, `
	<!DOCTYPE html>
<html>

<head>
    <link href="https://sso.osdp.open.ch/static/favicon.svg" rel="icon" type="image/svg+xml">
    <style>
        body {
            color: #5EE0C8;
            background-color: black;
            font-family: 'Courier New';
            padding: 15%%;
            min-height: 100vh;
            min-width: 100vh;
        }

        a {
            color: #5EE0C8;
        }

        #collapsible {
            cursor: pointer;
            padding: 18px;
            border: none;
            text-align: left;
            display: none;
            outline: none;
        }

        #content {
            padding: 0 18px;
            display: none;
            overflow: hidden;
        }

        #logo {
            width: 10%%;
            position: absolute;
            bottom: 10px;
            right: 10px;
        }

        .looper {
            animation: perspect 4s infinite linear
        }

        @keyframes perspect {
            0%% {
                transform: perspective(5000px) rotateY(0deg)
            }

            25%% {
                transform: perspective(5000px) rotateY(90deg)
            }

            50%% {
                transform: perspective(5000px) rotateY(180deg)
            }

            75%% {
                transform: perspective(5000px) rotateY(270deg)
            }

            100%% {
                transform: perspective(5000px) rotateY(360deg)
            }
        }
    </style>
</head>

<body>
    <p id="d"></p>
    <p id="collapsible">OSDP <b>enabling</b> namespace-as-a-service features on display here:</p>
    <div id="content">
        <ul>
            <li>Incredible speed of setting up a new service / <b>product</b> / namespace</li>
            <li>Powerfull integrated <b>DNS</b> and <b>SSL</b> setup (https://)</li>
            <li>Amazing Microsoft <b>Azure Active Directory</b> Integration for Authentication and Authorization (you are %q,
                right?)</li>
            <li>Easy to use External <b>Secrets</b> Integration with hashicorp vault (%q): <a
                    href="https://hcv.dev.open.ch/ui/vault/secrets/central-dev/show/_shared/osdp-generic-ns/_public/townhall-verb">hcv</a>
            </li>
			<li>Automatic <b>uninterrupted</b> restart on secret change</li>
            <li>And much more out of the box as part of OSDP enabling <b>namespace-as-a-service</b>: Integrated monitoring and alerting, dashboards, various
                backends, etc </li>
			<li>This service - now even faster - can and is already used in <b>production</b> by teams for front and backend services, api's, pipelines, etc</li>
        </ul>
        <br />
        <p>
            Want to start your own project?<br />
            Want to know more about the features shown here?<br />
			<br />
            --> <a
                href="https://docs.open.ch/docs/display/DEV/OSDP+Townhall+Teardown">https://docs.open.ch/docs/display/DEV/OSDP+Townhall+Teardown</a>
        </p>
    </div>
    <div class="looper" id="logo">
        <svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 90 98'>
            <style>
                path {
                    fill: #5EE0C8;
                }
            </style>
            <path class='icon'
                d='M70.171 25.325a1.063 1.063 0 00-.088-.21v-.042s-.061-.097-.095-.13l-.088-.062a.418.418 0 00-.095-.079L26.707.117a.87.87 0 00-1.308.757v28.145L.446 42.883A.87.87 0 000 43.65v28.676c.005.307.17.588.435.74l43.037 24.652a.87.87 0 001.308-.759V68.84l24.973-13.942a.87.87 0 00.444-.757V25.525a.788.788 0 00-.026-.2zM27.142 2.382l40.415 23.134-23.624 13.071-16.79-9.585zm-1.741 28.633V57.66L1.743 70.81V44.163zm17.652 64.41L2.657 72.328l23.616-13.124 16.78 9.585v26.638zm25.401-41.764l-23.657 13.15V40.155l23.657-13.139v26.647z' />
        </svg>
    </div>
    <script>
        var i = 0;
        var t = 'Hi <b>%s</b>, delighted to see you. OSDP whishes you <b>%s</b> day 💥💥💥💥💥💥💥💥!';
        var s = 60;
        function tW() {
            if (i < t.length) {
                document.getElementById('d').innerHTML += t.charAt(i);
                i++;
                setTimeout(tW, s);
            } else {
                document.getElementById('d').innerHTML = t;
                document.getElementById('collapsible').style.display = "block";
            }

        }

        document.getElementById("collapsible").addEventListener("click", function () {
            var content = document.getElementById("content");
            if (content.style.display === "block") {
                content.style.display = "none";
            } else {
                content.style.display = "block";
            }
        });

        document.getElementById("logo").addEventListener("click", function () {
            if (document.body.style.color === "black") {
                document.body.style.color = "#5EE0C8";
                document.body.style.backgroundColor = "black";
            } else {
                document.body.style.color = "black";
                document.body.style.backgroundColor = "white";
            }
        });

        window.onload = tW;
    </script>
</body>

</html>
	
	`, name, verb, name, verb)
	// var keys strings.Builder
    // for k := range jwtPayload {
    //     keys.WriteString(k)
    // }
	// fmt.Printf("Serving: %s %v \n", nameB64, keys.String())
	fmt.Printf("Serving: %s \n", name)
}

func main() {
	http.HandleFunc("/", helloHandler) // Update this line of code

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
