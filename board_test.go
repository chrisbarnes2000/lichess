package lichess

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/lukemilby/lichess/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetBoardGameState(t *testing.T) {
	setUp()

	boardGameStateResponse := `
		{
		  "type": "gameFull",
		  "id": "5IrD6Gzz",
		  "rated": true,
		  "variant": {
			"key": "standard",
			"name": "Standard",
			"short": "Std"
		  },
		  "clock": {
			"initial": 1200000,
			"increment": 10000
		  },
		  "speed": "classical",
		  "perf": {
			"name": "Classical"
		  },
		  "createdAt": 1523825103562,
		  "white": {
			"id": "lovlas",
			"name": "lovlas",
			"provisional": false,
			"rating": 2500,
			"title": "IM"
		  },
		  "black": {
			"id": "leela",
			"name": "leela",
			"rating": 2390,
			"title": null
		  },
		  "initialFen": "startpos",
		  "state": {
			"type": "gameState",
			"moves": "e2e4 c7c5 f2f4 d7d6 g1f3 b8c6 f1c4 g8f6 d2d3 g7g6 e1g1 f8g7",
			"wtime": 7598040,
			"btime": 8395220,
			"winc": 10000,
			"binc": 10000,
			"status": "started"
		  },
		  "winner": "black"
		}`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(boardGameStateResponse)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	boardState, err := client.GetBoardGameState("game_id")
	assert.NotNil(t, boardState)
	assert.Nil(t, err)
	assert.EqualValues(t, "5IrD6Gzz", boardState.Id)
	assert.EqualValues(t, "black", boardState.Winner)
}

func TestMakeBoardMove(t *testing.T) {
	setUp()

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(`{"ok": true}`)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	ok, err := client.MakeBoardMove("board_id", "c2c4")
	assert.Nil(t, err)
	assert.EqualValues(t, true, ok.Ok)
}
