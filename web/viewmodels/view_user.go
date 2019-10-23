package viewmodels

type ViewUser struct {
	Id			int 	`form:"id"`
	UserName	string  `form:"user_name"`
	BlackTime	string  `form:"black_time"`
	RealName	string  `form:"real_name"`
	Mobile		string  `form:"mobile"`
	Address		string  `form:"address"`
}
