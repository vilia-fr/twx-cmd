package api

type Api struct {
    User string
    Password string
    AppKey string
    BaseUrl string
}

func (this *Api) NewApi() string {
    return "ok"
}

