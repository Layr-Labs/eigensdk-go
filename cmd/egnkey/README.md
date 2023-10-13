## egnkey

### How to install
#### Install from source
```bash
   go install github.com/Layr-Labs/eigensdk-go/cmd/egnkey@latest
```

#### Build from source
Navigate to [egnkey](../egnkey/) directory and run
```bash
   go build
```

### How to run

To create in a random folder
```bash
   ./egnkey generate --key-type ecdsa --num-keys <num_key>
```

To create  in specific folder
```bash
   ./egnkey generate --key-type ecdsa --num-keys <num_key> --output-dir <path_to_folder>
```