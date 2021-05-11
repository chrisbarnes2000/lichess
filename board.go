package lichess

import (
	"fmt"
)

type Board struct {
	GameId string
}

type BoardGameState struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Rated bool   `json:"rated"`
	Clock struct {
		Initial   int `json:"initial"`
		Increment int `json:"Increment"`
	}
	Speed string `json:"speed"`
	Perf  struct {
		Name string `json:"name"`
	}
	CreatedAt int `json:"createdAt"`
	White     struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Provisional bool   `json:"provisional"`
		Rating      int    `json:"rating"`
		Title       string `json:"title"`
	}
	Black struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Rating int    `json:"rating"`
		Title  string `json:"title"`
	}
	InitialFen string `json:"initialFen"`
	State      struct {
		Type   string `json:"type"`
		WTime  int    `json:"wtime"`
		BTime  int    `json:"btime"`
		Status string `json:"status"`
		Winner string `json:"winner"`
	}
	Moves  string `json:"moves"`
	WTime  int    `json:"wtime"`
	BTime  int    `json:"btime"`
	Status string `json:"status"`
	Winner string `json:"winner"`
}

func (c *Client) GetBoardGameState(game_id string) (*BoardGameState, error) {
	boardGameState := new(BoardGameState)
	req, err := c.newRequest("GET", fmt.Sprintf("/api/board/game/stream/%s", game_id), nil)
	if err != nil {
		return nil, err
	}
	_, err = c.do(req, &boardGameState)
	if err != nil {
		return nil, err
	}

	return boardGameState, err
}

func (c *Client) MakeBoardMove(board_id string, move string) (*Ok, error) {
	ok := new(Ok)
	req, err := c.newRequest("POST", fmt.Sprintf("/api/board/game/%s/move/%s", board_id, move), nil)
	if err != nil {
		return nil, err
	}

	_, err = c.do(req, &ok)
	if err != nil {
		return nil, err
	}

	return ok, err
}
