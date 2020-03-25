package extractor

func NewExtractor() *Extractor {
	dataCollector := NewDataCollector()
	return &Extractor{dataCollector}
}

type Extractor struct {
	*DataCollector
}

func (e *Extractor) Extract(criteria Criteria) {
	if criteria.Stream {
		e.CollectStream(criteria.Name, criteria.Selection, criteria.Rules)
	} else {
		e.Collect(criteria.Name, criteria.Selection, criteria.Rules)
	}
}
