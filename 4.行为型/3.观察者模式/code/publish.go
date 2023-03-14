package my_observer

type IPublish interface {
	Register(IObserver)
	RemoveRegister()
	Publish(string)
}

type Publisher struct {
	obsList []IObserver
}

func NewPublisher() IPublish {
	return &Publisher{}
}

func (pub *Publisher) Register(obs IObserver) {
	pub.obsList = append(pub.obsList, obs)
}
func (pub *Publisher) RemoveRegister() {
	pub.obsList = pub.obsList[1:]
}

func (pub *Publisher) Publish(message string) {
	for _, obs := range pub.obsList {
		obs.Update(message)
	}
}
