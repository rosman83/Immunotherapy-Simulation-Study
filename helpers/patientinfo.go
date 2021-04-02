package patientinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"reflect"

	wr "github.com/mroth/weightedrand"
)

type Patient struct {
	Name          int    `json:"name"`          // Patients are assigned identifiers instead of direct names for privacy.
	Agegroup      string `json:"agegroup"`      // The general age group of the patient. There is a specified number of patients in each age group.
	Age           int    `json:"age"`           // Patients are assigned an specific age from the above age group.
	Gender        string `json:"gender"`        // Patients are assigned a gender.
	Severity      string `json:"severity"`      // Patients are assigned a severity level for the Melanoma disease. Can be Mild, Moderate, Or Severe.
	Ethnicity     string `json:"ethnicity"`     // Patients are assigned a standard ethnicity.
	Medication    string `json:"medication"`    // Patients are assigned one of three cancer medications detailed in paper.
	AdverseEvents bool   `json:"adverseevents"` // Is there a prescence of adverse events
	Fatality      string `json:"fatality"`
	FirstCycle    string `json:"firstcycle"`
	SecondCycle   string `json:"secondcycle"`
	ThirdCycle    string `json:"thirdcycle"`
	FourthCycle   string `json:"fourthcycle"`
}

func MapRandomKeyGet(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()

	return keys[rand.Intn(len(keys))].Interface()
}

func optionGetter(filename string, rarity string) (chosenOption string, chosensubOption string) {
	// Logic for opening and parsing the data file
	adversityData, _ := os.Open(filename)

	byteValue, _ := ioutil.ReadAll(adversityData)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	rarityoptions := result[rarity].(map[string]interface{})

	// Get randomized adverse event
	chosenOption = fmt.Sprint(MapRandomKeyGet(rarityoptions).(string))
	chosensubOption = fmt.Sprint(rarityoptions[chosenOption])

	return chosenOption, chosensubOption
}

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

// Temporary Statistic Counters

func GeneratePatientInfo(medication string) *Patient {

	// We define functions that need more logic to determine the chosen
	// trait, such as age that then is classified in it's respect age group

	// First define the general probabilities

	// 65.82% of immunotherapy patients developed any adverse effects.
	immunoprob, _ := wr.NewChooser(
		wr.Choice{Item: "adverse", Weight: 23},
		wr.Choice{Item: "none", Weight: 77},
	)
	// 85.19% of chemotherapy patients developed any adverse effects.
	chemoprob, _ := wr.NewChooser(
		wr.Choice{Item: "adverse", Weight: 38},
		wr.Choice{Item: "none", Weight: 62},
	)
	// Fatality probabilities for Immunotherapy
	immunofatal, _ := wr.NewChooser(
		wr.Choice{Item: "fatality", Weight: 87},
		wr.Choice{Item: "none", Weight: 9913},
	)
	// Fatality probabilities for Chemotherapy
	chemofatal, _ := wr.NewChooser(
		wr.Choice{Item: "fatality", Weight: 128},
		wr.Choice{Item: "none", Weight: 9872},
	)

	// Secondly the chemotherapy probabilities

	// Chemotherapy Pediatric/Elderly Male Asian = [47,23,20]
	chemo1, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 47},
		wr.Choice{Item: "lesscommon", Weight: 23},
		wr.Choice{Item: "rare", Weight: 20},
	)
	// Chemotherapy Pediatric/Elderly Male Black = [47,24,19]
	chemo2, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 47},
		wr.Choice{Item: "lesscommon", Weight: 24},
		wr.Choice{Item: "rare", Weight: 19},
	)
	// Chemotherapy Pediatric/Elderly Male White = [47,25,18]
	chemo3, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 47},
		wr.Choice{Item: "lesscommon", Weight: 25},
		wr.Choice{Item: "rare", Weight: 18},
	)
	// Chemotherapy Pediatric/Elderly Female Asian = [47,22,21]
	chemo4, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 47},
		wr.Choice{Item: "lesscommon", Weight: 22},
		wr.Choice{Item: "rare", Weight: 21},
	)
	// Chemotherapy Pediatric/Elderly Female Black = [47,23,20]
	chemo5, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 47},
		wr.Choice{Item: "lesscommon", Weight: 23},
		wr.Choice{Item: "rare", Weight: 20},
	)
	// Chemotherapy Pediatric/Elderly Female White = [47,24,19]
	chemo6, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 47},
		wr.Choice{Item: "lesscommon", Weight: 24},
		wr.Choice{Item: "rare", Weight: 19},
	)
	// Chemotherapy Middle Age/Young Male Asian = [47,26,17]
	chemo7, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 47},
		wr.Choice{Item: "lesscommon", Weight: 26},
		wr.Choice{Item: "rare", Weight: 17},
	)
	// Chemotherapy Middle Age/Young Male Black = [47,27,16]
	chemo8, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 47},
		wr.Choice{Item: "lesscommon", Weight: 27},
		wr.Choice{Item: "rare", Weight: 16},
	)
	// Chemotherapy Middle Age/Young Male White = [47,28,15]
	chemo9, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 47},
		wr.Choice{Item: "lesscommon", Weight: 28},
		wr.Choice{Item: "rare", Weight: 15},
	)
	// Chemotherapy Middle Age/Young Female Asian = [46,26,18]
	chemo10, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 46},
		wr.Choice{Item: "lesscommon", Weight: 26},
		wr.Choice{Item: "rare", Weight: 18},
	)
	// Chemotherapy Middle Age/Young Female Black = [46,27,17]
	chemo11, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 46},
		wr.Choice{Item: "lesscommon", Weight: 27},
		wr.Choice{Item: "rare", Weight: 17},
	)
	// Chemotherapy Middle Age/Young Female White = [46,28,16]
	chemo12, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 46},
		wr.Choice{Item: "lesscommon", Weight: 28},
		wr.Choice{Item: "rare", Weight: 16},
	)

	// Thirdly the immunotherapy probabilities

	// Immunotherapy Pediatric/Elderly Male Asian = [51,43,6]
	immuno1, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 51},
		wr.Choice{Item: "lesscommon", Weight: 43},
		wr.Choice{Item: "rare", Weight: 6},
	)
	// Immunotherapy Pediatric/Elderly Male Black = [51,44,5]
	immuno2, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 51},
		wr.Choice{Item: "lesscommon", Weight: 44},
		wr.Choice{Item: "rare", Weight: 5},
	)
	// Immunotherapy Pediatric/Elderly Male White = [51,45,4]
	immuno3, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 51},
		wr.Choice{Item: "lesscommon", Weight: 45},
		wr.Choice{Item: "rare", Weight: 4},
	)
	// Immunotherapy Pediatric/Elderly Female Asian = [50,43,7]
	immuno4, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 50},
		wr.Choice{Item: "lesscommon", Weight: 43},
		wr.Choice{Item: "rare", Weight: 7},
	)
	// Immunotherapy Pediatric/Elderly Female Black = [50,44,6]
	immuno5, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 50},
		wr.Choice{Item: "lesscommon", Weight: 44},
		wr.Choice{Item: "rare", Weight: 6},
	)
	// Immunotherapy Pediatric/Elderly Female White = [50,45,5]
	immuno6, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 50},
		wr.Choice{Item: "lesscommon", Weight: 45},
		wr.Choice{Item: "rare", Weight: 5},
	)
	// Immunotherapy Middle Age/Young Male Asian = [53,43,4]
	immuno7, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 53},
		wr.Choice{Item: "lesscommon", Weight: 43},
		wr.Choice{Item: "rare", Weight: 4},
	)
	// Immunotherapy Middle Age/Young Male Black = [53,44,3]
	immuno8, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 53},
		wr.Choice{Item: "lesscommon", Weight: 44},
		wr.Choice{Item: "rare", Weight: 3},
	)
	// Immunotherapy Middle Age/Young Male White = [53,45,2]
	immuno9, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 53},
		wr.Choice{Item: "lesscommon", Weight: 44},
		wr.Choice{Item: "rare", Weight: 2},
	)
	// Immunotherapy Middle Age/Young Female Asian = [52,42,5]
	immuno10, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 52},
		wr.Choice{Item: "lesscommon", Weight: 42},
		wr.Choice{Item: "rare", Weight: 5},
	)
	// Immunotherapy Middle Age/Young Female Black = [52,44,4]
	immuno11, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 52},
		wr.Choice{Item: "lesscommon", Weight: 44},
		wr.Choice{Item: "rare", Weight: 4},
	)
	// Immunotherapy Middle Age/Young Female White = [52,45,3]
	immuno12, _ := wr.NewChooser(
		wr.Choice{Item: "morecommon", Weight: 52},
		wr.Choice{Item: "lesscommon", Weight: 45},
		wr.Choice{Item: "rare", Weight: 3},
	)

	// We define arrays of options that don't need the level of
	// logic that require entire functions
	// TODO Decrease chance of certain ethnicities to appear in
	// clinical study to represent real world situations
	var severities [3]string
	var ethnicities [4]string
	var genders [2]string

	severities[0], severities[1], severities[2] = "mild", "moderate", "severe"
	ethnicities[0], ethnicities[1], ethnicities[2], ethnicities[3] = "Asian", "Black or African American", "Hispanic or Latino", "White"
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

	// Here we begin logic for other factors that are not automatically generated from
	// the start - such as the actual trial results

	var events [4]string
	var eventvalues [4]string
	var fatalResults [1]string
	var AdversityResults bool

	// CHEMOTHERAPY LOGIC
	for i := 0; i < 4; i++ {
		if p.Medication == "doxycycline" {
			result1 := chemoprob.Pick().(string)
			if result1 == "adverse" {
				AdversityResults = true
				if p.Gender == "male" {
					if p.Agegroup == "pediatric" || p.Agegroup == "elderly" {
						if p.Ethnicity == "Black or African American" || p.Ethnicity == "Hispanic or Latino" {
							res1 := chemo2.Pick().(string)
							switch res1 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else if p.Ethnicity == "Asian" {
							res2 := chemo1.Pick().(string)
							switch res2 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else {
							res3 := chemo3.Pick().(string)
							switch res3 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						}

					} else {
						if p.Ethnicity == "Black or African American" || p.Ethnicity == "Hispanic or Latino" {
							res4 := chemo8.Pick().(string)
							switch res4 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else if p.Ethnicity == "Asian" {
							res5 := chemo7.Pick().(string)
							switch res5 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else {
							res6 := chemo9.Pick().(string)
							switch res6 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						}
					}
				} else {
					if p.Agegroup == "pediatric" || p.Agegroup == "elderly" {
						if p.Ethnicity == "Black or African American" || p.Ethnicity == "Hispanic or Latino" {
							res7 := chemo5.Pick().(string)
							switch res7 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else if p.Ethnicity == "Asian" {
							res8 := chemo4.Pick().(string)
							switch res8 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else {
							res9 := chemo6.Pick().(string)
							switch res9 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						}

					} else {
						if p.Ethnicity == "Black or African American" || p.Ethnicity == "Hispanic or Latino" {
							res10 := chemo11.Pick().(string)
							switch res10 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else if p.Ethnicity == "Asian" {
							res11 := chemo10.Pick().(string)
							switch res11 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else {
							res12 := chemo12.Pick().(string)
							switch res12 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						}
					}
				}
			} else {
				events[i] = "None"
			}
		} else {
			result2 := immunoprob.Pick().(string)
			if result2 == "adverse" {
				AdversityResults = true
				if p.Gender == "male" {
					if p.Agegroup == "pediatric" || p.Agegroup == "elderly" {
						if p.Ethnicity == "Black or African American" || p.Ethnicity == "Hispanic or Latino" {
							res1 := immuno2.Pick().(string)
							switch res1 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else if p.Ethnicity == "Asian" {
							res2 := immuno1.Pick().(string)
							switch res2 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else {
							res3 := immuno3.Pick().(string)
							switch res3 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						}

					} else {
						if p.Ethnicity == "Black or African American" || p.Ethnicity == "Hispanic or Latino" {
							res4 := immuno8.Pick().(string)
							switch res4 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else if p.Ethnicity == "Asian" {
							res5 := immuno7.Pick().(string)
							switch res5 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else {
							res6 := immuno9.Pick().(string)
							switch res6 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						}
					}
				} else {
					if p.Agegroup == "pediatric" || p.Agegroup == "elderly" {
						if p.Ethnicity == "Black or African American" || p.Ethnicity == "Hispanic or Latino" {
							res7 := immuno5.Pick().(string)
							switch res7 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else if p.Ethnicity == "Asian" {
							res8 := immuno4.Pick().(string)
							switch res8 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else {
							res9 := immuno6.Pick().(string)
							switch res9 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						}

					} else {
						if p.Ethnicity == "Black or African American" || p.Ethnicity == "Hispanic or Latino" {
							res10 := immuno11.Pick().(string)
							switch res10 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else if p.Ethnicity == "Asian" {
							res11 := immuno10.Pick().(string)
							switch res11 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						} else {
							res12 := immuno12.Pick().(string)
							switch res12 {
							case "morecommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "lesscommon":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "morecommon")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							case "rare":
								chosenResult, chosensubResult := optionGetter("supplementary/adverse.json", "rare")
								events[i] = chosenResult
								eventvalues[i] = chosensubResult
							}
						}
					}
				}
			} else {
				events[i] = "None"
			}
		}
	}

	switch AdversityResults {
	case true:
		p.AdverseEvents = true
	case false:
		p.AdverseEvents = false
	}

	switch p.Medication {
	case "doxycycline":
		res1 := chemofatal.Pick().(string)
		fatalResults[0] = res1
	default:
		res2 := immunofatal.Pick().(string)
		fatalResults[0] = res2
	}

	p.Fatality = fatalResults[0]
	// IMMUNOTHERAPY LOGIC

	p.FirstCycle = events[0]
	p.SecondCycle = events[1]
	p.ThirdCycle = events[2]
	p.FourthCycle = events[3]

	// End of the final logic, here we return the patient to be saved in data files

	// Announce basic statistics, and announce completion

	return &p
}

func OldDataCleanup(directory string) {
	_, err := os.Stat(directory)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(directory, 0755)
		if errDir != nil {
			fmt.Println("There was an error cleaning up old data.")
			log.Fatal(err)
		}

	}
}

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			fmt.Println("There was an error creating JSON data.")
			return err
		}
	}
	return nil
}
