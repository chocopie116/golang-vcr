package recorder

import (
	"net/http"
	"fmt"
	"encoding/gob"
	"os"
)
type Casette struct {
	Encoder gob.Encoder
	Decoder gob.Decoder
}

func NewCasette(path string) (Casette , error){
	file, err := os.Create(path)
	if err != nil{
		return nil, err

	}
	en := gob.Encoder{}
	return Casette{}
}

func (c *Casette) Put (req http.Request) (){

}
func (c *Casette) load() ([]http.Request, error){

}
