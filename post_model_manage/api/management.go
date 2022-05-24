package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"post_model_manage/model"
)

func GetModelStates(modelName string) ([]*model.ModelState, error) {
	url := "http://211.71.76.189:6081/models/" + modelName
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var modelStates []*model.ModelState
	err = json.Unmarshal(body, &modelStates)
	if err != nil {
		log.Printf("[Error] ummarshal:%v", err)
		return nil, err
	}
	return modelStates,nil
}
