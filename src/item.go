package main

// <model>
type Item struct {
	Id 			string	`json:"id"`
	Category 	string	`json:"category"`
	Name 		string	`json:"name"`
	Quantity 	int		`json:"quantity"`
	Price		float32	`json:"price"`
	Clearance	bool	`json:"clearance"`
}
// </model>
