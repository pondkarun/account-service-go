package Services

import (
	"fmt"
	"log"
)

type AccountModel struct {
	Id        string  `json:"id"`
	IdAccount string  `json:"id_account"`
	IdCard    string  `json:"id_card"`
	Fname     string  `json:"fname"`
	Lname     string  `json:"lname"`
	Balance   float32 `json:"balance"`
}

func FindAllAccounts() []AccountModel {
	db := Condb()
	rows, err := db.Query("SELECT * FROM public.accounts ORDER BY id ASC")

	if err != nil {
		log.Fatal("query error")
	}
	data := []AccountModel{}
	for rows.Next() {
		acc := AccountModel{}
		err := rows.Scan(&acc.Id, &acc.IdAccount, &acc.IdCard, &acc.Fname, &acc.Lname, &acc.Balance)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		data = append(data, acc)
	}
	return data
}
