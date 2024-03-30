package translate

import "log"

var English map[string]string = map[string]string{
	"emailLoginNote": "Use the same email address you used for registeration.",
	"welcome":        "Welecome to Interface Management System (IMS).",
	"loginHelp":      "Please use your credentials to login or create a new account",
	"regCoordinator": "Register as a coordinator",
	"regMember":      "Register as a team member",
	"title_name":     "IMS",
	"password":       "Password",
	"login":          "Login",
	"email":          "Email address",
}

var ValidProjectDeliverayStratigies []string = []string{}
var ValidProjectContractingStratigies []string
var ValidProjectTypes []string = []string{
	"Chemical manufacturing",
	"Stadium musuem",
	"Dam",
	"Metal refining/processing",
	"Oil exploration/production",
	"Oil refining",
	"Natural gas processing",
	"Highway",
	"Power generation",
	"Water/wastewater",
	"Consumer products manufacturing",
}

var ValidProjectStages []string = []string{
	"Front-end planning",
	"Design",
	"Procurement",
	"Construction",
	"Commissions",
	"Start-up",
	"Completed",
}

func Message(k string) string {
	if val, ok := English[k]; ok {
		return val
	}
	m := "Couldn't find key " + k
	log.Println(m)
	return m
}
