package handler





type RegisterUser struct {
	// Name is the first name of the user.
	Name string `json:"name,omitempty"`
	
	// Email is the email of the user.
	Email string `json:"email,omitempty"`
}


type UpdateUser struct {
		// Name is the first name of the user.
	Name string `json:"name,omitempty"`
	
	// Email is the email of the user.
	Email string `json:"email,omitempty"`
}


