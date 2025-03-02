package libbpfgo

const (
	// Maximum number of channels (RingBuffers + PerfBuffers) supported
	maxEventChannels = 512
)

var (
	eventChannels = newRWArray(maxEventChannels)
)
