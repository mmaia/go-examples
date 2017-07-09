package nutrilife

type Patient struct {
	Name 	string 	`json: "name"`
	Email 	string	`json: "email"`
}


type Patients struct {
	patientList []Patient
}