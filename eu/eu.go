package eu

//CubeParent data
type CubeParent struct {
	Cube CubeTime `xml:"Cube"`
}

//CubeTime data
type CubeTime struct {
	Time string `xml:"time,attr"`
	Cube []Cube `xml:"Cube"`
}

//Cube data
type Cube struct {
	Currency string  `xml:"currency,attr"`
	Rate     float32 `xml:"rate,attr"`
}

//EuroCenterBankXML Europen Bank XMl data structure
type EuroCenterBankXML struct {
	Subject string `xml:"subject"`
	Sender  struct {
		Name string `xml:"name"`
	} `xml:"Sender"`
	Cube CubeParent `xml:"Cube"`
}

const (
	_url = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
)

//Convert interface struct
type Convert struct {
}

func init() {
	//Initialize source data here
	//check cached data
	//check for s3 data
}

//Convert converts from to values
func (eu Convert) Convert(from, to string) (float32, error) {
	//if no data available
	//generate new data

	return 0.0, nil
}

func (eu Convert) generate() error {
	//Generate json file
	//Generate svg -> png file

	return nil
}
