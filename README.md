## Golang Lichess Package [![Go Report Card - Badge]][Go Report Card - Report] [![GitHub Forks - Badge]][GitHub Forks - Link] [![GitHub Issues - Badge]][GitHub Issues - Link]

[Go Report Card - Badge]: https://goreportcard.com/badge/github.com/chrisbarnes2000/lichess
[Go Report Card - Report]: https://goreportcard.com/report/github.com/chrisbarnes2000/lichess
<!-- [![Go Report Card](https://goreportcard.com/badge/github.com/chrisbarnes2000/lichess)](https://goreportcard.com/report/github.com/chrisbarnes2000/lichess) -->

[GitHub Forks - Badge]: https://img.shields.io/github/forks/ChrisBarnes2000/lichess.svg?style=flat-square
[GitHub Forks - Link]: https://github.com/ChrisBarnes2000/lichess/network
<!-- [![GitHub](https://img.shields.io/github/forks/ChrisBarnes2000/lichess.svg?style=flat-square)](https://github.com/ChrisBarnes2000/lichess/network) -->

[GitHub Issues - Badge]: https://img.shields.io/github/issues/ChrisBarnes2000/lichess.svg?style=flat-square
[GitHub Issues - Link]: https://github.com/ChrisBarnes2000/lichess/issues
<!-- [![GitHub](https://img.shields.io/github/issues/ChrisBarnes2000/lichess.svg?style=flat-square)](https://github.com/ChrisBarnes2000/lichess/issues) -->

<!-- [![NPM version](https://img.shields.io/npm/v/chris-barnes-lichess.svg?style=flat)](https://www.npmjs.com/package/chris-barnes-lichess) -->
<!-- [![NPM downloads](https://img.shields.io/npm/dm/chris-barnes-lichess.svg?style=flat)](https://npmjs.org/package/chris-barnes-lichess) -->
<!-- [![Build Status](https://img.shields.io/travis/ChrisBarnes2000/chris-barnes-lichess.svg?style=flat)](https://travis-ci.org/ChrisBarnes2000/chris-barnes-lichess) -->

A Go implementation of [Lichess](https://lichess.org)'s [API](https://lichess.org/api)

### Usage

```go
package main

import (
 "fmt"
 "github.com/chrisbarnes2000/lichess"
 "log"
 "net/http"
 "net/url"
)

func main() {
  baseURL, err := url.Parse("https://lichess.org")
  if err != nil {
    log.Fatal(err)
  }
  client := new(lichess.Client)

  client.BaseURL = baseURL
  client.APIKey = "<API Key>"
  client.UserAgent = "Golang Client"
  client.HttpClient = new(http.Client)

  user, err := client.GetProfile()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(user.ID)
  }
```

## 💻 Local Development

## Getting Started

- First, Create an API Token from [here](https://lichess.org/account/oauth/token/create)
- Second, [fork this repo](https://github.com/chrisbarnes2000/lichess/fork),
- Third, run these commands to clone it locally and get started

```zsh
# Clone and CD into/Open this project
$ git clone git@github.com:YOUR_GITHUB_USERNAME/lichess.git && cd lichess
# Create a .env file to hold our secrets
$ touch .env > LICHESS_TOKEN=YOUR_API_TOKEN
# Download & Install the dependancies. Then Compile the program
$ go build
# Run the program locally
$ go run lichess.go
```

## Available Features


- [x] Accounts <details>
  - [X] GetProfile [API Docs](https://lichess.org/api#operation/accountMe), [location in our code](./account.go)
  - [X] GetEmail [API Docs](https://lichess.org/api#operation/accountEmail), [location in our code](./account.go)
  - [X] GetPreferences [API Docs](https://lichess.org/api#operation/account), [location in our code](./account.go)
  - [X] GetKidModeStatus [API Docs](https://lichess.org/api#operation/accountKid), [location in our code](./account.go)
  - [X] SetKidModeStatus[API Docs](https://lichess.org/api#operation/accountKidPost), [location in our code](./account.go)
  - [X] Unit Testing [location in our code](./account_test.go)
  </details>
- [ ] Users <details>
  - [ ] GetRLUsersStatus [API Docs](https://lichess.org/api#operation/apiUsersStatus), [location in our code](./users.go)
  - [ ] GetTopPlayers [API Docs](https://lichess.org/api#operation/player), [location in our code](./users.go)
  - [ ] GetLeaderboard [API Docs](https://lichess.org/api#operation/playerTopNbPerfType), [location in our code](./users.go)
  - [ ] GetUserPublicData [API Docs](https://lichess.org/api#operation/apiUser), [location in our code](./users.go)
  - [ ] GetPlayerHistory [API Docs](https://lichess.org/api#operation/apiUserRatingHistory), [location in our code](./users.go)
  - [ ] GetPlayerActivity [API Docs](https://lichess.org/api#operation/apiUserActivity), [location in our code](./users.go)
  - [ ] GetPlayerByID [API Docs](https://lichess.org/api#operation/apiUsers), [location in our code](./users.go)
  - [ ] GetTeamMembers [API Docs](https://lichess.org/api#operation/teamIdUsers), [location in our code](./users.go)
  - [ ] GetLiveStream [API Docs](https://lichess.org/api#operation/streamerLive), [location in our code](./users.go)
  - [ ] GetCrossTable [API Docs](https://lichess.org/api#operation/apiCrosstable), [location in our code](./users.go)
  - [ ] Unit Testing [location in our code](./users_test.go)
  </details>
- [ ] Relations <details>
  - [ ] GetFollowing [API Docs](https://lichess.org/api#operation/apiUserFollowing), [location in our code](./relations.go)
  - [ ] GetFollowers [API Docs](https://lichess.org/api#operation/apiUserFollowers), [location in our code](./relations.go)
  - [ ] Unit Testing [location in our code](./relations_test.go)
  </details>
- [ ] Games <details>
  - [ ] ExportOneGame [API Docs](https://lichess.org/api#operation/gamePgn), [location in our code](./game.go)
  - [ ] ExportOngoingGame [API Docs](https://lichess.org/api#operation/apiUserCurrentGame), [location in our code](./game.go)
  - [ ] ExportGames [API Docs](https://lichess.org/api#operation/apiGamesUser), [location in our code](./game.go)
  - [ ] ExportGameByID [API Docs](https://lichess.org/api#operation/gamesExportIds), [location in our code](./game.go)
  - [ ] StreamCurrentGames [API Docs](https://lichess.org/api#operation/gamesByUsers), [location in our code](./game.go)
  - [ ] GetOngoingGames [API Docs](https://lichess.org/api#operation/apiAccountPlaying), [location in our code](./game.go)
  - [ ] GetCurrentTvGames [API Docs](https://lichess.org/api#operation/tvChannels), [location in our code](./game.go)
  - [ ] StreamCurrentTvGame [API Docs](https://lichess.org/api#operation/tvFeed), [location in our code](./game.go)
  - [ ] StreamMovesFromGame [API Docs](https://lichess.org/api#operation/streamGame), [location in our code](./game.go)
  - [ ] ImportOneGame [API Docs](https://lichess.org/api#operation/gameImport), [location in our code](./game.go)
  - [ ] Unit Testing [location in our code](./game_test.go)
  </details>
- [ ] Puzzles <details>
  - [ ] GetDailyPuzzle [API Docs](https://lichess.org/api#operation/apiPuzzleDaily), [location in our code](./puzzles.go)
  - [ ] GetPuzzleActivity [API Docs](https://lichess.org/api#operation/apiPuzzleActivity), [location in our code](./puzzles.go)
  - [ ] GetPuzzleDashboard [API Docs](https://lichess.org/api#operation/apiPuzzleDashboard), [location in our code](./puzzles.go)
  - [ ] GetStormDashboard [API Docs](https://lichess.org/api#operation/apiStormDashboard), [location in our code](./puzzles.go)
  - [ ] Unit Testing [location in our code](./puzzle_test.go)
  </details>
- [ ] Teams <details>
  - [ ] GetTeamSwissTournaments [API Docs](https://lichess.org/api#operation/apiTeamSwiss), [location in our code](./teams.go)
  - [ ] GetSingleTeam [API Docs](https://lichess.org/api#operation/teamShow), [location in our code](./teams.go)
  - [ ] GetPopularTeams [API Docs](https://lichess.org/api#operation/teamAll), [location in our code](./teams.go)
  - [ ] TeamsofPlayer [API Docs](https://lichess.org/api#operation/teamOfUsername), [location in our code](./teams.go)
  - [ ] SearchTeams [API Docs](https://lichess.org/api#operation/teamSearch), [location in our code](./teams.go)
  - [ ] GetMembersofTeam [API Docs](https://lichess.org/api#operation/teamIdUsers), [location in our code](./teams.go)
  - [ ] GetTeamArenaTournaments [API Docs](https://lichess.org/api#operation/apiTeamArena), [location in our code](./teams.go)
  - [ ] JoinTeam [API Docs](https://lichess.org/api#operation/teamIdJoin), [location in our code](./teams.go)
  - [ ] LeaveTeam [API Docs](https://lichess.org/api#operation/teamIdQuit), [location in our code](./teams.go)
  - [ ] KickFromTeam [API Docs](https://lichess.org/api#operation/teamIdKickUserId), [location in our code](./teams.go)
  - [ ] MessageAllTeamMembers [API Docs](https://lichess.org/api#operation/teamIdPmAll), [location in our code](./teams.go)
  - [ ] Unit Testing [location in our code](./teams_test.go)
  </details>
- [ ] Board <details>
  - [ ] StreamIncomingEvents [API Docs](https://lichess.org/api#operation/apiStreamEvent), [location in our code](./board.go)
  - [ ] CreateSeek [API Docs](https://lichess.org/api#operation/apiBoardSeek), [location in our code](./board.go)
  - [ ] StreamBaordGameState/GetBoardGameState [API Docs](https://lichess.org/api#operation/boardGameStream), [location in our code](./board.go)
  - [ ] MakeBoardMove [API Docs](https://lichess.org/api#operation/boardGameMove), [location in our code](./board.go)
  - [ ] WriteInChat [API Docs](https://lichess.org/api#operation/boardGameChat), [location in our code](./board.go)
  - [ ] AbortGame [API Docs](https://lichess.org/api#operation/boardGameAbort), [location in our code](./board.go)
  - [ ] ResignGame [API Docs](https://lichess.org/api#operation/boardGameResign), [location in our code](./board.go)
  - [ ] HandleDrawOffers [API Docs](https://lichess.org/api#operation/boardGameDraw), [location in our code](./board.go)
  - [ ] Unit Testing [location in our code](./board_test.go)
  </details>
- [ ] Chessbot <details>
  - [ ] StreamIncomingEvents [API Docs](https://lichess.org/api#operation/apiStreamEvent), [location in our code](./board.go)
  - [ ] UpgradeToBotAccount [API Docs](https://lichess.org/api#operation/botAccountUpgrade), [location in our code](./bot.go)
  - [ ] StreamBotGameState [API Docs](https://lichess.org/api#operation/botGameStream), [location in our code](./bot.go)
  - [ ] MakeBotMove [API Docs](https://lichess.org/api#operation/botGameMove), [location in our code](./bot.go)
  - [ ] BotWriteInChat [API Docs](https://lichess.org/api#operation/botGameChat), [location in our code](./bot.go)
  - [ ] BotAbortGame [API Docs](https://lichess.org/api#operation/botGameAbort), [location in our code](./bot.go)
  - [ ] BotResginGame [API Docs](https://lichess.org/api#operation/botGameResign), [location in our code](./bot.go)
  - [ ] Unit Testing [location in our code](./chessbot_test.go)
  </details>
- [ ] Challenges <details>
  - [ ] CreateChallenge [API Docs](https://lichess.org/api#operation/challengeCreate), [location in our code](./challenges)
  - [ ] AcceptChallenge [API Docs](https://lichess.org/api#operation/challengeAccept), [location in our code](./challenges)
  - [ ] DeclineChallenge [API Docs](https://lichess.org/api#operation/challengeDecline), [location in our code](./challenges)
  - [ ] CancelChallenge [API Docs](https://lichess.org/api#operation/challengeCancel), [location in our code](./challenges)
  - [ ] ChallengeTheAI [API Docs](https://lichess.org/api#operation/challengeAi), [location in our code](./challenges)
  - [ ] OpenEndedChallenge [API Docs](https://lichess.org/api#operation/challengeOpen), [location in our code](./challenges)
  - [ ] StartClocks [API Docs](https://lichess.org/api#operation/challengeStartClocks), [location in our code](./challenges)
  - [ ] AddToOpponentClock [API Docs](https://lichess.org/api#operation/roundAddTime), [location in our code](./challenges)
  - [ ] CreateAdminChallenge [API Docs](https://lichess.org/api#operation/challengeCreateAdmin), [location in our code](./challenges)
  - [ ] Unit Testing [location in our code](./challenges_test.go)
   </details>
- [ ] Bulk Pairings <details>
  - [ ] ViewUpcomingBulkPairings [API Docs](https://lichess.org/api#operation/bulkPairingGet), [location in our code](./bulkpairing.go)
  - [ ] CreateBulkPairing [API Docs](https://lichess.org/api#operation/bulkPairingCreate), [location in our code](./bulkpairing.go)
  - [ ] ManuallyStartClocks [API Docs](https://lichess.org/api#operation/bulkPairingStartClocks), [location in our code](./bulkpairing.go)
  - [ ] CancelBulkPairing [API Docs](https://lichess.org/api#operation/bulkPairingDelete), [location in our code](./bulkpairing.go)
  - [ ] Unit Testing [location in our code](bulkpairings_test.go)
  </details>
- [ ] Arena Tournaments <details>
  - [ ] GetCurrentArenaTournaments [API Docs](https://lichess.org/api#operation/apiTournament), [location in our code](./arenatournaments.go)
  - [ ] CreateArenaTournament [API Docs](https://lichess.org/api#operation/apiTournamentPost), [location in our code](./arenatournaments.go)
  - [ ] GetArenaTournamentInfo [API Docs](https://lichess.org/api#operation/tournament), [location in our code](./arenatournaments.go)
  - [ ] UpdateArenaTournament [API Docs](https://lichess.org/api#operation/apiTournamentUpdate), [location in our code](./arenatournaments.go)
  - [ ] JoinArenaTournament [API Docs](https://lichess.org/api#operation/apiTournamentJoin), [location in our code](./arenatournaments.go)
  - [ ] TerminateArenaTournament [API Docs](https://lichess.org/api#operation/apiTournamentTerminate), [location in our code](./arenatournaments.go)
  - [ ] UpdateTeamBattle [API Docs](https://lichess.org/api#operation/apiTournamentTeamBattlePost), [location in our code](./arenatournaments.go)
  - [ ] ExportArenaTournamentGames [API Docs](https://lichess.org/api#operation/gamesByTournament), [location in our code](./arenatournaments.go)
  - [ ] GetArenaTournamentResults [API Docs](https://lichess.org/api#operation/resultsByTournament), [location in our code](./arenatournaments.go)
  - [ ] GetTeamStandingFromArenaTournament [API Docs](https://lichess.org/api#operation/teamsByTournament), [location in our code](./arenatournaments.go)
  - [ ] GetArenaTournamentsCreatedByUser [API Docs](https://lichess.org/api#operation/apiUserNameTournamentCreated), [location in our code](./arenatournaments.go)
  - [ ] GetTeamArenaTournaments [API Docs](https://lichess.org/api#operation/apiTeamArena), [location in our code](./arenatournaments.go)
  - [ ] Unit Testing [location in our code](./arenatournaments_test.go)
  </details>
- [ ] Swiss Tournaments <details>
  - [ ] CreateSwissTournament [API Docs](https://lichess.org/api#operation/apiSwissNew), [location in our code](./swisstournaments.go)
  - [ ] TerminateSwissTournament [API Docs](https://lichess.org/api#operation/apiSwissTerminate), [location in our code](./swisstournaments.go)
  - [ ] ExportTRFSwissTournament [API Docs](https://lichess.org/api#operation/swissTrf), [location in our code](./swisstournaments.go)
  - [ ] ExportGameFromSwissTournament [API Docs](https://lichess.org/api#operation/gamesBySwiss), [location in our code](./swisstournaments.go)
  - [ ] GetSwissTournamentResults [API Docs](https://lichess.org/api#operation/resultsBySwiss), [location in our code](./swisstournaments.go)
  - [ ] GetTeamSwissTournaments [API Docs](https://lichess.org/api#operation/apiTeamSwiss), [location in our code](./swisstournaments.go)
  - [ ] Unit Testing [location in our code](./swisstournaments_test.go)
  </details>
- [ ] Simuls <details>
  - [ ] GetCurrentSimuls [API Docs](https://lichess.org/api#operation/apiSimul), [location in our code](./simuls.go)
  - [ ] Unit Testing [location in our code](./simuls_test.go)
  </details>
- [ ] Studies <details>
  - [ ] ExportOneStudyChapter [API Docs](https://lichess.org/api#operation/studyChapterPgn), [location in our code](./studies.go)
  - [ ] ExportAllChapters [API Docs](https://lichess.org/api#operation/studyAllChaptersPgn), [location in our code](./studies.go)
  - [ ] ExportAllStudieFromUser [API Docs](https://lichess.org/api#operation/studyExportAllPgn), [location in our code](./studies.go)
  - [ ] Unit Testing [location in our code](./studies_test.go)
  </details>
- [ ] Messaging <details>
  - [ ] SendPrivateMessage [API Docs](https://lichess.org/api#operation/inboxUsername), [location in our code](./messages.go)
  - [ ] Unit Testing [location in our code](./messages_test.go)
  </details>
- [ ] Broadcasts <details>
  - [ ] GetOfficialBroadcasts [API Docs](https://lichess.org/api#operation/broadcastIndex), [location in our code](./broadcasts.go)
  - [ ] CreateBroadcast [API Docs](https://lichess.org/api#operation/broadcastCreate), [location in our code](./broadcasts.go)
  - [ ] GetBroadcast [API Docs](https://lichess.org/api#operation/broadcastGet), [location in our code](./broadcasts.go)
  - [ ] UpdateBroadcast [API Docs](https://lichess.org/api#operation/broadcastUpdate), [location in our code](./broadcasts.go)
  - [ ] PushPNGToBroadcast [API Docs](https://lichess.org/api#operation/broadcastPush), [location in our code](./broadcasts.go)
  - [ ] Unit Testing [location in our code](broadcasts_test.go)
  </details>
- [ ] Analysis <details>
  - [ ] GetCloudEvaluation [API Docs](https://lichess.org/api#operation/apiCloudEval), [location in our code](./analysis.go)
  - [ ] Unit Testing [location in our code](./analysis_test.go)
  </details>
- [ ] Open Explorer <details>
  - [ ] MasterGames [API Docs](https://lichess.org/api#operation/openingExplorerMaster), [location in our code](./openexplorer.go)
  - [ ] LichessGames [API Docs](https://lichess.org/api#operation/openingExplorerLichess), [location in our code](./openexplorer.go)
  - [ ] OTBMasterGames [API Docs](https://lichess.org/api#operation/openingExplorerMasterGame), [location in our code](./openexplorer.go)
  - [ ] OpeningExplorerStats [API Docs](https://lichess.org/api#operation/openingExplorerStats), [location in our code](./openexplorer.go)
  - [ ] Unit Testing [location in our code](./openexplorer_test.go)
  </details>
- [ ] Tablebase <details>
  - [ ] TablebaseLookup [API Docs](https://lichess.org/api#operation/tablebaseStandard), [location in our code](./tablebase.go)
  - [ ] TablebaseLookupForAtomicChess [API Docs](https://lichess.org/api#operation/tablebaseAtomic), [location in our code](./tablebase.go)
  - [ ] TablebaseLookupForAntiChess [API Docs](https://lichess.org/api#operation/antichessAtomic), [location in our code](./tablebase.go)
  - [ ] Unit Testing [location in our code](./tablebase_test.go)
  </details>
- [ ] OAuth <details>
  <!-- ACCOUNT -->
  - [ ] GetPuzzleActivity [API Docs](https://lichess.org/api#operation/apiPuzzleActivity), [location in our code](./oauth.go)
  - [ ] GetPuzzleDashBoard [API Docs](https://lichess.org/api#operation/apiPuzzleDashboard), [location in our code](./oauth.go)
  - [ ] GetProfile [API Docs](https://lichess.org/api#operation/accountMe), [location in our code](./oauth.go)
  - [ ] GetEmail [API Docs](https://lichess.org/api#operation/accountEmail), [location in our code](./oauth.go)
  - [ ] GetPreferences [API Docs](https://lichess.org/api#operation/account), [location in our code](./oauth.go)
  - [ ] GetKidModeStatus [API Docs](https://lichess.org/api#operation/accountKid), [location in our code](./oauth.go)
  - [ ] SetKidModeStatus [API Docs](https://lichess.org/api#operation/accountKidPost), [location in our code](./oauth.go)
  <!-- GAME -->
  - [ ] ExportGames [API Docs](https://lichess.org/api#operation/apiGamesUser), [location in our code](./oauth.go)
  - [ ] GetOngoingGames [API Docs](https://lichess.org/api#operation/apiAccountPlaying), [location in our code](./oauth.go)
  - [ ] ImportOneGame [API Docs](https://lichess.org/api#operation/gameImport), [location in our code](./oauth.go)
  <!-- STUDIES -->
  - [ ] ExportAllStudieFromUser [API Docs](https://lichess.org/api#operation/studyExportAllPgn), [location in our code](./oauth.go)
  <!-- BRODCASTS -->
  - [ ] CreateBroadcast [API Docs](https://lichess.org/api#operation/broadcastCreate), [location in our code](./oauth.go)
  - [ ] GetBroadcast [API Docs](https://lichess.org/api#operation/broadcastGet), [location in our code](./oauth.go)
  - [ ] UpdateBroadcast [API Docs](https://lichess.org/api#operation/broadcastUpdate), [location in our code](./oauth.go)
  - [ ] PushPNGToBroadcast [API Docs](https://lichess.org/api#operation/broadcastPush), [location in our code](./oauth.go)
  <!-- TEAM -->
  - [ ] JoinTeam [API Docs](https://lichess.org/api#operation/teamIdJoin), [location in our code](./oauth.go)
  - [ ] LeaveTeam [API Docs](https://lichess.org/api#operation/teamIdQuit), [location in our code](./oauth.go)
  - [ ] KickFromTeam [API Docs](https://lichess.org/api#operation/teamIdKickUserId), [location in our code](./oauth.go)
  - [ ] MessageAllTeamMembers [API Docs](https://lichess.org/api#operation/teamIdPmAll), [location in our code](./oauth.go)
  - [ ] Unit Testing [location in our code](./oauth_test.go)
  </details>
- [ ] Client <details>
  - [X] newRequest [API Docs](https://lichess.org/api#section/Introduction/Clients), [location in our code](./client.go)
  - [X] do [API Docs](https://lichess.org/api#section/Introduction/Clients), [location in our code](./client.go)
  - [X] Unit Testing [location in our code](./client_test.go)
  </details>

<details> <summary>Ideas For Tournament Main Code</summary>

Afetr setting things up for the Chess club at MakeSchool to compete and learn from eachother I have noticed these features are missing from the tournament sections of lichess.org and will be looking into how one mighth add these

  1. Ability to Kick/Pause Members
  2. Ability to Invite people you follow
  3. Ability to Pick who you compete against
  4. Ability to Change How Matches are paird up
  5. Ability to Link to Twitter/Twitch/Discord/Slack Group Channel for voice/disscussion
  6. Dispay Tournaments you've joined via quick button on mobile and desktop versions

</details>

## 📚 Resources

- [Go Template CheeatSheet](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet)
- [Intro to go testing](https://tutorialedge.net/golang/intro-testing-in-go/)
- [Intro to go benchmark testing](https://tutorialedge.net/golang/benchmarking-your-go-programs/)
- [Go: tests with HTML coverage report](https://kenanbek.medium.com/go-tests-with-html-coverage-report-f977da09552d)
- [Makefiles for Go Developers](https://tutorialedge.net/golang/makefiles-for-go-developers/)
- [Runes in golang](https://www.geeksforgeeks.org/rune-in-golang/)
- [Default Vaules for Struct Fields in golang](https://www.geeksforgeeks.org/how-to-assign-default-value-for-struct-field-in-golang/)
