package InteractionLayer

import "fmt"

type InteractionLayerExecute struct {
	printedStrings []string
}

func (l *InteractionLayerExecute) Print(text string) {
	fmt.Println(text)
	l.printedStrings = append(l.printedStrings, text)
}

func (l *InteractionLayerExecute) GetOutput() []string {
	return l.printedStrings
}
