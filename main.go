package main

import (
	"flag"
	"log"
	tgClient "my-golang-project/clients/telegram"
	"my-golang-project/events/telegram"
	"my-golang-project/storage/files"
    event_consumer "my-golang-project/consumer/event-consumer"

)

const (
    tgBotHost = "api.telegram.org"
    storagePath = "storage"
    batchSize = 100 
)

func main() {

    eventProcessor := telegram.New(
        tgClient.New(tgBotHost, mustToken()), 
        files.New(storagePath),
    )
    
    log.Print("service started")
    
    consumer := event_consumer.New(eventProcessor, eventProcessor, batchSize)

    if err := consumer.Start(); err != nil {
        log.Fatal("service is stopped", err)
    }


}

func mustToken() string {
    token := flag.String(
        "tg-bot-token", 
        "", 
        "token for access to telegram bot")
        
        flag.Parse()

        if *token == ""{
            log.Fatal("token is not specified")
        }
    return *token
}