package repository

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/binary"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/json"
	"github.com/cnf/structhash"
	"github.com/valyala/fasthttp"
	"github.com/vmihailenco/msgpack/v5"
	"math/big"
	"net"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
	"unsafe"
)

type X struct {
	Id string
}

func TestName(t *testing.T) {

	v := make([]*PostRow, 0, 1)

	rt := reflect.ValueOf(v)
	x := rt.Interface()

	fmt.Println(&x)
	fmt.Println(fff(&x))
	fmt.Println(x)

	s := "sss"
	fmt.Println(fmt.Sprintf("%x", md5.Sum([]byte(s))))

	z := X{Id: "sss"}
	p, _ := json.Marshal(z)
	fmt.Println(string(p))

}

func fff(v interface{}) (err error) {
	rv := reflect.Indirect(reflect.ValueOf(v))
	if reflect.TypeOf(rv.Interface()).Kind() != reflect.Slice {
		err = fmt.Errorf("fns SQL Rows: scan failed for target elem is not slice or struct")
		return
	}
	var elemType reflect.Type
	elemIsPtr := false
	rvt := reflect.TypeOf(rv.Interface())
	elem := rvt.Elem()
	fmt.Println(elem)
	if elem.Kind() == reflect.Ptr {
		if elem.Elem().Kind() != reflect.Struct {
			err = fmt.Errorf("fns SQL Rows: scan failed for element of target is not struct or ptr of struct")
			return
		}
		elemIsPtr = true
		elemType = elem.Elem()
	} else if elem.Kind() == reflect.Struct {
		elemIsPtr = false
		elemType = elem
	} else {
		err = fmt.Errorf("fns SQL Rows: scan failed for element of target is not struct or ptr of struct")
		return
	}
	rv0 := reflect.MakeSlice(rvt, 0, 1)
	x := reflect.New(elemType)
	x.Elem().FieldByName("Id").SetString("id")
	if elemIsPtr {
		rv0 = reflect.Append(rv0, x)
	} else {
		rv0 = reflect.Append(rv0, x.Elem())
	}

	fmt.Println(rv0.Interface())
	rv.Set(rv0)
	return
}

func xxx() *UserRow {
	fmt.Println("xxx")
	return &UserRow{
		Id: "fff",
	}
}

func TestXXX(t *testing.T) {
	p := &PostRow{
		Id:       "x",
		CreateBY: "",
		CreateAT: time.Time{},
		ModifyBY: "",
		ModifyAT: time.Time{},
		Version:  0,
		Title:    "",
		Content:  "",
		Author:   nil,
		Likes:    0,
		Comments: nil,
	}
	rv := reflect.Indirect(reflect.ValueOf(p))
	fv := rv.FieldByName("Author")
	ptr := reflect.ValueOf(xxx).Pointer()
	fmt.Println(ptr)
	x := reflect.NewAt(reflect.TypeOf(xxx), unsafe.Pointer(ptr)).Elem()
	fmt.Println(x.CanInterface(), x.CanAddr(), x.CanConvert(reflect.TypeOf(&UserRow{})))

	fv.Set(x.Elem().Call(nil)[0])
	fmt.Println("!")
	fmt.Println(p)
	a := p.Author
	fmt.Println("!")
	a.Id = ""
	fmt.Println("!")
	fmt.Println(a.Id)

}

func TestHostClient(t *testing.T) {
	client := &fasthttp.HostClient{
		Addr:  "www.runoob.com",
		Name:  "fns-proxy",
		IsTLS: true,
	}
	req := fasthttp.AcquireRequest()
	//req.UseHostHeader = true
	req.Header.SetHost("www.runoob.com")
	err := client.DoTimeout(req, fasthttp.AcquireResponse(), 5*time.Second)
	fmt.Println(err)
}

type Context interface {
	context.Context
	D() map[string]string
}

type ctx struct {
	context.Context
}

func (c ctx) D() map[string]string {
	return c.Value("a").(map[string]string)
}

func TestCtx(t *testing.T) {
	c1 := context.TODO()
	c2 := context.WithValue(c1, "a", 1)
	n := c2.Value("a").(int) + 1
	c2 = context.WithValue(c1, "a", n)
	c3 := context.WithValue(c2, "a", n)
	fmt.Println(c1.Value("a"), c2.Value("a"), c3.Value("a"))
	u, err := url.Parse("https://www.dd/asdf asdf ")
	fmt.Println(err, u.Scheme, u.Host, u.Path)
	fmt.Println("base64:123"[7:])
	remoteIp := "125.22.212.144:2020"
	fmt.Println(remoteIp[0:strings.Index(remoteIp, ":")])

}

func TestTimeoutCTx(t *testing.T) {
	r := context.WithValue(context.Background(), "now", time.Now())
	p, pc := context.WithTimeout(r, 2*time.Second)
	c, cc := context.WithTimeout(p, 1*time.Second)
	go func(p context.Context, c context.Context) {
		pd := false
		cd := false
		for {
			select {
			case <-p.Done():
				if !pd {
					fmt.Println("p", p.Value("now"), time.Now())
					pd = true
				}
			case <-c.Done():
				if !cd {
					fmt.Println("c", p.Value("now"), time.Now())
					cd = true
				}
			}
			if pd && cd {
				break
			}
		}
	}(p, c)
	time.Sleep(5 * time.Second)
	cc()
	pc()
}

func TestCreateSSL(t *testing.T) {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	subject := pkix.Name{
		Country:            []string{"CN"},
		Organization:       []string{"FNS"},
		OrganizationalUnit: []string{"AACFACTORY"},
		SerialNumber:       "",
		CommonName:         "app title",
	}

	certificate := x509.Certificate{
		SerialNumber: serialNumber,
		Issuer:       pkix.Name{},
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 1024)

	//
	derBytes, _ := x509.CreateCertificate(rand.Reader, &certificate, &certificate, &pk.PublicKey, pk)
	certOut := bytes.NewBufferString("")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyOut := bytes.NewBufferString("")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})

	fmt.Println(certOut.String())
	fmt.Println(keyOut.String())
}

func TestHost(t *testing.T) {
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname, _ = os.LookupEnv("HOSTNAME")
	}
	if hostname == "" {
		return
	}
	fmt.Println(hostname)
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return
	}
	host := ""
	for _, ip := range ips {
		if ip.IsGlobalUnicast() {
			host = ip.To4().String()
			break
		}
	}
	fmt.Println(host)
}

func TestScanner(t *testing.T) {
	s := &Scanner{v: &ABC{
		Id: 1,
	}}

	abc := &ABC{}
	fmt.Println(s.Scan(abc))
	fmt.Println(abc.Id)
}

type Scanner struct {
	v interface{}
}

func (s *Scanner) Scan(dest interface{}) (err error) {
	if dest == nil {
		err = fmt.Errorf("dest is nil, dest = %v", dest)
		return
	}
	dpv := reflect.ValueOf(dest)
	if dpv.Kind() != reflect.Ptr {
		err = fmt.Errorf("dest's type is not Ptr, dest = %v", dpv.Kind())
		return
	}
	sv := reflect.ValueOf(s.v)
	dv := reflect.Indirect(dpv)
	if sv.Kind() == reflect.Ptr {
		sv = sv.Elem()
	}
	if sv.IsValid() && sv.Type().AssignableTo(dv.Type()) {
		dv.Set(sv)
		return
	}
	if dv.Kind() == sv.Kind() && sv.Type().ConvertibleTo(dv.Type()) {
		dv.Set(sv.Convert(dv.Type()))
		return
	}
	err = fmt.Errorf("scan failed, not match, src = %v, dest = %v", sv.Type(), dv.Type())
	return
}

type ABC struct {
	Id int
}

func TestBig(t *testing.T) {
	p := make([]byte, 8)
	binary.BigEndian.PutUint64(p, uint64(9999999999999999))
	fmt.Println(len(p), string(p), binary.BigEndian.Uint64(p))
	p, _ = json.Marshal(json.RawMessage([]byte{'{', '}'}))
	fmt.Println(string(p))
	fmt.Println(time.Now().Add(-24 * time.Hour))
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = structhash.Hash(&ABC{
			Id: 1,
		}, 1)
	}
}

func BenchmarkMsgPack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(&ABC{
			Id: 1,
		})
	}
}

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&ABC{
			Id: 1,
		})
	}
}

func BenchmarkGOB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = gob.NewEncoder(bytes.NewBufferString("")).Encode(&ABC{
			Id: 1,
		})
	}
}

func TestJsonBytes(t *testing.T) {
	p := json.RawMessage([]byte("{\"a\":1}"))
	b, err := json.Marshal(p)
	fmt.Println(string(b), err)
}

func TestErrorJson(t *testing.T) {
	e := errors.Warning("fns: fooo").WithMeta("a", "a").WithCause(fmt.Errorf("bar"))
	p, _ := json.Marshal(e)
	err := errors.Empty()
	de := json.Unmarshal(p, err)
	fmt.Println(de)
	fmt.Println(fmt.Sprintf("%+v", err))
}

func TestSearch(t *testing.T) {
	s := []string{"0", "4", "3", "2", "2"}
	fmt.Println(s)

	x := sort.SearchStrings(s, "2")
	fmt.Println(x, x < len(s))
	fmt.Println(sort.SearchStrings(s, "5") < len(s))
}
