package keystore

import(
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"strings"
	logging "github.com/ipfs/go-log"
	ci "github.com/libp2p/go-libp2p-crypto"
)

var log = logging.Logger("keystore")

// Keystore 提供了一个key管理的接口
type Keystore interface{
	// 判断Key是否存在
	Has(string) (bool, error)
	// 将Key存储到keystore中, 如果已经存在相同名称的key，则返回已存在错误。
	Put(string, ci.PrivKey) error
	// 从keystore中获取key，如果不存在返回不存在错误
	Get(string) (ci.PrivKey, error)
	// 从Keystore中删除key
	Delete(string) error
	// 列出所有的key
	List()([]string, error)
}

var ErrNoSuchKey = fmt.Errorf("Key不存在")
var ErrKeyExists = fmt.Errorf("Key已经存在")

// FSKeystore 将keystore存储在磁盘文件中 
type FSKeystore struct{
	dir string
}

func validateName(name string) error{
	if name == ""{
		return fmt.Errorf("密钥名称至少包含一个字符")
	}
	if strings.Contains(name, "/"){
		return fmt.Errorf("密钥名称不能包含反斜杠")
	}
	if strings.HasPrefix(name, "."){
		return fmt.Errorf("密钥不能以.开头")
	}
	return nil
}

// NewFSKeystore 创建FSKeystore文件夹
func NewFSKeystore(dir string) (*FSKeystore, error){
	_, err := os.Stat(dir)
	if err != nil{
		if !os.IsNotExist(err){
			return nil, err
		}
		if err := os.Mkdir(dir, 0700); err != nil{
			return nil, err
		}
	}
	return &FSKeystore{dir}, nil
}

// Has 方法返回key是否存在
func (ks *FSKeystore) Has(name string) (bool, error){
	kp := filepath.Join(ks.dir, name)
	
	_, err := os.Stat(kp)
	if os.IsNotExist(err){
		return false, nil
	}
	if err != nil{
		return false, err
	}

	if err := validateName(name); err != nil{
		return false, err
	}

	return true, nil
}

// Put 将key存储在keystore当中。如果出现同名的key，则返回已存在错误
func (ks *FSKeystore) Put(name string, k ci.PrivKey) error{
	if err := validateName(name); err != nil{
		return err
	}

	b, err := k.Bytes()
	if err != nil{
		return err
	}

	kp := filepath.Join(ks.dir, name)
	_, err = os.Stat(kp)
	if err != nil{
		return ErrKeyExists
	}else if !os.IsNotExist(err){
		return err
	}

	fi, err := os.Create(kp)
	if err != nil{
		return err
	}
	defer fi.Close()

	_, err = fi.Write(b)
	return err
}

// Get 获取key的内容, 如果不存在，返回不存在错误
func (ks *FSKeystore) Get(name string) (ci.PrivKey, error){
	if err := validateName(name); err != nil{
		return nil, err
	}
	kp := filepath.Join(ks.dir, name)
	data, err :=  ioutil.ReadFile(kp)
	if err != nil{
		if os.IsNotExist(err){
			return nil, ErrNoSuchKey
		}
		return nil, err
	}
	return ci.UnmarshalPrivateKey(data)
}

// Delete 从keystore当中删除key
func (ks *FSKeystore) Delete(name string) error{
	if err := validateName(name); err != nil{
		return err
	}
	kp := filepath.Join(ks.dir, name)
	
	return os.Remove(kp)
}

// List 返回key列表
func (ks *FSKeystore) List() ([]string, error){
	dir, err := os.Open(ks.dir)
	if err != nil{
		return nil, err
	}
	dirs, err := dir.Readdirnames(0)
	if err != nil{
		return nil, err
	}

	list := make([]string, 0, len(dirs))

	for _, name := range dirs{
		err := validateName(name)
		if err == nil{
			list = append(list, name)
		}else{
			log.Warningf("略过无效的密钥文件: %s", name)
		}
	}
	
	return list, nil
}