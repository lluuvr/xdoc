package leveldb

import (
	"encoding/hex"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LevelDB struct {
	file string
}

func (ldb *LevelDB) Clean() error {
	if err := os.RemoveAll(ldb.file); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func Put(key string, cert *pb.Cert) error {
	db, err := leveldb.OpenFile(config.Viper.Node.Database.FilePath, nil)
	if err != nil {
		log.Error(err)
		return err
	}
	defer db.Close()

	k, _ := hex.DecodeString(key)

	buf, err := proto.Marshal(cert)
	if err != nil {
		log.Error(err)
		return err
	}

	if err := db.Put(k, buf, nil); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func Get(key string) (*pb.Cert, error) {
	db, err := leveldb.OpenFile(config.Viper.Node.Database.FilePath, &opt.Options{

	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer db.Close()

	k, _ := hex.DecodeString(key)

	buf, err := db.Get(k, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	cert := &pb.Cert{}

	if err := proto.Unmarshal(buf, cert); err != nil {
		log.Error(err)
		return nil, err
	}

	return cert, nil
}