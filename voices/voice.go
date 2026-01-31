package voices

import (
	"math/rand/v2"
	"tts/db"
	"tts/models"
)

var males []models.Voice
var females []models.Voice

func getVoices(isMale bool) ([]models.Voice, error) {
	var set *[]models.Voice
	if isMale {
		set = &males
	} else {
		set = &females
	}
	if len(*set) > 0 {
		return *set, nil
	}
	return db.SelectVoices(isMale)
}

func GetRandomVoice(isMale bool) (models.Voice, error) {
	set, err := getVoices(isMale)
	if err != nil {
		return models.Voice{}, err
	}
	i := rand.IntN(len(set))
	return set[i], err
}
