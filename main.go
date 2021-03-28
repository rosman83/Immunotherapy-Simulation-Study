package main

import (
	pi "./helpers"
	"math/rand"
	"time"
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
	for index := range medications {
		for i := 0; i < 16; i++ {
			pi.GeneratePatientInfo(medications[index])

		}
	}
}
