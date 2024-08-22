package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	// Cria uma nova requisição HTTP
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Configura os headers da requisição
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer ")

	// Envia a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	defer resp.Body.Close()

	// Verifica o status da resposta
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %v", resp.Status)
	}

	// Decodifica a resposta JSON
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}

	// Extrai a resposta do modelo
	choices := result["choices"].([]interface{})
	firstChoice := choices[0].(map[string]interface{})
	message := firstChoice["message"].(map[string]interface{})
	content := message["content"].(string)

	return content
}

func main() {
	// Exemplo de uso da função
	response := GetResponse("Hello!")
	fmt.Println("Response from GPT:", response)
}
