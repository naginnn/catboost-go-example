package main

import (
	"fmt"
	"github.com/naginnn/catboostgo/wrapper/catboost"
	"slices"
)

func main() {
	// only cbm format
	model, err := catboost.Load("testmodel")
	if err != nil {
		fmt.Println(err)
	}

	var param catboost.ClassParams
	// get metadata from model
	err = model.GetMetaData("class_params", &param)
	if err != nil {
		return
	}
	fmt.Println(param.ClassToLabel)
	fmt.Println(param.ClassNames)
	fmt.Println(model.GetFloatFeaturesCount())
	fmt.Println(model.GetCatFeaturesCount())
	fmt.Println(model.GetDimensionsCount())
	floats := [][]float32{{
		5.0, 0.0, 0.0, 1973.0, 0.0,
		2048755.0, 9.0, 12.0, 431.0, 21103.0,
		21088.0, 15.0, 0.0, 0.0, 0.0,
		2048929.0, 22728486.0, 12.0, 0.0, 0.0,
		22289201.0, 0.0, 42875644.0, 58761330.0,
		45063584.0, 0.0, 0.0}}
	cats := [][]string{}
	prediction, err := model.CalcModelPredictionProba(floats, cats)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result", prediction)
	fmt.Println(slices.Max(prediction))
	defer model.Free()
}
