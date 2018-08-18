package views

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets80c0d7c4b5f47513d79ddc23ee8b0f4488cb1b9b = "<!DOCTYPE html>\r\n<html>\r\n  <head>\r\n    <meta charset=\"utf-8\">\r\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\r\n    <title>{{ .title }}</title>\r\n\r\n    <!-- CSS -->\r\n    <!--<link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.0/css/bulma.min.css\">   Bulma cdn -->\r\n    <link rel=\"stylesheet\" href=\"/public/css/bulma.css\">  <!-- Bulma cdn -->\r\n    <link rel=\"stylesheet\" href=\"http://code.ionicframework.com/ionicons/2.0.1/css/ionicons.min.css\">    <!-- Ionicons -->\r\n    <script src=\"https://unpkg.com/vue@2.5.2/dist/vue.js\" charset=\"utf-8\"></script> <!-- VUE --> \r\n\r\n     <style media=\"screen\">\r\n          section.hero{\r\n            background-image: -webkit-linear-gradient(45deg, #45145a, #ff5300);\r\n            background-image: linear-gradient(45deg, #45145a, #ff5300);\r\n          }\r\n         \r\n          section.hero h2 {\r\n            color:#fff7e0;\r\n            margin-top:10px;\r\n          }\r\n\r\n          section.hero h2 img {\r\n            border-radius: 95%;\r\n            box-shadow: 0 10px 16px 0 rgba(0,0,0,0.2),0 6px 20px 0 rgba(0,0,0,0.19) !important;\r\n          }\r\n\r\n          section.hero h3 {\r\n            color:#fff7e0;\r\n          }\r\n\r\n          blockquote {\r\n            background-color: #f5f5f5;\r\n            border-left: 5px solid #dbdbdb;\r\n            padding: .25em 1.5em;\r\n            margin-bottom:5px;\r\n          }  \r\n     </style>\r\n\r\n  </head>\r\n  <body>\r\n\r\n\r\n    <section class=\"hero\" >\r\n      <div class=\"hero-body\"  style=\"padding:0px;\">\r\n        <div class=\"container\">\r\n          <h2 class=\"title is-2 has-text-centered\">\r\n             <img width=\"200\" src=\"public/assets/storm.png\" />{{ .greeting }}\r\n          </h2>\r\n          <h3 class=\"title is-5 has-text-centered\">\r\n            {{ .words }}\r\n          </h3><br />\r\n\r\n        </div>\r\n      </div>\r\n    </section>\r\n\r\n    <div class=\"container\"><br />\r\n      <form method=\"POST\" action=\"/check\">\r\n            <div class=\"field\">\r\n              <label class=\"label\">Example POST form </label>\r\n              <div class=\"control\">\r\n                <input class=\"input\" name=\"test_field\" type=\"text\" placeholder=\"Enter text\">\r\n              </div>\r\n              <p class=\"help\">Enter some words and see in console data from this form</p>\r\n            </div>\r\n         <input class=\"button\" type=\"submit\" value=\"Submit form\">\r\n      </form>\r\n\r\n      <br /><br />\r\n\r\n      <div>\r\n\r\n        <label class=\"label\">Example of loop data</label>\r\n        <ul>\r\n          {{ range .items }}\r\n            <li>{{ . }}</li>\r\n          {{ end }}  \r\n        </ul>\r\n\r\n        <br />\r\n\r\n        <label class=\"label\">Example of loop data (Nested)</label>\r\n        <ul>\r\n          {{ range .items2 }}\r\n            <li>{{ .id }}) <b>{{ .name }}</b></li>\r\n          {{ end }}  \r\n        </ul>\r\n      </div>\r\n\r\n      <!-- <div class=\"field\">\r\n        <label class=\"label\">Test Sockets</label>\r\n        <blockquote>Be sure that you enable in <b>config.json</b> - <b>sockets.enabled: true</b> & <b>npm i socket.io -S</b></blockquote>\r\n        <div class=\"control\">\r\n          <input id=\"socket_field\" class=\"input\" type=\"text\" placeholder=\"Type 'hello'\">\r\n        </div>\r\n        <p class=\"help\">Type \"hello\" and receive answer from sockets</p>\r\n      </div>\r\n\r\n       <button onclick=\"SendToSockets();\" class=\"button is-light\">Send &nbsp;<i class=\"ion-paper-airplane\"></i></button> -->\r\n    \r\n      </div>\r\n\r\n  </body>\r\n</html>\r\n"
var _Assetsd275912fb2b3edf5b5b1728ba5f8470841591bce = "<html>\r\n<head>\r\n  <title>test page</title>\r\n</head> \r\n\r\n<body>\r\n   <h2>Hello there</h2>\r\n</body>\r\n\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/views": []string{"home.html", "test.html"}, "/": []string{"views"}}, map[string]*assets.File{
	"/views": &assets.File{
		Path:     "/views",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1534601187, 1534601187438035900),
		Data:     nil,
	}, "/views/home.html": &assets.File{
		Path:     "/views/home.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1533843856, 1533843856495835700),
		Data:     []byte(_Assets80c0d7c4b5f47513d79ddc23ee8b0f4488cb1b9b),
	}, "/views/test.html": &assets.File{
		Path:     "/views/test.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1534601212, 1534601212975496600),
		Data:     []byte(_Assetsd275912fb2b3edf5b5b1728ba5f8470841591bce),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1534602674, 1534602674265077500),
		Data:     nil,
	}}, "")
