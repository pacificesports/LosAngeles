package service

import (
	"encoding/json"
	"io"
	"losangeles/model"
	"losangeles/utils"
	"net/http"
)

func GetAllTeamsForTournament(tournamentID string) []model.TournamentTeam {
	var tournamentTeams []model.TournamentTeam
	result := DB.Where("tournament_id = ?", tournamentID).Find(&tournamentTeams)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	for i := range tournamentTeams {
		tournamentTeams[i].Team = FetchTeam(tournamentTeams[i].TeamID)
	}
	return tournamentTeams
}

func GetTeamForTournament(tournamentID string, teamID string) model.TournamentTeam {
	var tournamentTeam model.TournamentTeam
	result := DB.Where("tournament_id = ? AND team_id = ?", tournamentID, teamID).First(&tournamentTeam)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	if tournamentTeam.TeamID != "" {
		tournamentTeam.Team = FetchTeam(teamID)
	}
	return tournamentTeam
}

func SetTeamForTournament(tournamentTeam model.TournamentTeam) error {
	result := DB.Create(&tournamentTeam)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func RemoveTeamFromTournament(tournamentID string, teamID string) error {
	result := DB.Where("tournament_id = ? AND team_id = ?", tournamentID, teamID).Delete(&model.TournamentTeam{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTournamentsForTeam(teamID string) []model.Tournament {
	var tournamentTeams []model.TournamentTeam
	var tournaments []model.Tournament
	result := DB.Where("team_id = ?", teamID).Find(&tournamentTeams)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	for i := range tournamentTeams {
		tournaments = append(tournaments, GetTournamentByID(tournamentTeams[i].TournamentID))
	}
	return tournaments
}

func FetchTeam(teamID string) json.RawMessage {
	var responseJson json.RawMessage = []byte("{}")
	mappedService := MatchRoute("teams", "-")
	if mappedService.ID != 0 {
		proxyClient := &http.Client{}
		//proxyRequest, _ := http.NewRequest("GET", "http://localhost"+":"+strconv.Itoa(mappedService.Port)+"/schools/"+schoolID, nil) // Use this when not running in Docker
		proxyRequest, _ := http.NewRequest("GET", mappedService.URL+"/teams/"+teamID, nil)
		proxyRequest.Header.Set("Request-ID", "-")
		proxyResponse, err := proxyClient.Do(proxyRequest)
		if err != nil {
			utils.SugarLogger.Errorln("Failed to get team information from " + mappedService.Name + ": " + err.Error())
			return responseJson
		}
		defer proxyResponse.Body.Close()
		proxyResponseBodyBytes, _ := io.ReadAll(proxyResponse.Body)
		json.Unmarshal(proxyResponseBodyBytes, &responseJson)
	}
	return responseJson
}
