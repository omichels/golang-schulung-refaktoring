package numbergen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestNewGenerator(t *testing.T) {
	g := GeneratorInstance()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {

		go func(ix int) {
			time.Sleep(time.Millisecond * 60)
			fmt.Println(ix, " = ", g.GetNumber())
			wg.Done()
		}(i)
	}
	wg.Add(100)
	wg.Wait()

}

func TestBeingBad(t *testing.T) {

	g := &Generator{}
	g.n = 77

	// panics with nil pointer
	// assert.Equal(t, 77, g.GetNumber())
	GeneratorInstance().GetNumber()
	GeneratorInstance().GetNumber()
	// 3 as result is unpredictable because of other tests in project
	// either 3 or 103 is a good guess
	// setting to 78 will never work
	assert.True(t, GeneratorInstance().GetNumber() != 78)
	// assert.Equal(t, 1, GeneratorInstance().GetNumber())

	// fails because access to g.n does not work / has no effect
	//assert.Equal(t, 78, GeneratorInstance().GetNumber())
}
