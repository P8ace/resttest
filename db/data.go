package db

type mockdata struct {
	Sn    uint   `json:"sn"`
	Make  string `json:"make"`
	Model string `json:"model"`
}

var Data = []mockdata{
	{Sn: 1, Make: "Honda", Model: "CB100"},
	{Sn: 2, Make: "Honda", Model: "CB100"},
	{Sn: 3, Make: "Honda", Model: "CB100"},
	{Sn: 4, Make: "Honda", Model: "CB100"},
	{Sn: 5, Make: "Honda", Model: "CB100"},
	{Sn: 6, Make: "Honda", Model: "CB100"},
	{Sn: 7, Make: "Honda", Model: "CB100"},
	{Sn: 8, Make: "Honda", Model: "CB100"},
	{Sn: 9, Make: "Honda", Model: "CB100"},
	{Sn: 10, Make: "Honda", Model: "CB100"},
}
