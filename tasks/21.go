package tasks

import "fmt"

/*Шаблон адаптера состоит из четырех компонентов:
Целевой интерфейс: интерфейс, с которым клиент ожидает взаимодействовать.
Адаптер: объект, реализующий целевой интерфейс и обертывающий адаптируемый объект.
Адаптируемый объект: интерфейс, который необходимо адаптировать для использования клиентом.
Клиент: компонент, использующий целевой интерфейс.
*/
type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

type adapteeImpl struct{}

func (a *adapteeImpl) SpecificRequest() string {
	return "Adaptee request"
}

type adapter struct {
	adaptee Adaptee
}

func (a *adapter) Request() string {
	return "Adapter: " + a.adaptee.SpecificRequest()
}

func L1_21() {
	adaptee := &adapteeImpl{}
	target := &adapter{
		adaptee: adaptee,
	}

	fmt.Println(target.Request())
}
