package cache

type Stat struct {
	Count	int64
	KeySize	int64
	ValSize	int64
}

func (s *Stat) add(k string, v []byte) {
	s.Count++
	s.KeySize += int64((len(k)))
	s.ValSize += int64((len(v)))
}

func (s *Stat) del(k string, v []byte) {
	s.Count--
	s.KeySize -= int64((len(k)))
	s.ValSize -= int64((len(v)))
}
