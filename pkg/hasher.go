package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) string {
	salt := make([]byte, 16)

	if _, err := rand.Read(salt); err != nil {
		log.Fatal(err)
	}

	var time uint32 = 1
	var memory uint32 = 1024 * 64
	var threads uint8 = 4
	var keyLen uint32 = 32

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	// Convertimos los bytes binarios a texto seguro con Base64 Raw (sin los '=' del final)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Formato estándar: $argon2id$v=19$m=memoria,t=tiempo,p=hilos$sal$hash
	// Esto es lo que guardarás en el campo u.HashedPassword de tu BD
	hashedPassword := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, time, threads, b64Salt, b64Hash)

	return hashedPassword
}

func VerifyPassword(hashed string, password string) bool {
	// 1. Separar el string guardado usando el delimitador '$'
	parts := strings.Split(hashed, "$")
	if len(parts) != 6 {
		return false // El formato en la BD está corrupto o no es Argon2
	}

	// parts[1] = "argon2id"
	// parts[2] = version ("v=19")
	// parts[3] = parámetros ("m=65536,t=1,p=4")
	// parts[4] = sal en base64
	// parts[5] = hash en base64

	// 2. Extraer los parámetros dinámicamente
	var memory, time uint32
	var threads uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return false
	}

	// 3. Decodificar la sal y el hash original desde Base64 a bytes puros
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false
	}

	originalHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	// 4. Derivar un hash de prueba usando LA MISMA SAL y parámetros extraídos
	verificationHash := argon2.IDKey([]byte(password), salt, time, memory, threads, uint32(len(originalHash)))

	// 5. Comparar en tiempo constante con la librería subtle
	return subtle.ConstantTimeCompare(originalHash, verificationHash) == 1
}
