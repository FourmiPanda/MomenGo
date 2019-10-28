/**
 * ???
 *
 * @description :: ????.
 */

package entities

type Mean struct {
	TemperatureMean float64
	PressureMean    float64
	WindMean        float64
}

func CreateMean(meanT float64, meanP float64, meanW float64) *Mean {
	return &Mean{
		TemperatureMean: meanT,
		PressureMean:    meanP,
		WindMean:        meanW,
	}
}
