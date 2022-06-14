package repository

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns"
	"github.com/go-acme/lego/v4/registration"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

func TestACME(t *testing.T) {
	m := &autocert.Manager{
		Prompt:          autocert.AcceptTOS,
		Cache:           autocert.DirCache("./certs"),
		HostPolicy:      autocert.HostWhitelist("example.com"), // Replace with your domain.
		RenewBefore:     0,
		Client:          nil,
		Email:           "",
		ExtraExtensions: nil,
	}
	cfg := &tls.Config{
		GetCertificate: m.GetCertificate,
		NextProtos: []string{
			"http/1.1", acme.ALPNProto,
		},
	}
	ln, lnErr := tls.Listen("tcp", "0.0.0.0:8080", cfg) /* #nosec G102 */
	if lnErr != nil {
		panic(lnErr)
	}
	err := http.Serve(ln, &Handler{})
	if err != nil {
		panic(err)
	}
}

type Handler struct {
	config *tls.Config
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(time.Now().String()))
	p := h.config.Certificates[0].Certificate[0]
	crt, _ := x509.ParseCertificate(p)
	fmt.Println(crt.Subject, crt.NotAfter)
	writer.Write([]byte(crt.Subject.String()))
	writer.Write([]byte(crt.NotAfter.String()))
	cert, _ := tls.LoadX509KeyPair("G:/ssl/server.crt", "G:/ssl/server.key")
	h.config.Certificates[0] = cert
}

type MyUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *MyUser) GetEmail() string {
	return u.Email
}
func (u MyUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *MyUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func TestLego(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Error(err)
		return
	}

	myUser := MyUser{
		Email: "you@yours.com",
		key:   privateKey,
	}
	config := lego.NewConfig(&myUser)
	// This CA URL is configured for a local dev instance of Boulder running in Docker in a VM.
	config.Certificate.KeyType = certcrypto.RSA2048

	// A client facilitates communication with the CA server.
	client, err := lego.NewClient(config)
	if err != nil {
		t.Error(err)
		return
	}

	// 用户登录名称 dns@1368854082260567.onaliyun.com
	//AccessKey ID LTAI5tJrqXZSeFkKNBSWScQV
	//AccessKey Secret wB6tdMyTy14gWM5tZXmmMfyyQrx6kl
	os.Setenv("ALICLOUD_ACCESS_KEY", "LTAI5tJrqXZSeFkKNBSWScQV")
	os.Setenv("ALICLOUD_SECRET_KEY", "wB6tdMyTy14gWM5tZXmmMfyyQrx6kl")
	//os.Setenv("", "")
	provider, err := dns.NewDNSChallengeProviderByName("alidns")
	if err != nil {
		t.Error(err)
		return
	}
	err = client.Challenge.SetDNS01Provider(provider)
	if err != nil {
		t.Error(err)
		return
	}
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		t.Error(err)
		return
	}
	myUser.Registration = reg
	request := certificate.ObtainRequest{
		Domains: []string{"*.aacfactory.com"},
		Bundle:  true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		t.Error(err)
		return
	}

	// Each certificate comes back with the cert bytes, the bytes of the client's
	// private key, and a certificate URL. SAVE THESE TO DISK.
	fmt.Printf("%#v\n", certificates)
	resp, err := http.Get(certificates.CertStableURL)
	if err != nil {
		t.Error(err)
		return
	}
	certPEM, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	ioutil.WriteFile("G:/acme.cert", certPEM, 0600)
	ioutil.WriteFile("G:/acme.key", certificates.PrivateKey, 0600)
}

func TestLegoResult(t *testing.T) {
	cert, _ := tls.LoadX509KeyPair("G:/acme.cert", "G:/acme.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	ln, lnErr := tls.Listen("tcp", ":8080", config)
	if lnErr != nil {
		t.Error(lnErr)
		return
	}
	err := http.Serve(ln, &Handler{config: config})

	//err := http.ListenAndServeTLS(":8080", "G:/acme.cert", "G:/acme.key", &Handler{})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestParseCERT(t *testing.T) {
	p, _ := ioutil.ReadFile("G:/acme.cert")
	block, _ := pem.Decode(p)
	cert, pErr := x509.ParseCertificate(block.Bytes)
	if pErr != nil {
		t.Error(pErr)
		return
	}
	fmt.Println(cert.Subject.CommonName)
	fmt.Println(cert.NotBefore, cert.NotAfter)
}

func TestUDP(t *testing.T) {
	fmt.Println(net.ResolveUDPAddr("udp", ":https"))
	ss := []int{1, 2, 3, 4, 5, 6}
	ssLen := len(ss)
	for i := ssLen; i > 0; i-- {
		fmt.Println(ss[i-1])
	}
}

func TestPrivateNet(t *testing.T) {
	s := "123.168.0.1"
	ip := net.ParseIP(s)
	fmt.Println(ip.IsPrivate())
	fmt.Println(&M{})
}

type M struct {
	s []string
	r sync.Map
}
