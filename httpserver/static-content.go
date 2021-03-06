package httpserver

//go:generate go run static-content/make_static.go httpserver stylesDotCSS              text/css                      s static-content/styles.css                                         styles_dot_css.go

//go:generate go run static-content/make_static.go httpserver jsontreeDotJS             application/javascript        s static-content/jsontree.js                                        jsontree_dot_js.go

//go:generate go run static-content/make_static.go httpserver bootstrapDotCSS           text/css                      s static-content/bootstrap.min.css                                  bootstrap_dot_min_dot_css.go

//go:generate go run static-content/make_static.go httpserver bootstrapDotJS            application/javascript        s static-content/bootstrap.min.js                                   bootstrap_dot_min_dot_js.go
//go:generate go run static-content/make_static.go httpserver jqueryDotJS               application/javascript        s static-content/jquery-3.2.1.min.js                                jquery_underline_3_dot_2_dot_1_dot_min_dot_js.go
//go:generate go run static-content/make_static.go httpserver popperDotJS               application/javascript        b static-content/popper.min.js                                      popper_dot_min_dot_js.go

//go:generate go run static-content/make_static.go httpserver openIconicBootstrapDotCSS text/css                      s static-content/open-iconic/font/css/open-iconic-bootstrap.min.css open_iconic_bootstrap_dot_css.go

//go:generate go run static-content/make_static.go httpserver openIconicDotEOT          application/vnd.ms-fontobject b static-content/open-iconic/font/fonts/open-iconic.eot             open_iconic_dot_eot.go
//go:generate go run static-content/make_static.go httpserver openIconicDotOTF          application/font-sfnt         b static-content/open-iconic/font/fonts/open-iconic.otf             open_iconic_dot_otf.go
//go:generate go run static-content/make_static.go httpserver openIconicDotSVG          image/svg+xml                 s static-content/open-iconic/font/fonts/open-iconic.svg             open_iconic_dot_svg.go
//go:generate go run static-content/make_static.go httpserver openIconicDotTTF          application/font-sfnt         b static-content/open-iconic/font/fonts/open-iconic.ttf             open_iconic_dot_ttf.go
//go:generate go run static-content/make_static.go httpserver openIconicDotWOFF         application/font-woff         b static-content/open-iconic/font/fonts/open-iconic.woff            open_iconic_dot_woff.go
