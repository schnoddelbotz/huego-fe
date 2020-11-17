package web

import (
	"bytes"
	"github.com/amimof/huego"
	"html/template"
	"log"
	"net/http"
)
import "github.com/schnoddelbotz/huego-fe/hueController"

type server struct {
	hc *hueController.Controller
}

type indexTemplateData struct {
	HueIP  string
	Lights []huego.Light
}

func Serve(port string, hc *hueController.Controller) error {
	srv := &server{hc: hc}
	http.Handle("/assets/", http.FileServer(_escFS(false)))
	http.HandleFunc("/control", srv.controlHandler)
	http.HandleFunc("/", srv.indexHandler)
	log.Printf("Controlling Hue: %s", hc.IP())
	log.Printf("Starting huego-fe webserver on port %s", port)
	return http.ListenAndServe(port, nil)
}

func (s *server) indexHandler(w http.ResponseWriter, r *http.Request) {
	lights, err := s.hc.Lights()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	templateData := indexTemplateData{
		HueIP:  s.hc.IP(),
		Lights: lights,
	}
	w.Write(renderIndexTemplate(templateData))
}

func (s *server) controlHandler(w http.ResponseWriter, r *http.Request) {
	print("CTRL")
}

func renderIndexTemplate(data indexTemplateData) []byte {
	buf := &bytes.Buffer{}
	templateBinary := _escFSMustByte(false, "/assets/index.tpl.html")
	tpl, err := template.New("index").Parse(string(templateBinary))
	if err != nil {
		log.Fatalf("Template parsing error: %v\n", err)
	}
	err = tpl.Execute(buf, data)
	if err != nil {
		log.Printf("Template execution error: %v\n", err)
	}
	return buf.Bytes()
}
