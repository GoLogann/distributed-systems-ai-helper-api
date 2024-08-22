package gpt

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func GetResponse(prompt string) string {
	apiUrl := "https://api.openai.com/v1/chat/completions"

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": prompt},
		},
	})
	if err != nil {
		log.Fatalf("Error marshaling request body: %v", err)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer ")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %v", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}

	choices := result["choices"].([]interface{})
	firstChoice := choices[0].(map[string]interface{})
	message := firstChoice["message"].(map[string]interface{})
	content := message["content"].(string)

	return content
}
