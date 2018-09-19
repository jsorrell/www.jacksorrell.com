package templateexecuter

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/jsorrell/www.jacksorrell.com/data"
	"github.com/jsorrell/www.jacksorrell.com/log"
)

// TODO store catching options on creation

const DefaultContentType = "text/html; charset=utf-8"

var bufPool *sync.Pool

func init() {
	bufPool = &sync.Pool{
		New: func() interface{} {
			return pooledBuffer{new(bytes.Buffer)}
		},
	}
}

func (g *TemplateGroup) checkInit() {
	if !g.inited {
		g.Init()
		log.Warn("Template group " + g.name + " not inited. Initing...")
	}
}

type ServerPush struct {
	Path string
	As   string
}

type ReadSeekerCloser interface {
	io.Reader
	io.Seeker
	io.Closer
}

type ByteReadSeekerCloser struct {
	*bytes.Reader
}

func (ByteReadSeekerCloser) Close() error {
	return nil
}

type pooledBuffer struct {
	*bytes.Buffer
}

func (b pooledBuffer) getReader() BufferReadSeekerCloser {
	return BufferReadSeekerCloser{ByteReadSeekerCloser{bytes.NewReader(b.Bytes())}, b}
}

type BufferReadSeekerCloser struct {
	ByteReadSeekerCloser
	buf pooledBuffer
}

func (b BufferReadSeekerCloser) Close() error {
	return b.buf.Close()
}

func (r pooledBuffer) Close() error {
	r.Reset()
	bufPool.Put(r)
	return nil
}

func addTemplate(tmpl **template.Template, templateName string, templateFile string) error {
	templateString, err := readTemplateString(templateFile)
	if err != nil {
		return err
	}
	*tmpl, err = (*tmpl).New(templateName).Parse(templateString)
	return err
}

func readTemplateString(filename string) (string, error) {
	templateFile, err := data.Templates.Open(filename)
	if err != nil {
		return "", err
	}
	defer templateFile.Close()
	return data.ReadFileToString(templateFile)
}

func (tmpl *DynamicTemplate) ServeHTTP(w http.ResponseWriter, req *http.Request, args interface{}) {
	r := tmpl.GetReader(args)
	defer r.Close()
	// TODO fully disabling caching will work for current static site, but should probably be smarter at some point in future
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	w.Header().Set("Content-Type", tmpl.GetContentType())
	if len(tmpl.pushes) > 0 {
		w.Header().Set("Link", getServerPushHeader(tmpl.pushes))
	}
	http.ServeContent(w, req, "", time.Unix(0, 0), r)
}

func (tmpl *DynamicTemplate) SetContentType(t string) *DynamicTemplate {
	tmpl.contentType = t
	return tmpl
}

func (tmpl *DynamicTemplate) AddServerPushes(pushes ...ServerPush) *DynamicTemplate {
	tmpl.pushes = append(tmpl.pushes, pushes...)
	return tmpl
}

func (tmpl *DynamicTemplate) GetContentType() string {
	return tmpl.contentType
}

func (tmpl *StaticTemplate) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r, etag := tmpl.GetReader()
	defer r.Close()
	w.Header().Set("Cache-Control", "max-age=86400, public")
	w.Header().Set("Etag", etag)
	w.Header().Set("Content-Type", tmpl.GetContentType())
	if len(tmpl.pushes) > 0 {
		w.Header().Set("Link", getServerPushHeader(tmpl.pushes))
	}
	http.ServeContent(w, req, "", time.Unix(0, 0), r)
}

func (tmpl *StaticTemplate) SetContentType(t string) *StaticTemplate {
	tmpl.contentType = t
	return tmpl
}

func (tmpl *StaticTemplate) AddServerPushes(pushes ...ServerPush) *StaticTemplate {
	tmpl.pushes = append(tmpl.pushes, pushes...)
	return tmpl
}

func (tmpl *StaticTemplate) GetContentType() string {
	return tmpl.contentType
}

func genEtag(data []byte) string {
	sum := md5.Sum(data)
	b64 := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(sum[:])
	return "\"" + b64 + "\""
}

func getServerPushHeader(pushes []ServerPush) string {
	builder := strings.Builder{}
	buffer := new(bytes.Buffer)

	for _, push := range pushes {
		fmt.Fprintf(buffer, "<%s>; rel=preload; as=%s, ", push.Path, push.As)
		builder.Write(buffer.Bytes())
		buffer.Reset()
	}

	return builder.String()[:builder.Len()-2] // Remove final ", "
}
