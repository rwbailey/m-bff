package bff

import (
	"github.com/rs/zerolog/log"
	pbGameengine "github.com/rwbailey/m-apis/game-engine/v1"
	pbHighscore "github.com/rwbailey/m-apis/highscore/v1"
	"google.golang.org/grpc"
)

type gameResource struct {
	gameClient       pbHighscore.GameClient
	gameEngineClient pbGameengine.GameEngineClient
}

func NewGameResource(gameClient pbHighscore.GameClient, gameEngineClient pbGameengine.GameEngineClient) *gameResource {
	return &gameResource{
		gameClient:       gameClient,
		gameEngineClient: gameEngineClient,
	}
}

func NewGrpcGameServiceClient(serverAddr string) (pbHighscore.GameClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msg("Failed to dial: " + err.Error())
		return nil, err
	}
	log.Info().Msg("Connected to: " + serverAddr)

	if conn == nil {
		log.Info().Msg("m-highscore conn is nil in m-bff")
		return nil, err
	}

	client := pbHighscore.NewGameClient(conn)

	return client, nil
}

func NewGrpcGameEngineServiceClient(serverAddr string) (pbGameengine.GameEngineClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msg("Failed to dial: " + err.Error())
		return nil, err
	}
	log.Info().Msg("Connected to: " + serverAddr)

	if conn == nil {
		log.Info().Msg("m-game-engine conn is nil in m-bff")
		return nil, err
	}

	client := pbGameengine.NewGameEngineClient(conn)

	return client, nil
}
