package timer

type Timer interface {
	AddCallBack(*CallBackInfo)
	DelCallBack(*CallBackInfo)
	Run()
}
