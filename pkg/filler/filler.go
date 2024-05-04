package filler

import (
	"encoding/json"
	"fmt"

	"github.com/kirillmc/data_filler/pkg/model"
)

const (
	oneCount    = 1
	smallCount  = 5
	normalCount = 11
	bigCount    = 31

	picturesCount = 5
)

func FindByteSizeOfJson(programs model.TrainPrograms) (int64, error) {
	numOfSets, err := json.Marshal(programs)
	if err != nil {
		return 0.0, fmt.Errorf("fail to get json: %v", err)
	}

	return int64(len(numOfSets)), nil
}

func CreateEmptySetOfPrograms() model.TrainPrograms {
	return model.TrainPrograms{}
}

func CreateSetOfOneProgram() model.TrainPrograms {
	return model.TrainPrograms{
		TrainPrograms: fillTrainPrograms(oneCount),
	}
}

func CreateSmallSetOfPrograms() model.TrainPrograms {
	return model.TrainPrograms{
		TrainPrograms: fillTrainPrograms(smallCount),
	}
}

func CreateNormalSetOfPrograms() model.TrainPrograms {
	return model.TrainPrograms{
		TrainPrograms: fillTrainPrograms(normalCount),
	}
}

func CreateBigSetOfPrograms() model.TrainPrograms {
	return model.TrainPrograms{
		TrainPrograms: fillTrainPrograms(bigCount),
	}
}

func CreateOwnSetOfPrograms(count int) model.TrainPrograms {
	return model.TrainPrograms{
		TrainPrograms: fillTrainPrograms(count),
	}
}

func fillTrainPrograms(count int) []model.TrainProgram {
	var setOfTrainPrograms []model.TrainProgram
	for i := 1; i < count+1; i++ {
		trainPrograms := model.TrainProgram{
			Id:          int64(i),
			ProgramName: fmt.Sprintf("Название программы тренировок №%d", i),
			Description: fmt.Sprintf("Описание программы тренирово №%d", i),
			Status:      "На модерации",
			TrainDays:   fillTrainDays(count),
		}

		setOfTrainPrograms = append(setOfTrainPrograms, trainPrograms)
	}

	return setOfTrainPrograms
}

func fillTrainDays(count int) []model.TrainDay {
	var setOfTrainDays []model.TrainDay
	for i := 1; i < count+1; i++ {
		trainDay := model.TrainDay{
			Id:          int64(i),
			DayName:     fmt.Sprintf("Название тренировочного дня №%d", i),
			Description: fmt.Sprintf("Описание тренировочного дня №%d", i),
			Exercises:   fillExercises(count),
		}

		setOfTrainDays = append(setOfTrainDays, trainDay)
	}

	return setOfTrainDays
}

func fillExercises(count int) []model.Exercise {
	var setOfExercises []model.Exercise
	for i := 1; i < count+1; i++ {
		exercise := model.Exercise{
			Id:           int64(i),
			ExerciseName: fmt.Sprintf("Название упражнения №%d", i),
			Pictures:     fillPictures(picturesCount),
			Description:  fmt.Sprintf("Описание упражнения №%d", i),
			Sets:         fillSets(count),
		}

		setOfExercises = append(setOfExercises, exercise)
	}

	return setOfExercises
}

func fillPictures(count int) []string {
	var setOfPictures []string
	for i := 0; i < count; i++ {
		setOfPictures = append(setOfPictures, fmt.Sprintf("http://picture_%d.png", i))
	}

	return setOfPictures
}

func fillSets(count int) []model.Set {
	var setOfSets []model.Set
	for i := 1; i < count+1; i++ {
		set := model.Set{
			Id:       int64(i),
			Quantity: 11,
			Weight:   55.5,
		}

		setOfSets = append(setOfSets, set)
	}

	return setOfSets
}
