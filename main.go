package main

import (
	//"encoding/json"
	//"fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	pi "github.com/rashidosman/Immunotherapy-Simulation-Study/helpers"
)

// We define the structure and different traits for a patient for the study.
// We use a structure object to do this as it defines a template to create a large number of patients.
// We define the different functions and methods that generate each individual batch of data and then utilize
// these functions to generate a specific trait in the final structure.

func main() {
	// This function initiates the logic of our program and is run every time the main program is executed.
	// We first seed the program to ensure randomness
	rand.Seed(time.Now().UTC().UnixNano())
	// Here we define an array of the three treatments, so that we can later generate
	// three pools of patients for each treatment type.
	var medications [3]string
	medications[0], medications[1], medications[2] = "ipilimumab", "nivolumab", "doxycycline"
	// We simulate the creation of approximately 166-167 patients each
	// assigned to one of the three medications, which gives us an end result of 500.

	// First delete any previous data that may interfere
	if _, err := os.Stat("data"); err == nil {
		fmt.Println("Previous data folder detected - Deleting said old data.")
		os.RemoveAll("data")
	}

	// We create an array of the names of the file names we are going to create for future reference
	// and analysis. We also create a container for all the single patients to become one bigger pool
	// of trial data

	var patientsCollection []*pi.Patient

	trialStorage := make([]string, 3)

	for index := range medications {
		filename := "data/population" + fmt.Sprint(rand.Intn(99)) + ".json"
		pi.OldDataCleanup("data")
		pi.CheckFile(filename)
		trialStorage = append(trialStorage, filename)
		time.Sleep(1 * time.Second)

		for i := 0; i < 166; i++ {
			unparsedPatient := pi.GeneratePatientInfo(medications[index])
			patientsCollection = append(patientsCollection, unparsedPatient)
		}
		parsedPatient, err := json.MarshalIndent(patientsCollection, "", "  ")
		if err != nil {
			fmt.Println(err)
		}

		// Write data of each petient into the JSON data file for the trial

		err = ioutil.WriteFile(filename, parsedPatient, 0644)
		if err != nil {
			fmt.Println(err)
		}

	}
	fmt.Println(trialStorage)
	fmt.Println(fmt.Println("Finished Stage 1: All data populations have been created for the three trials."))

}
