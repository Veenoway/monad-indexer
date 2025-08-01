package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

type MonadTransaction struct {
    Hash             string `json:"hash"`
    From             string `json:"from"`
    To               string `json:"to"`
    Value            string `json:"value"`
    Gas              string `json:"gas"`
    GasPrice         string `json:"gasPrice"`
    BlockNumber      string `json:"blockNumber"`
    TransactionIndex string `json:"transactionIndex"`
}

type JSONRPCRequest struct {
    ID      int           `json:"id"`
    Method  string        `json:"method"`
    Params  []interface{} `json:"params"`
    JSONRPC string        `json:"jsonrpc"`
}

type JSONRPCResponse struct {
    ID      int             `json:"id"`
    Result  json.RawMessage `json:"result"`
    Error   *JSONRPCError   `json:"error"`
    JSONRPC string          `json:"jsonrpc"`
}

type JSONRPCError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}


func main() {

	err := godotenv.Load("../../.env")

	pwd, _ := os.Getwd()
    fmt.Println("🔍 Répertoire de travail actuel:", pwd)

	if err != nil {
		fmt.Println("⚠️ Fichier .env non trouvé, utilisation des valeurs par défaut")
		return
	}
	
	router := gin.Default()

	router.GET("/hello/:name", func(c * gin.Context) {
		name:= c.Param("name")

		c.JSON(200,gin.H{
			"message":"Salut " + name + " !",
		})
	})

	go listenToBlockEvents()

	fmt.Println("Server is running on port 8080")
	fmt.Println("Essaye d'aller sur:")
    fmt.Println("- http://localhost:8080/")
    fmt.Println("- http://localhost:8080/hello/tonnom")
	router.Run(":8080")	
	
	
}

func listenToBlockEvents() {

	url:= os.Getenv("WSS_ENDPOINT")
	if url == "" {
		log.Println(".env isn't findable")
	}

	fmt.Println("Connecting to Monad Infra...")

	ws, _, err := websocket.DefaultDialer.Dial(url,nil)
	if err != nil {
		log.Println("Error happen while connecting wss...")
		return
	}

	defer ws.Close()

	subscribeRequest:= JSONRPCRequest{
		ID: 1,
		Method: "eth_subscribe",
		Params: []interface{}{"newHeads"},
		JSONRPC: "2.0",
	}

	requestBytes, err := json.Marshal(subscribeRequest)
	if err != nil {
		log.Println("Impossible to parse into json",err)
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, requestBytes)
	if err != nil {
		log.Println("Error while connecting to connect wss")
		return
	}

	fmt.Println("Connected to Monad WebSocket")

	for {
		_, messageBytes,err := ws.ReadMessage()
		if err != nil {
			log.Println("Error while reading message")
			continue
		}

		var response map[string]interface{}

		err = json.Unmarshal(messageBytes, &response)
		if err != nil {
			log.Println("Error while unmarshaling message byte")
		}

		if id, exists := response["id"]; exists {
			if id == 999 {
				fmt.Println("We got a block here:\n\n %v\n\n%v", response, id)
			} 
		} else if result, ok := response["method"]; ok {
			fmt.Println("Result is a method",result)
		} else {
			fmt.Println("Result is a method",result)
		}
		fmt.Println("This is the response :\n%v",response)
	}
}