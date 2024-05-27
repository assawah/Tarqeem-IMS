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
	"create_project": "Create Project",
	"edit_project":   "Edit Project",
}

var ValidProjectDeliveryStrategies []string = []string{
	"Waterfall",
	"Agile",
	"Lean",
}
var ValidProjectContractingStrategies = []string{
	"Maximum Price",
	"Fixed Price",
	"Unit Price",
}
var ValidProjectTypes []string = []string{
	"Chemical Manufacturing",
	"Stadium Musuem",
	"Dam",
	"Metal refining or processing",
	"Oil exploration or production",
	"Oil refining",
	"Natural gas processing",
	"Highway",
	"Power generation",
	"Water or wastewater",
	"Consumer products manufacturing",
}

var ValidProjectStates []string = []string{
	"Front-end planning",
	"Design",
	"Procurement",
	"Construction",
	"Construction",
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
