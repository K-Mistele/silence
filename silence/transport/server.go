package transport

// TransportServer SHOULD RUN ON THE SERVER
type TransportServer struct {
	ClientMap		map[uint16] string
}
