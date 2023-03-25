# Copyright (C) 2022 Storx Labs, Inc.
# See LICENSE for copying information.
#!/bin/bash

solc --bin --abi --overwrite -o build/ TestToken.sol \
    --base-path . \
    --include-path node_modules/
