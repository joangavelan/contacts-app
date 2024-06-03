package toast

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	INFO    = "info"
	SUCCESS = "success"
	WARNING = "warning"
	ERROR   = "error"
)

type Toast struct {
	Variant string `json:"variant"`
	Message string `json:"message"`
}

type TriggerToastEvent struct {
	TriggerToast Toast `json:"triggerToast"`
}

// Creates a new Toast with the given variant and message.
func new(variant, message string) Toast {
	return Toast{variant, message}
}

func Info(message string) Toast {
	return new(INFO, message)
}

func Success(message string) Toast {
	return new(SUCCESS, message)
}

func Warning(message string) Toast {
	return new(WARNING, message)
}

func Error(message string) Toast {
	return new(ERROR, message)
}

// ToJsonEvent constructs the "triggerToast" event in JSON format.
func (t Toast) ToJsonEvent() ([]byte, error) {
	event := TriggerToastEvent{t}

	jsonData, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("error marshalling toast event: %w", err)
	}

	return jsonData, nil
}

// WriteToHeader writes the toast event to the HX-Trigger response header.
func (t Toast) WriteToHeader(w http.ResponseWriter) error {
	event, err := t.ToJsonEvent()
	if err != nil {
		return fmt.Errorf("error writing toast event: %w", err)
	}

	w.Header().Set("HX-Trigger", string(event))

	return nil
}
