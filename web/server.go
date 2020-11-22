package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/amimof/huego"

	"github.com/schnoddelbotz/huego-fe/huecontroller"
)

type server struct {
	hc      *huecontroller.Controller
	Port    string
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
	Attempt int
	Port    string
}

const (
	// PowerOn string as used in URLs
	PowerOn = "PowerOn"
	// PowerOff string as used in URLs
	PowerOff = "PowerOff"
	// Brightness string as used in URLs
	Brightness = "Brightness"
)

// Serve starts the huego-fe web server on the given port, to enable light control via browser
func Serve(port string, hc *huecontroller.Controller, huegofeVersion string) error {
	srv := &server{hc: hc, Port: port, Version: huegofeVersion}
	log.Printf("Starting huego-fe %s webserver for Controlling Hue: %s", huegofeVersion, hc.IP())
	log.Printf("Listening on %s ( visit http://localhost:%s/ )", huegofeVersion, port)
	return http.ListenAndServe(port, accessLogHandler(serveMux(srv)))
}

func serveMux(srv *server) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.FileServer(_escFS(false)))
	mux.HandleFunc("/control/", srv.controlHandler)
	mux.HandleFunc("/login", srv.loginHandler)
	mux.HandleFunc("/", srv.indexHandler)
	return mux
}

func (s *server) loginHandler(w http.ResponseWriter, r *http.Request) {
	_attempt := r.URL.Query().Get("attempt")
	attempt, err := strconv.Atoi(_attempt)
	if _attempt == "" || err != nil {
		http.Error(w, "No attempt?", http.StatusBadRequest)
		return
	}
	if !s.hc.IsLoggedIn() {
		if attempt > 29 {
			w.Header().Set("Content-Type", "text/html")
			_, _ = w.Write([]byte(`Unable to discover/login to Hue. <a href="/">Click here to retry</a>.`))
			return
		}
		err := s.hc.Login()
		if err == nil {
			err = s.hc.SavePrefs()
			if err != nil {
				log.Fatalf("could not write Hue settings file: %s", err)
			}
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		// maybe put param/counter into template ... and give up after n retries?
		loginTemplate := _escFSMustByte(false, "/assets/login.tpl.html")
		loginPage := &bytes.Buffer{}
		tpl, _ := template.New("index").Parse(string(loginTemplate))
		data := &indexTemplateData{
			Attempt: attempt + 1,
			Port:    s.Port,
		}
		_ = tpl.Execute(loginPage, data)
		_, _ = w.Write(loginPage.Bytes())
		return
	}
	http.Redirect(w, r, "/", 302)
}

func (s *server) indexHandler(w http.ResponseWriter, r *http.Request) {
	if !s.hc.IsLoggedIn() {
		http.Redirect(w, r, "/login?attempt=0", http.StatusFound)
		return
	}
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

	message, err := s.executeControlCommand(pathParts[2:], r.URL.Query())
	data := response{
		Message: message,
	}
	statusCode := 200
	if err != nil {
		data.Error = err.Error()
		statusCode = 500
		log.Printf("Control command failed: %s", err)
	}

	response, _ := json.Marshal(data) // panic on marshal error...?
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}

func (s *server) executeControlCommand(args []string, query url.Values) (string, error) {
	command := args[0]
	light, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}
	switch command {
	case PowerOn:
		return command, s.hc.PowerOn(light)
	case PowerOff:
		return command, s.hc.PowerOff(light)
	case Brightness:
		toValue, err := strconv.Atoi(query.Get("to"))
		if err != nil {
			return "!", err
		}
		return command, s.hc.SetBrightness(light, uint8(toValue))
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
