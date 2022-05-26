package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"time"

	c "AutodeleteV2/src/config"
)

func main() {
	fmt.Println("Input a date as the ending value of the range of delete function")
	//Get User Input
	var uInput string
	fmt.Scan(&uInput)

	//call function in main
	deleter(&uInput)
}

func deleter(endDate *string) error {
	//define default directory
	//directory := filepath.Join("./kiplebox", "vzcenter_release", "ubuntu64", "result")

	//Set file name of the configurations file
	viper.SetConfigName("config")
	//Set the path to look for the configurations file
	viper.AddConfigPath(".")
	//Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configurations c.Configurations

	if err := viper.ReadInConfig(); err != err {
		fmt.Printf("Error reading config file, %v", err)
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	var directory string
	directory = configurations.EPATH

	//raed all camera in directory
	cameras, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal("err1")
	}

	//loop through the datefolders
	for _, camera := range cameras {

		cameraDirectory := directory + "/" + camera.Name()
		dates, dateErr := ioutil.ReadDir(cameraDirectory)
		if dateErr != nil {
			log.Fatal("date error")
		}

		for _, date := range dates {

			parsedEndDate, error := time.Parse("2006-01-02", *endDate)
			if error != nil {
				log.Fatal("err4")
			}

			parsedCurrentDate, error1 := time.Parse("2006-01-02", date.Name())
			if error1 != nil {
				log.Fatal("err5")
			}

			if parsedEndDate.Before(parsedCurrentDate) {
				break
			}

			//insert date into new directory path
			var picDirectory string = cameraDirectory + "/" + date.Name()
			pictures, err := ioutil.ReadDir(picDirectory)
			if err != nil {
				log.Fatal("err2")
			}
			for _, pic := range pictures {
				picDir := picDirectory + "/" + pic.Name()
				switch pic.Name()[:4] {
				case "http":
					err := os.Remove(picDir)
					if err != nil {
						log.Fatal("err3")
					}
					fmt.Println(picDir, "deleted")
				default:
					fmt.Println(picDir + "is locally stored.")
				}
			}
			fmt.Println(picDirectory, "processed.")
		}
	}

	return nil
}
