package report

type Handler struct {
	token string
}

func Init(account string) Handler {
	return Handler{
		token: accounts,
	}
}

func (rh Handler) getToken() string {
	return rh.token
}

func (rh Handler) GetCostReport() {

}
