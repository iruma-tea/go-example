package main

func scoreboardManager(in <-chan func(map[string]int), done <-chan struct{}) {
	scorebord := map[string]int{}
	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(scorebord)
		}
	}
}

type ChannelScoreboardManager chan func(map[string]int)

func NewChannelScoreboardManager() (ChannelScoreboardManager, func()) {
	ch := make(ChannelScoreboardManager)
	done := make(chan struct{})
	go scoreboardManager(ch, done)
	return ch, func() {
		close(done)
	}
}

func (csm ChannelScoreboardManager) Update(name string, val int) {
	csm <- func(m map[string]int) {
		m[name] = val
	}
} //listend2

func (csm ChannelScoreboardManager) Read(name string) (int, bool) { //liststart3
	var out int
	var ok bool
	done := make(chan struct{})
	csm <- func(m map[string]int) {
		out, ok = m[name]
		close(done)
	}
	<-done
	return out, ok
} //listend3
