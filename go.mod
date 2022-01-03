module puppet-steady-state

go 1.14

require (
	github.com/Jeffail/gabs v1.4.0 // indirect
	github.com/akira/go-puppetdb v0.0.0-20200122132916-4bc34e483a6e
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b
)

replace k8s.io/client-go => k8s.io/client-go v0.17.4
