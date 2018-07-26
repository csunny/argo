### Test

Terminal one
```
go run main.go -l 10000 -secio
```

Terminal two
```
go run main.go -l 10001 -d /ip4/127.0.0.1/tcp/10000/ipfs/Qmf4X6ypWtb6g1wFFASGgbkUqci5sEA3PYVvgFP81imDR8 -secio
```

Terminal three
```
go run main.go -l 10002 -d /ip4/127.0.0.1/tcp/10001/ipfs/QmZrCB3ELTTa1ChjqcyopZ9aRw23Sadt1HFaZ7ZaX4Mrxc -secio

```

