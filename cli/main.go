package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/rwbailey/m-bff/bff"
)

func main() {
	grpcAddrHighScore := flag.String("hsAddr", "localhost:50051", "address of the high score service")
	grpcAddrGameEngine := flag.String("geAddr", "localhost:60051", "address of the game engine service")
	serverAddr := flag.String("httpAddr", "localhost:8081", "address of the http server")
	flag.Parse()

	gameClient, err := bff.NewGrpcGameServiceClient(*grpcAddrHighScore)
	if err != nil {
		log.Error().Err(err).Msg("Error getting gameClient")
	}

	gameEngineClient, err := bff.NewGrpcGameEngineServiceClient(*grpcAddrGameEngine)
	if err != nil {
		log.Error().Err(err).Msg("Error getting gameEngineClient")
	}
	gr := bff.NewGameResource(gameClient, gameEngineClient)

	router := gin.Default()

	router.GET("/geths", gr.GetHighScore)
	router.GET("/seths/:hs", gr.SetHighScore)
	router.GET("/getsize", gr.GetSize)
	router.GET("/setscore/:score", gr.SetScore)

	err = router.Run(*serverAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start bff server")
	}
	log.Info().Msgf("Started http server at %v", *serverAddr)
}
