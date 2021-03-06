package lichess

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/lukemilby/lichess/utils/mocks"
	"github.com/stretchr/testify/assert"
)

const gameResponce = `{
	"id": "q7ZvsdUF",
	"createdAt": 1514505150384,
	"lastMoveAt": 1514505592843,
	"status": "draw",
	"players": {
	  "white": {
		"user": {
		  "name": "Lance5500",
		  "title": "LM",
		  "patron": true,
		  "id": "lance5500"
		},
		"rating": 2389,
		"ratingDiff": 4
	  },
	  "black": {
		"user": {
		  "name": "TryingHard87",
		  "id": "tryinghard87"
		},
		"rating": 2498,
		"ratingDiff": -4
	  }
	},
	"moves": "d4 d5 c4 c6 Nc3 e6 e4 Nd7 exd5 cxd5 cxd5 exd5 Nxd5 Nb6 Bb5+ Bd7 Qe2+ Ne7 Nxb6 Qxb6 Bxd7+ Kxd7 Nf3 Qa6 Ne5+ Ke8 Qf3 f6 Nd3 Qc6 Qe2 Kf7 O-O Kg8 Bd2 Re8 Rac1 Nf5 Be3 Qe6 Rfe1 g6 b3 Bd6 Qd2 Kf7 Bf4 Qd7 Bxd6 Nxd6 Nc5 Rxe1+ Rxe1 Qc6 f3 Re8 Rxe8 Nxe8 Kf2 Nc7 Qb4 b6 Qc4+ Nd5 Nd3 Qe6 Nb4 Ne7 Qxe6+ Kxe6 Ke3 Kd6 g3 h6 Kd3 h5 Nc2 Kd5 a3 Nc6 Ne3+ Kd6 h4 Nd8 g4 Ne6 Ke4 Ng7 Nc4+ Ke6 d5+ Kd7 a4 g5 gxh5 Nxh5 hxg5 fxg5 Kf5 Nf4 Ne3 Nh3 Kg4 Ng1 Nc4 Kc7 Nd2 Kd6 Kxg5 Kxd5 f4 Nh3+ Kg4 Nf2+ Kf3 Nd3 Ke3 Nc5 Kf3 Ke6 Ke3 Kf5 Kd4 Ne6+ Kc4",
	"clock": {
	  "initial": 300,
	  "increment": 3,
	  "totalTime": 420
	}
  }`

func TestExportGame(t *testing.T) {
	setUp()

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(gameResponce)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	game, err := client.ExportGame("qaddcxa2")
	assert.NotNil(t, game)
	assert.Nil(t, err)
	assert.EqualValues(t, "lance5500", game.Players.White.User.ID)
	assert.EqualValues(t, 420, game.Clock.TotalTime)
}

func TestOngoingGames(t *testing.T) {
	setUp()

	ongoingGamesResponse := `{
		"nowPlaying": [
		  {
			"fullId": "knbD9FPUqhra",
			"gameId": "knbD9FPU",
			"fen": "rnbqkbnr/pp3pp1/8/3p3p/1PpPp3/P1P1P3/5PPP/RNBQKBNR",
			"color": "white",
			"lastMove": "c5c4",
			"variant": {
			  "key": "standard",
			  "name": "Standard"
			},
			"speed": "correspondence",
			"perf": "correspondence",
			"rated": true,
			"opponent": {
			  "id": "thibot",
			  "username": "BOT thibot",
			  "rating": 1500
			},
			"isMyTurn": true
		  }
		]
	  }`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(ongoingGamesResponse)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	ongoingGames, err := client.OngoingGames()
	assert.NotNil(t, ongoingGames)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, len(*ongoingGames))
	assert.EqualValues(t, "thibot", (*ongoingGames)[0].Opponent.Id)
}
