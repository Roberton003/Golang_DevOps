package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func DecryptLargefiles(fileInput, fileOutput string, key []byte) error {
	infile, err := os.Open(fileInput)
	if err != nil {
		return err
	}
	defer infile.Close()
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	fi, err := infile.Stat()
	if err != nil {
		return err
	}
	iv := make([]byte, block.BlockSize())
	msgLen := fi.Size() - int64(len(iv))
	_, err = infile.ReadAt(iv, msgLen)
	if err != nil {
		log.Fatal("erro lendo o iv: ", err)
	}

	outfile, err := os.OpenFile(fileOutput, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer outfile.Close()
	buf := make([]byte, 4096)
	stream := cipher.NewCTR(block, iv)
	for {
		n, err := infile.Read(buf)
		if n > 0 {
			if n > int(msgLen) {
				n = int(msgLen)
			}
			msgLen -= int64(n)
			stream.XORKeyStream(buf, buf[:n])
			outfile.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("tamanho de bytes lidos: %d", n)
			break
		}
	}
	return nil
}

func EncryptLargeFiles(infile, outfile string, key []byte) error {
	buf := make([]byte, 4096)
	in, err := os.Open(infile)
	if err != nil {
		return err
	}

	defer in.Close()
	out, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer out.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}
	stream := cipher.NewCTR(block, iv)
	for {
		n, err := in.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			out.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	out.Write(iv)
	return nil
}
