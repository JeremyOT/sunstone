package sunstone

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type joinRequest struct {
	Nodes []string `json:"nodes"`
}

type resetRequest struct {
}

func (m *Sunstone) joinHandler(body []byte) interface{} {
	var request joinRequest
	if err := json.Unmarshal(body, &request); err != nil {
		log.Println("[Error] Bad command:", err)
		return nil
	}
	log.Println("Received join command:", request.Nodes)
	if nodes, err := m.Join(request.Nodes); err != nil {
		log.Println("[Error] Failed to join:", err)
		return nil
	} else {
		log.Println("Joined", nodes, "nodes")
	}
	return request
}

func (m *Sunstone) resetHandler(body []byte) interface{} {
	var request resetRequest
	if err := json.Unmarshal(body, &request); err != nil {
		log.Println("[Error] Bad command:", err)
		return nil
	}
	log.Println("Received shutdown command")
	m.Shutdown()
	return request
}

func addHandler(serveMux *http.ServeMux, pattern string, handler func(body []byte) interface{}) {
	serveMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		response := handler(buf)
		if response != nil {
			responseBody, _ := json.Marshal(response)
			w.Header().Add("Content-Type", "application/json")
			w.Write(responseBody)
		}
	})
}

func (m *Sunstone) ListenForControl(address string) error {
	serveMux := http.NewServeMux()
	addHandler(serveMux, "/join", m.joinHandler)
	addHandler(serveMux, "/reset", m.resetHandler)
	return http.ListenAndServe(address, serveMux)
}

func CommandJoin(address string, nodes []string) error {
	body, _ := json.Marshal(joinRequest{Nodes: nodes})
	resp, err := http.Post("http://"+address+"/join", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func CommandReset(address string) error {
	body, _ := json.Marshal(resetRequest{})
	resp, err := http.Post("http://"+address+"/reset", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
