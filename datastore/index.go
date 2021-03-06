package datastore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	//"math/rand"
	//"fmt"
	//"errors"
	//bson "launchpad.net/mgo/bson"
)

func init() {
	/*
		var pos, length uint64 = 0, 0
		bFile, err := os.Create("/tmp/idx.bson"); if err != nil { fmt.Println(err.Error()) }
		jFile, err := os.Create("/tmp/idx.json"); if err != nil { fmt.Println(err.Error()) }

		//checksum := "716484d000da82756fe593bcec43213f" 
		idx := Index{Filename : "test", CType: "None", Idx: RecordIndex{}}
		for i := 0; i < 1024000; i++ {
			length = uint64(rand.Int63n(10240))
			idx.Idx = append(idx.Idx, Record{pos,length})
			pos = pos+length
		}

		b, err := bson.Marshal(idx); if err != nil { fmt.Println(err.Error()) }
		bFile.Write(b)
		j, err := json.Marshal(idx); if err != nil { fmt.Println(err.Error()) }
		jFile.Write(j)
		fmt.Println("done create")
		bFile.Close()
		jFile.Close()

		bf, err := ioutil.ReadFile("/tmp/idx.json"); if err != nil { fmt.Println(err.Error()) }
		idx = Index{}
		err = bson.Unmarshal(bf, &idx); if err != nil { fmt.Println(err.Error()) }
		fmt.Println("done")
		panic("this is a panic folks")
	*/
}

/*	
Shock Index format:
<position> - unsigned 64bit int
<length>   - unsigned 64bit int
<checksum> - optional (none,md5,sha1,sha256)

#filename=<filename>:checksum=<type>\n
<position><length><checksum><position><length><checksum>...

Json representation:
{
	index_type : <type>,
	filename : <filename>,
	checksum_type : <type>,
	version : <version>,
	index : [
		[<position>,<length>,<optional_checksum>]...
	]
}
*/

type Index struct {
	Type     string      `bson:"index_type" json:"index_type"`
	Filename string      `bson:"filename" json:"filename"`
	CType    string      `bson:"checksum_type" json:"checksum_type"`
	Idx      RecordIndex `bson:"index" json:"index"`
	Version  int         `bson:"version" json:"version"`
}

type RecordIndex []Record

type Record []interface{}

func NewIndex() *Index {
	return &Index{Filename: "", CType: "", Idx: RecordIndex{}}
}

func (idx *Index) Save(filename string) (err error) {
	jFile, err := os.Create(filename)
	if err != nil {
		return
	}
	defer jFile.Close()
	j, err := json.Marshal(idx)
	if err != nil {
		return
	}
	jFile.Write(j)
	return
}

func (idx *Index) Load(filename string) (err error) {
	bf, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(bf, idx)
	return
}
