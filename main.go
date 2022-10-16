package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Status struct {
	Status map[string]int `json:"status"`
}

func main() {
	go func() {
		for range time.Tick(time.Second * 15) {
			WriteJSONFile()
		}
	}()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions, http.MethodPut},
		AllowHeaders:     []string{"Content-Type", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.GET("/status", requestHandler)
	router.Run()
}

func WriteJSONFile() {
	data := Status{
		Status: map[string]int{
			"water": rand.Intn(100),
			"wind":  rand.Intn(100),
		},
	}

	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("status.json", jsonData, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

type DTO struct {
	WaterStatus string `json:"waterStatus"`
	WindStatus  string `json:"windStatus"`
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
}

func requestHandler(ctx *gin.Context) {
	dto := DTO{}
	status := Status{}
	data, err := ioutil.ReadFile("status.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &status)
	if err != nil {
		fmt.Println(err)
	}

	water := status.Status["water"]
	wind := status.Status["wind"]

	if water <= 5 {
		dto.WaterStatus = "Aman"
	} else if 6 <= water && water <= 8 {
		dto.WaterStatus = "Siaga"
	} else {
		dto.WaterStatus = "Bahaya"
	}

	if wind <= 6 {
		dto.WindStatus = "Aman"
	} else if 7 <= wind && wind <= 15 {
		dto.WindStatus = "Siaga"
	} else {
		dto.WindStatus = "Bahaya"
	}

	dto.Water = water
	dto.Wind = wind

	ctx.JSON(http.StatusOK, dto)
}
