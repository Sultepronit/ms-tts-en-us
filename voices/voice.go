package voices

import (
	"fmt"
	"math/rand/v2"
	"slices"
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

func GetRandomVoice(isMale bool, exclude []string) (models.Voice, error) {
	set, err := getVoices(isMale)
	if err != nil {
		return models.Voice{}, err
	}
	fmt.Println(exclude)

	var voice models.Voice
	for range 20 {
		i := rand.IntN(len(set))
		// i := rand.IntN(2)
		fmt.Println(i)
		voice = set[i]
		if !slices.Contains(exclude, voice.CodeName) {
			break
		}
	}

	return voice, err
}
