package bff

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (gr *gameResource) SetHighScore(ctx *gin.Context) {
	highScoreStr := ctx.Param("hs")
	highScore, err := strconv.ParseFloat(highScoreStr, 64)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse high score")
	}
	gr.gameClient.SetHighScore(context.Background(), &pbHighscore.SetHighScoreRequest{
		HighScore: highScore,
	})
}

func (gr *gameResource) GetHighScore(ctx *gin.Context) {
	resp, err := gr.gameClient.GetHighScore(context.Background(), &pbHighscore.GetHighScoreRequest{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get high score")
		return
	}
	hs := strconv.FormatFloat(resp.GetHighScore(), 'e', -1, 64)
	ctx.JSONP(200, gin.H{
		"hs": hs,
	})
}

func (gr *gameResource) GetSize(ctx *gin.Context) {
	resp, err := gr.gameEngineClient.GetSize(context.Background(), &pbGameengine.GetSizeRequest{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get size")
		return
	}

	ctx.JSONP(200, gin.H{
		"size": resp.GetSize(),
	})
}

func (gr *gameResource) SetScore(ctx *gin.Context) {
	scoreStr := ctx.Param("score")
	score, _ := strconv.ParseFloat(scoreStr, 64)

	_, err := gr.gameEngineClient.SetScore(context.Background(), &pbGameengine.SetScoreRequest{
		Score: score,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to set score")
		return
	}
}
