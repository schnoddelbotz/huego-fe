package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/amimof/huego"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

type server struct {
	hc      *hueController.Controller
	Version string
}

type response struct {
	Message string
	Error   string
}

type indexTemplateData struct {
	HueIP   string
	Lights  []huego.Light
	Version string
}

const (
	PowerOn  = "PowerOn"
	PowerOff = "PowerOff"
)

func Serve(port string, hc *hueController.Controller, huegofeVersion string) error {
	srv := &server{hc: hc, Version: huegofeVersion}
	http.Handle("/assets/", http.FileServer(_escFS(false)))
	http.HandleFunc("/control/", srv.controlHandler)
	http.HandleFunc("/", srv.indexHandler)
	log.Printf("Controlling Hue: %s", hc.IP())
	log.Printf("Starting huego-fe webserver %s on port %s", huegofeVersion, port)
	return http.ListenAndServe(port, nil)
}

func (s *server) indexHandler(w http.ResponseWriter, r *http.Request) {
	lights, err := s.hc.Lights()
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
	templateData := indexTemplateData{
		HueIP:   s.hc.IP(),
		Lights:  lights,
		Version: s.Version,
	}
	_, _ = w.Write(renderIndexTemplate(templateData))
}

func (s *server) controlHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 || pathParts[1] != "control" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message, err := s.executeControlCommand(pathParts[2:])
	data := response{
		Message: message,
	}
	statusCode := 200
	if err != nil {
		data.Error = err.Error()
		statusCode = 500
	}

	response, _ := json.Marshal(data) // panic on marshal error...?
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}

func (s *server) executeControlCommand(args []string) (string, error) {
	command := args[0]
	light, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}
	log.Printf("Executing command: %s on light %d", command, light)
	switch command {
	case PowerOn:
		return command, s.hc.PowerOn(light)
	case PowerOff:
		return command, s.hc.PowerOff(light)
	// TBD MORE
	default:
		return "?", errors.New("unknown command")
	}
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
