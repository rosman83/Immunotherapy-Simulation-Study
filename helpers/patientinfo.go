package patientinfo

import (
	"math/rand"
)

type patient struct {
	name       int    // Patients are assigned identifiers instead of direct names for privacy.
	agegroup   string // The general age group of the patient. There is a specified number of patients in each age group.
	age        uint16 // Patients are assigned an specific age from the above age group.
	gender     string // Patients are assigned a gender.
	severity   string // Patients are assigned a severity level for the Melanoma disease. Can be Mild, Moderate, Or Severe.
	ethnicity  string // Patients are assigned a standard ethnicity.
	medication string // Patients are assigned one of three cancer medications detailed in paper.
}

// We define functions that need more logic to determine
// the chosen trait, such as age that then is classified in
// it's respect age group

func age() {

}

func agegroup() {

}

func GeneratePatientInfo(medication string) *patient {

	// We define arrays of options that don't need the level of
	// logic that require entire functions

	var severities [3]string
	var ethnicities [6]string
	var genders [2]string

	severities[0], severities[1], severities[2] = "mild", "moderate", "severe"
	ethnicities[0], ethnicities[1], ethnicities[2] = "American Indian or Alaska Native", "Asian", "Black or African American"
	ethnicities[3], ethnicities[4], ethnicities[5] = "Hispanic or Latino", "Native Hawaiian or Other Pacific Islander", "White"
	genders[0], genders[1] = "male", "female"
	// We proceed with the logic for choosing traits and then
	// sending them back to the main application file
	var p patient

	p.name = rand.Int()
	p.agegroup = "32-223"
	p.age = 12
	p.gender = genders[rand.Intn(len(genders))]
	p.severity = severities[rand.Intn(len(severities))]
	p.ethnicity = ethnicities[rand.Intn(len(ethnicities))]
	p.medication = medication
	return &p
}
