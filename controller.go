package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var mutex = &sync.Mutex{}
var chatHistory = make(map[string][]ChatMessage)

func CallOpenAI(message string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": message,
			},
		},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		logger.Println("Error marshaling JSON:", err)
		return "", nil
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Println("Error API OpenAI,", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var openAIResp OpenAIResponse
	err = json.Unmarshal(body, &openAIResp)
	if err != nil {
		logger.Println("Error unmarshalling response:", err.Error())
		return "", err
	}

	if len(openAIResp.Choices) > 0 {

		logger.Println("OpenAI Response:", openAIResp.Choices[0].Message.Content)
		return openAIResp.Choices[0].Message.Content, nil
	} else {
		logger.Println("my response ai error", openAIResp)
	}

	return "", nil
}

func SendMessageHandler(c *gin.Context) {
	var message ChatMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		logger.Println("Detect for body:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message.Timestamp = time.Now()

	mutex.Lock()
	chatHistory[message.Receiver] = append(chatHistory[message.Receiver], message)
	mutex.Unlock()

	aiResponse, err := CallOpenAI(message.Message)
	if err != nil {
		logger.Println("Proses OpenAi Error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process message with OpenAI"})
		return
	} else {
		logger.Println("response AI now", aiResponse)
	}

	aiMessage := ChatMessage{
		Sender:    "User",
		Receiver:  "System",
		Message:   aiResponse,
		Timestamp: time.Now(),
		IsAI:      true,
	}

	mutex.Lock()
	chatHistory[aiMessage.Receiver] = append(chatHistory[aiMessage.Receiver], aiMessage)
	mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"status": "Message sent successfully"})
}

func GetMessagesHandler(c *gin.Context) {

	mutex.Lock()
	messages, ok := chatHistory["System"]
	mutex.Unlock()

	if !ok {
		logger.Println("Null Data")
		c.JSON(http.StatusOK, gin.H{"message": "No messages found"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	})
}
