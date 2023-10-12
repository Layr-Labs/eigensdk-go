## elkey

### How to build
Navigate to [elkey](../elkey/) directory and run
```bash
   go build
```

### How to run

To create in a random folder
```bash
   ./elkey generate --key-type ecdsa --num-keys <num_key>
```

To create  in specific folder
```bash
   ./elkey generate --key-type ecdsa --num-keys <num_key> --output-dir <path_to_folder>
```