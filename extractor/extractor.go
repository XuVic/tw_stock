package extractor

func NewExtractor() *Extractor {
	dataCollector := make(DataCollector)
	criterias := make(map[string]Criteria)
	return &Extractor{dataCollector, criterias}
}

type Extractor struct {
	Data      DataCollector
	Criterias map[string]Criteria
}

func (e *Extractor) ExtractFrom(criteria Criteria) {
	if _, ok := e.Criterias[criteria.Name]; !ok {
		e.Criterias[criteria.Name] = criteria
	}

	if criteria.Stream {
		e.Data.CollectStream(criteria.Name, criteria.Selection, criteria.Rules)
	} else {
		e.Data.Collect(criteria.Name, criteria.Selection, criteria.Rules)
	}
}

func (e *Extractor) Setup(criterias []Criteria) {
	for _, c := range criterias {
		e.Criterias[c.Name] = c
	}
}

func (e *Extractor) ExtractAll() {
	for _, c := range e.Criterias {
		e.ExtractFrom(c)
	}
}
