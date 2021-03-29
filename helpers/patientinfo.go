package patientinfo

import (
	//"encoding/json"
	//"fmt"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Patient struct {
	Name       int    `json:"name"`       // Patients are assigned identifiers instead of direct names for privacy.
	Agegroup   string `json:"agegroup"`   // The general age group of the patient. There is a specified number of patients in each age group.
	Age        int    `json:"age"`        // Patients are assigned an specific age from the above age group.
	Gender     string `json:"gender"`     // Patients are assigned a gender.
	Severity   string `json:"severity"`   // Patients are assigned a severity level for the Melanoma disease. Can be Mild, Moderate, Or Severe.
	Ethnicity  string `json:"ethnicity"`  // Patients are assigned a standard ethnicity.
	Medication string `json:"medication"` // Patients are assigned one of three cancer medications detailed in paper.
}

// We define functions that need more logic to determine
// the chosen trait, such as age that then is classified in
// it's respect age group

func agegroup(age int) (calcAgeGroup string) {
	// 0–14 years old (pediatric group)
	// 15–47 years old (young group)
	// 48–63 years old (middle age group)
	// ≥ 64 years old (elderly group)
	switch {
	case age >= 0 && age <= 14:
		var calcAgeGroup string = "pediatric"
		return calcAgeGroup
	case age >= 15 && age <= 47:
		var calcAgeGroup string = "young"
		return calcAgeGroup
	case age >= 48 && age <= 63:
		var calcAgeGroup string = "middle age"
		return calcAgeGroup
	case age >= 64:
		var calcAgeGroup string = "elderly"
		return calcAgeGroup
	}
	return calcAgeGroup
}

func GeneratePatientInfo(medication string) *Patient {

	// We define arrays of options that don't need the level of
	// logic that require entire functions
	// TODO Decrease chance of certain ethnicities to appear in
	// clinical study to represent real world situations
	var severities [3]string
	var ethnicities [6]string
	var genders [2]string

	severities[0], severities[1], severities[2] = "mild", "moderate", "severe"
	ethnicities[0], ethnicities[1], ethnicities[2] = "American Indian or Alaska Native", "Asian", "Black or African American"
	ethnicities[3], ethnicities[4], ethnicities[5] = "Hispanic or Latino", "Native Hawaiian or Other Pacific Islander", "White"
	genders[0], genders[1] = "male", "female"

	// We proceed with the logic for choosing traits and then
	// sending them back to the main application file
	var p Patient

	p.Name = rand.Int()
	p.Age = rand.Intn(78)
	p.Agegroup = agegroup(p.Age)
	p.Gender = genders[rand.Intn(len(genders))]
	p.Severity = severities[rand.Intn(len(severities))]
	p.Ethnicity = ethnicities[rand.Intn(len(ethnicities))]
	p.Medication = medication

	return &p
}

func OldDataCleanup(directory string) {
	_, err := os.Stat(directory)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(directory, 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
}

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("Created data of population patients for trial named: ", filename)
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
