package leveldb

import (
	"encoding/hex"
	"fmt"

	"github.com/golang/protobuf/proto"

	"pandora/pkg/pb"
	"pandora/pkg/utils/errors"
)

const (
	PrefixCert = "cert-"
)

// PutCert put cert in leveldb
func (l *Leveldb) PutCert(cert *pb.Cert) error {
	key, _ := hex.DecodeString(fmt.Sprintf("%s%s", PrefixCert, cert.ID))

	buf, err := proto.Marshal(cert)
	if err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(l.db.Put(key, buf, nil))
}

// LoadCert load cert from leveldb
func (l *Leveldb) LoadCert(id string) (*pb.Cert, error) {
	key, _ := hex.DecodeString(fmt.Sprintf("%s%s", PrefixCert, id))

	buf, err := l.db.Get(key, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	cert := &pb.Cert{}

	if err := proto.Unmarshal(buf, cert); err != nil {
		return nil, errors.WithStack(err)
	}

	return cert, nil
}
