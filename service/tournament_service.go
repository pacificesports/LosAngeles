package service

import (
	"losangeles/model"
	"losangeles/utils"
)

func GetAllTournaments() []model.Tournament {
	var tournaments []model.Tournament
	result := DB.Find(&tournaments)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	return tournaments
}

func GetTournamentByID(id string) model.Tournament {
	var tournament model.Tournament
	result := DB.Where("id = ?", id).First(&tournament)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	return tournament
}

func CreateTournament(tournament model.Tournament) error {
	if DB.Where("id = ?", tournament.ID).Select("*").Updates(&tournament).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New tournament created with id: " + tournament.ID)
		if result := DB.Create(&tournament); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("Tournament with id: " + tournament.ID + " has been updated!")
	}
	return nil
}
