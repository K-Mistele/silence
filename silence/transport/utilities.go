package transport

import "bytes"

func datagramsAreEqual(d1 *Datagram, d2 *Datagram) bool {

	if d1.EndpointID != d2.EndpointID ||
			d1.SequenceNumber != d2.SequenceNumber ||
			d1.AckNumber != d2.AckNumber ||
			d1.Flags != d2.Flags ||
			d1.Checksum != d2.Checksum ||
			d1.FragmentNumber != d2.FragmentNumber ||
			d1.Reserved != d2.Reserved ||
			bytes.Compare(d1.Data, d2.Data) != 0 {
		return false
	} else {
		return true
	}
}

