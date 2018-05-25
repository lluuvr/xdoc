package membership

import (
	"encoding/hex"

	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/storage"
	"github.com/spacelavr/pandora/pkg/utils/crypto/ed25519"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
)

type runtime struct {
	storage *storage.Storage
}

func (r *Runtime) AcceptCandidate(candidate *pb.Candidate) *pb.Member {
	publicKey, privateKey := ed25519.GenerateKeys()

	return &pb.Member{
		Meta: &pb.MemberMeta{
			Email: candidate.Email,
			Name:  candidate.Name,
		},
		PublicKey:  &pb.PublicKey{PublicKey: hex.EncodeToString(publicKey)},
		PrivateKey: privateKey,
	}
}

func (r *Runtime) SignCert(cert *pb.Cert, issuerKey []byte, recipientKey []byte) *pb.Cert {
	buf, _ := hex.DecodeString(cert.String())
	hash := sha256.SumBytes(buf)

	ISignature := ed25519.Sign(hash, issuerKey)
	RSignature := ed25519.Sign(hash, recipientKey)

	cert.Issuer.Signature = ISignature
	cert.Recipient.Signature = RSignature

	return cert
}