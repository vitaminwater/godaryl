package daryl

type Daryl struct {
}

func darylProcess(d *Daryl) {
	for {

	}
}

func NewDaryl(identifier string) *Daryl {
	d := &Daryl{}
	go darylProcess(d)
	return d
}
