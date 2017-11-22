package main

import (
	"log"
	"os"
	"strconv"

	"kairu/src/sdk"
)

func main() {
	botName := "Kairu"
	connection := sdk.NewConnection(botName)

	logging := true
	if logging {
		directoryName := "log"
		fileName := directoryName + "/" + strconv.Itoa(connection.PlayerTag) + ".log"

		if _, err := os.Stat(directoryName); os.IsNotExist(err) {
			os.MkdirAll(directoryName, os.ModePerm)
		} else if err != nil {
			log.Fatal(err)
		}

		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	log.Println("--- NEW GAME ---")
	gameTurn := 1
	for true {
		log.Printf("Turn %v\n", gameTurn)

		gameMap := connection.UpdateMap()
		myPlayer := gameMap.Players[gameMap.MyID]
		myShips := myPlayer.Ships

		commands := []string{}
		for i := 0; i < len(myShips); i++ {
			ship := myShips[i]
			if ship.DockingStatus == sdk.UNDOCKED {
				commands = append(commands, sdk.StrategyBasicBot(ship, gameMap))
			}
		}
		connection.SubmitCommands(commands)

		gameTurn++
	}
}
