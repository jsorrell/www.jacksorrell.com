package io

import (
	"bytes"
	"io"
	"os"
	"sync"

	"github.com/jsorrell/www.jacksorrell.com/data"
)

// ReadSeekerCloser is a wrapper interface for io.Reader, io.Seeker, and io.Closer
type ReadSeekerCloser interface {
	io.Reader
	io.Seeker
	io.Closer
}

// RSCloseWrapper wraps an io.ReaderSeeker to add a Close function that does nothing.
type RSCloseWrapper struct {
	io.ReadSeeker
}

// Close does nothing
func (RSCloseWrapper) Close() error {
	return nil
}

// ByteBufPool a pool of ByteBufs that wraps sync.Pool
type ByteBufPool struct {
	pool sync.Pool
}

// ByteBuf is a buffer that implements io.Writer that comes from a ByteBufPool.
type ByteBuf struct {
	*bytes.Buffer
	bbp *ByteBufPool
}

// ByteBufReader allows reading from a ByteBuf.
type ByteBufReader struct {
	*bytes.Reader
	buf ByteBuf
}

// CreateByteBufPool Creates a ByteBufPool. Each created ByteBuf will have the capacity of initialCapacity (but will expand as needed).
func CreateByteBufPool(initialCapacity int) *ByteBufPool {
	pool := &ByteBufPool{pool: sync.Pool{}}
	pool.pool.New = func() interface{} {
		return ByteBuf{bytes.NewBuffer(make([]byte, 0, initialCapacity)), pool}
	}
	return pool
}

// Get retrieves a ByteBuf from the pool. See sync.Pool.Get()
func (p *ByteBufPool) Get() ByteBuf {
	return p.pool.Get().(ByteBuf)
}

// GetReader gets a ByteBufReader from the ByteBuf. Writing to ByteBuf is undefined after this is called.
func (b ByteBuf) GetReader() ByteBufReader {
	return ByteBufReader{Reader: bytes.NewReader(b.Bytes()), buf: b}
}

// Close readds the ByteBuf to the pool. Any additional use of this ByteBuf is undefined.
func (b ByteBuf) Close() error {
	b.Reset()
	b.bbp.pool.Put(b)
	return nil
}

// Close readds the ByteBuf to the pool. Any additional use of this Reader is undefined.
func (r ByteBufReader) Close() error {
	r.buf.Reset()
	r.buf.bbp.pool.Put(r.buf)
	return nil
}

// WriteAssetToDisk copies a file from assets to os filesystem.
func WriteAssetToDisk(src, dst string) error {
	in, err := data.Assets.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}
