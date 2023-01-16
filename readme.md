# lottery
**lottery** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Design

There are mainly two data types:

- Lottery: is a struct to store the lottery information

    ```go
    type Lottery struct {
        Index   uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
        Winner  string `protobuf:"bytes,2,opt,name=winner,proto3" json:"winner,omitempty"`
        Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
    }
    ```

- Bet: is a struct to store the bet lists of the given lottery

    ```go
    type Bet struct {
        Index      string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
        Amount     string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
        Suffix     string `protobuf:"bytes,3,opt,name=suffix,proto3" json:"suffix,omitempty"`
        SuffixHash string `protobuf:"bytes,4,opt,name=suffixHash,proto3" json:"suffixHash,omitempty"`
        Player     string `protobuf:"bytes,5,opt,name=player,proto3" json:"player,omitempty"`
    }
    ```

    `Index` format is like `{lottery index}/{bet owner address}` to specify bets in the given lottery round.

There are mainly two messages (transactions):

- add-bet: is used for adding bet

    ```cmd
    lotteryd tx lottery add-bet [token] [suffix]
    ```

    The `suffix` is used to calc the hash for the winner choosing. To hide the real suffix, we only store `SuffixHash` in this stage.
    If we have more bets than `BetThresCount` in the end_block, we lock the lottery through setting the `Winnder` as the module name (`lottery`).

- reveal-bet: is used to reveal the suffix (assume the lottery is already locked)

    ```cmd
    lotteryd tx lottery reveal-bet [suffix]
    ```

    We compare the `suffix` with the `SuffixHash` of the bet storage for the verify purpose and determine the winner using suffixes of all bets.
    Through this way, we can protect the block proposer's cheats, because there is no way to reveal suffix which is used to determine the winner before the lottery lock. After the lock, there is no way to update the suffix.

    We have another option of using `RVF`, it is more complicated.

### Demo

The `config.yml` file is set with the initial state. We can init and run the chain using the following commands.

```cmd
ignite chain serve --reset-once
```

Open the new shell, and execute the following command to run the demo script.

```cmd
go run ./script/demo
```

### Bonus Strategy

We can calculate the expected amount of every bet through Monte Carlo Simulation. The key point is to keep as lowest bet amount but not last.
The formula will be complicated, the best way is to calculate this expected value through Monte Carlo methods.