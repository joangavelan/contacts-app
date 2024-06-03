package toast

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestToast(t *testing.T) {
	t.Run("New creates a new Toast with the given variant and message", func(t *testing.T) {
		toast := new(INFO, "Test message")
		if toast.Variant != INFO {
			t.Errorf("expected %s, got %s", INFO, toast.Variant)
		}
		if toast.Message != "Test message" {
			t.Errorf("expected %s, got %s", "Test message", toast.Message)
		}
	})

	t.Run("Variants create the correct Toast instances", func(t *testing.T) {
		tests := []struct {
			name    string
			method  func(string) Toast
			variant string
			message string
		}{
			{"Info", Info, INFO, "Info message"},
			{"Success", Success, SUCCESS, "Success message"},
			{"Warning", Warning, WARNING, "Warning message"},
			{"Error", Error, ERROR, "Error message"},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				toast := tc.method(tc.message)
				if toast.Variant != tc.variant {
					t.Errorf("expected %s, got %s", tc.variant, toast.Variant)
				}
				if toast.Message != tc.message {
					t.Errorf("expected %s, got %s", tc.message, toast.Message)
				}
			})
		}
	})

	t.Run("ToJsonEvent creates the correct JSON structure", func(t *testing.T) {
		toast := Info("Test message")
		jsonData, err := toast.ToJsonEvent()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		var event TriggerToastEvent
		err = json.Unmarshal(jsonData, &event)
		if err != nil {
			t.Fatalf("expected no error unmarshalling, got %v", err)
		}

		if event.TriggerToast.Variant != toast.Variant {
			t.Errorf("expected %s, got %s", toast.Variant, event.TriggerToast.Variant)
		}
		if event.TriggerToast.Message != toast.Message {
			t.Errorf("expected %s, got %s", toast.Message, event.TriggerToast.Message)
		}
	})

	t.Run("Write sets the correct HX-Trigger header", func(t *testing.T) {
		toast := Success("Test message")
		recorder := httptest.NewRecorder()
		err := toast.WriteToHeader(recorder)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		header := recorder.Header().Get("HX-Trigger")
		if header == "" {
			t.Fatalf("expected HX-Trigger header to be set")
		}

		var event TriggerToastEvent
		err = json.Unmarshal([]byte(header), &event)
		if err != nil {
			t.Fatalf("expected no error unmarshalling, got %v", err)
		}

		if event.TriggerToast.Variant != toast.Variant {
			t.Errorf("expected %s, got %s", toast.Variant, event.TriggerToast.Variant)
		}
		if event.TriggerToast.Message != toast.Message {
			t.Errorf("expected %s, got %s", toast.Message, event.TriggerToast.Message)
		}
	})
}
