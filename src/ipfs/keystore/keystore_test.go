package keystore

import (
	"fmt"
	"io/ioutil"
	"sort"
    "math/rand"
	"testing"
    ci "github.com/libp2p/go-libp2p-crypto"
)

type rr struct{}

func (rr rr) Read(b []byte) (int, error){
	return rand.Read(b) 
}

func privKeyOrFatal(t *testing.T) ci.PrivKey{
	priv, _, err := ci.GenerateEd25519Key(rr{})
	if err != nil{
		t.Fatal(err)
	}
	return priv
}

func TestKeystoreBasic(t *testing.T){
	tdir, err := ioutil.TempDir("/tmp", "keystore-test")
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println("------->", tdir)
	ks, err := NewFSKeystore(tdir)
	if err != nil{
		t.Fatal(err)
	}
	l, err := ks.List()
	if err != nil{
		t.Fatal(err)
	}

	if len(l) != 0 {
		t.Fatal("存在不该有的密钥")
	}
	k1 := privKeyOrFatal(t)
	k2 := privKeyOrFatal(t)
	k3 := privKeyOrFatal(t)
	k4 := privKeyOrFatal(t)

	err = ks.Put("magic_1", k1)
	if err != nil{
		t.Fatal(err)
	}
	
	err = ks.Put("magic_2", k2)
	if err != nil{
		t.Fatal(err)
	}

	l, err = ks.List()
	if err != nil{
		t.Fatal(err)
	}
	sort.Strings(l)
	if l[0] != "magic_1" || l[2] != "magic_2"{
		t.Fatal("密钥列表异常")
	}
	if err := assertDirContents(tdir, []string{"magic_1", "magic_2"}); err != nil{
		t.Fatal(err)
	}

	err = ks.Put("magic_1", k3)
	if err != nil{
		t.Fatal("不能重写密钥")
	}
	if err := assertDirContents(tdir, []string{"magic_1", "magic_2"}); err != nil{
		t.Fatal(err)
	}

	exist, err := ks.Has("magic_1")
	if !exist{
		t.Fatal("密钥magic_1应该存在")
	}
	if err != nil{
		t.Fatal(err)
	}
	exist, err = ks.Has("xxxx")
	if exist{
		t.Fatal("密钥xxxx不应该存在")
	}
	if err != nil{
		t.Fatal(err)
	}
	if err := ks.Delete("magic_1"); err != nil{
		t.Fatal(err)
	}

	if err := assertDirContents(tdir, []string{"magic_2"}); err != nil{
		t.Fatal(err)
	}

	if err := ks.Put("magic_3", k3); err != nil{
		t.Fatal(err)
	}
	if err := ks.Put("magic_4", k4); err != nil{
		t.Fatal(err)
	}

	if err := assertDirContents(tdir, []string{"magic_2", "magic_3", "magic_4"}); err != nil{
		t.Fatal(err)
	}
	if err := assertGetKey(ks, "magic_2", k2); err != nil{
		t.Fatal(err)
	}
	if err := assertGetKey(ks, "magic_3", k3); err != nil{
		t.Fatal(err)
	}
	if err := assertGetKey(ks, "magic_4", k4); err != nil{
		t.Fatal(err)
	}

}

func assertDirContents(dir string, exp []string) error{
	finfos, err := ioutil.ReadDir(dir)
	if err != nil{
		return err
	}

	if len(finfos) != len(exp){
		return fmt.Errorf("excepted %d directory entries", len(exp))
	}
	var names []string
	for _, fi := range finfos{
		names = append(names, fi.Name())
	}
	
	sort.Strings(names)
	sort.Strings(exp)
	if len(names) != len(exp){
		return fmt.Errorf("文件列表数量异常")
	}
	for i, v := range names{
		if v != exp[i]{
			return fmt.Errorf("文件有异常的输入")
		}
	}
	return nil
}

func assertGetKey(ks Keystore, name string, exp ci.PrivKey) error{
	out_k, err := ks.Get(name)
	if err != nil{
		return err
	}

	if !out_k.Equals(exp){
		return fmt.Errorf("获取到的密钥跟预期的不匹配")
	}
	return nil
}