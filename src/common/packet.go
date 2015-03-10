package common

import (
	"bytes"
	//"encoding/binary"
	"encoding/gob"
	//"io"
	//"log"
	//"net"
)

/**
 * 通讯协议:
 * 1. Version - 1 byte;
 * 2. Flags - 3 bytes;
 * 3. Sequence - 8 bytes;
 * 4. Length - 4 bytes;
 * 5. Data Package.
 */
type Packet struct {
	Version  byte
	Flags    uint32
	Sequence uint64
	Length   uint32
	Data     []byte
}

func (p *Packet) Decode(buf []byte) error {

	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(&p)
	if err != nil {
		return err
	}

	return nil

	/*header := make([]byte, PACKET_HEADER_LEN)
	for {
		n, err := io.ReadFull(conn, header)
		if err != nil {
			log.Println("error receiving header, bytes:", n, "reason:", err)
			break
		}
		if n == PACKET_HEADER_LEN {
			log.Println("get packet header")
			break
		}
	}

	p.Version = header[0]
	p.Flags = binary.BigEndian.Uint32(header[1:5])
	p.Sequence = binary.BigEndian.Uint64(header[5:13])
	p.Length = binary.BigEndian.Uint32(header[13:17])

	log.Println("packet version:", p.Version, ", flags:", p.Flags, ", Sequence:", p.Sequence)

	if p.Length > 0 {
		log.Println("ready to get packet data, len is:", p.Length)

		data := make([]byte, p.Length)
		for {
			n, err := io.ReadFull(conn, data)
			if err != nil {
				log.Println("error receiving msg, bytes:", n, "reason:", err)
				break
			}
			if uint32(n) == (p.Length - 1) {
				log.Println("get packet data size %d", n)
				break
			}
		}

		p.Data = data
	}*/
}

/*func (p Packet) Encode() []byte {
	log.Println("pingcmd version is:", p.Version)
	data := make([]byte, p.Length+PACKET_HEADER_LEN)
	data[0] = p.Version
	data[1] = 0x00
	data[2] = 0x00
	data[3] = 0x00
	data[4] = 0x00

	data[5] = 0x00
	data[6] = 0x00
	data[7] = 0x00
	data[8] = 0x00
	data[9] = 0x00
	data[10] = 0x00
	data[11] = 0x00
	data[12] = 0x00

	data[13] = 0x00
	data[14] = 0x00
	data[15] = 0x00
	data[16] = 0x00

	return data
}*/

func (p *Packet) Encode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(p)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
