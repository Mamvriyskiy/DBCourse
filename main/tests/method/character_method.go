package method 

import (
	"crypto/rand"
	"math/big"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"fmt"
)

const (

)

type TestCharacter struct {
	pkg.TypeCharacter
}

func NewCharacter() *TestCharacter {
	var b TestCharacter

	return b.BuilderAccess()
}

func (b *TestCharacter) BuilderAccess() *TestCharacter {
	b.Type = b.generateChar()
	b.UnitMeasure = b.generateChar()

	return b
}

func (b *TestCharacter) generateChar() string {
	chr:= make([]byte, length)
	for j := 0; j < length; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		chr[j] = charset[n.Int64()]
	}

	return string(chr)
}

func (tu TestCharacter) InsertObject(connDB *sqlx.DB) (int, error) {
	var characterID int
	query2 := fmt.Sprintf(`INSERT INTO typecharacter (typecharacter, unitmeasure)
		values ($1, $2) RETURNING typecharacterID`)
	row := connDB.QueryRow(query2, tu.Type, tu.UnitMeasure)
	
	err := row.Scan(&characterID)

	return characterID, err
}