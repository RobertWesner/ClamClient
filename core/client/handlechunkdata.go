package client

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/RobertWesner/ClamClient/core/packets"
)

type BlockData struct {
}

type ChunkData []BlockData

func HandleChunkData(packet *packets.Packet51MapChunk) (ChunkData, error) {
	r, err := zlib.NewReader(bytes.NewReader(packet.CompressedData))
	if err != nil {
		return nil, err
	}
	defer func() { _ = r.Close() }()

	foo, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("!!!")
	err = os.WriteFile("dump.bin", foo, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return nil, errors.New("WIP")
}
