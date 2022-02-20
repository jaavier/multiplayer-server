package models

type Card struct {
	Id 					string 	 `json:"id"`
	Name 				string 	 `json:"name"`
	Description string 	 `json:"description"`
	Image 			string 	 `json:"image"`
	CreatedAt 	string 	 `json:"createdAt"`
	Price 			int			 `json:"price"`
}
