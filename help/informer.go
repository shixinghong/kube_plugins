package help

import "k8s.io/client-go/informers"

type podHandler struct {
}

func (p *podHandler) OnAdd(obj interface{}) {
	//TODO implement me
	//panic("implement me")
}

func (p *podHandler) OnUpdate(oldObj, newObj interface{}) {
	//TODO implement me
	//panic("implement me")
}

func (p *podHandler) OnDelete(obj interface{}) {
	//TODO implement me
	//panic("implement me")
}

var fact informers.SharedInformerFactory

func InitCache() {
	fact = informers.NewSharedInformerFactory(clientSet, 0)
	podInformer := fact.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(&podHandler{})
	ch := make(chan struct{})
	fact.Start(ch)
	fact.WaitForCacheSync(ch)
}
