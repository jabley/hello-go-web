package main

import (
	"bytes"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	indexHTML = `<!DOCTYPE html>
<html>
	<head>
		<title>Welcome to my service</title>
		<style type="text/css">
			#footer {
				border-top: 10px solid #005ea5;
			    background-color: #dee0e2;
			}
			#footer ul {
				list-style: none;
			}
			#footer ul li {
    			display: inline-block;
    			margin: 0 15px 0 0;
			}
			#overview p {
				margin: 0 30px 0 30px;
			}
			.homepage-top {
				text-align: center;
    			background: #005ea5;
    			color: #fff;
			}
			.values-list li {
				display: list-item;
				list-style: none;
			}
			.values-list li h3 {
    			font-size: 19px;
    			line-height: 1.31579;
    		}
    		.values-list table {
				text-align: center;
				width: 80%;
    		}
			.visuallyhidden {
 			   position: absolute;
    			left: -9999em;
			}
		</style>
	</head>
	<body>
		<header class="homepage-top">
			<h1>Welcome!</h1>
			<p>A simple 12-factor app showing some environment configuration.</p>
		</header>
		<main>
			<section id="overview" aria-labelledby="overview-label">
				<h2 id="overview-label" class="visuallyhidden">Overview</h2>
				<p>Typically, this application is run in multiple hosting providers, with different values for <code>HELLO_VERSION</code> in each provider.</p>
				<p>This makes it possible to front the application with a CDN / load-balancer and see different values coming back depending on which origin served the request.</p>
				<p>It also allows you to demonstrate zero-downtime failover of the CDN/load-balancer, if suitably configured.</p>
			</section>
			<section id="environment-variables" aria-labelledby="environment-variables-label">
				<h2 id="environment-variables-label" class="visuallyhidden">Selected environment variables</h2>
				<div class="values-list">
					<table>
						<thead>
							<tr>
								<th width="50%">Key</th>
								<th width="50%">Value</th>
							</tr>
						</thead>
						<tbody>

					{{range .}}
						<tr>
							<td>{{.Key}}</td>
							<td>{{.Value}}</td>
						</tr>
					{{end}}
						</tbody>
					</table>
				</div>
			</section>
		</main>
		<footer id="footer">
			<div class="footer-meta">
				<h2 class="visuallyhidden">Support links</h2>
				<ul>
					<li><a href="https://github.com/jabley/hello-go-web">Source</a></li>
					<li>Built by <a href="https://twitter.com/jabley">James Abley</a></li>
				</ul>
			</div>
		</footer>
	</body>
</html>
`
)

// KeyValue makes the ENV vars into a first-class data structure
type KeyValue struct {
	Key   string
	Value string
}

// KeyValues is a shorter way of referencing an array
type KeyValues []*KeyValue

var (
	tmpl = template.Must(template.New("index.html").Parse(indexHTML))
	body []byte
)

func main() {
	flag.Parse()

	port := getDefaultConfig("PORT", "8080")

	values := getKeyValues()
	var b bytes.Buffer
	if err := tmpl.Execute(&b, values); err != nil {
		panic(err)
	}
	body = b.Bytes()

	http.HandleFunc("/", mainHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getDefaultConfig(name, fallback string) string {
	if val := os.Getenv(name); val != "" {
		return val
	}
	return fallback
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Set("Cache-Control", "private, no-cache, no-store, must-revalidate")
	w.Write(body)
}

func getKeyValues() KeyValues {
	result := make(KeyValues, 2)
	result[0] = &KeyValue{"PORT", os.Getenv("PORT")}
	result[1] = &KeyValue{"HELLO_VERSION", os.Getenv("HELLO_VERSION")}
	return result
}

func newKeyValue(kv string) *KeyValue {
	s := strings.Split(kv, "=")
	return &KeyValue{Key: s[0], Value: s[1]}
}
